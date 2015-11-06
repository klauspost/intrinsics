package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Downloaded from
// https://software.intel.com/sites/landingpage/IntrinsicsGuide/#techs=MMX,SSE,SSE2,SSE3,SSSE3,SSE4_1,SSE4_2,AVX,AVX2,FMA,KNC,SVML,Other&avx512techs=AVX512F,AVX512BW,AVX512CD,AVX512DQ,AVX512ER,AVX512IFMA52,AVX512PF,AVX512VBMI

func parsesimd() {
	genimport = `import . "github.com/bjwbell/gensimd/simd"`
	f, err := os.Open("allintrin.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	parseHTMLSimd(f)
	for _, pk := range packages {
		pk.goFile.Close()
		pk.toAsmFile.WriteString("}\n")
		pk.toAsmFile.Close()
		pk.instrFile.WriteString(")\n")
		pk.instrFile.Close()

	}
}

func parseHTMLSimd(f io.Reader) error {
	z := html.NewTokenizer(f)
	in := NewIntrinsic()
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			in.FinishSimd()
			return z.Err()
		case html.TextToken:
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			switch strings.ToLower(string(tn)) {
			case "span", "div":
				if tt == html.StartTagToken {
					in = parseDivSimd(z, in)
				}
			case "table":
				in = parseTableSimd(z, in)
			}
		}
	}
	in.FinishSimd()
	return nil
}

func parseDivSimd(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
	more := true
	var k, v []byte
	for more {
		k, v, more = z.TagAttr()
		//		fmt.Println("attr:", string(k))
		switch string(k) {
		case "class":
			val := string(v)
			if strings.Contains(val, "intrinsic") {
				in.FinishSimd()
				return NewIntrinsic()
			}
			switch val {
			case "cpuid":
				in.CpuID = getText(z)
			case "instruction":
				in.Instruction = strings.TrimSpace(strings.ToUpper(getText(z)))
			case "rettype":
				in.RetType = fixTypeSimd(getText(z))
			case "param_type":
				in.cParam = &Param{Type: fixTypeSimd(getText(z))}
			case "param_name":
				in.cParam.Name = getText(z)
				if !in.Params.HasParam(in.cParam.Name) {
					in.Params = append(in.Params, *in.cParam)
				}
				in.cParam = nil
			case "description":
				in.Description = strings.TrimSpace(getTextR(z))
			case "name":
				in.OrgName = getText(z)
				in.Name = fixFuncName(in.OrgName)
			case "operation":
				in.Operation = strings.TrimSpace(getText(z))
			default:
				//fmt.Println("unparsed class:", string(v))

			}
		}
	}
	return in
}

func (i Intrinsic) getFilesSimd() pack {
	pk := i.getPackage()
	p, ok := packages[pk]
	if !ok {
		var err error
		err = os.MkdirAll("simd/"+pk, 0777)
		if err != nil {
			panic(err)
		}
		err = os.MkdirAll("codegen/"+pk, 0777)
		if err != nil {
			panic(err)
		}
		p.goFile, err = os.Create("simd/" + pk + "/" + pk + ".go")
		if err != nil {
			panic(err)
		}
		p.goFile.WriteString("// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!\n")
		p.goFile.WriteString("//\n// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!\n")
		p.goFile.WriteString("//\n// See https://github.com/klauspost/intrinsics\n")
		p.goFile.WriteString("package " + pk + "\n\n")
		p.goFile.WriteString(`import "github.com/bjwbell/gensimd/simd"` + "\n\n")

		p.goFile.WriteString(`var _ = simd.M128{}  // Make sure we use simd package` + "\n\n")

		p.toAsmFile, err = os.Create("codegen/" + pk + ".go")
		if err != nil {
			panic(err)
		}
		p.toAsmFile.WriteString("package codegen\n\n")
		p.toAsmFile.WriteString("import (\n")
		p.toAsmFile.WriteString("\t\"github.com/bjwbell/gensimd/simd/" + pk + "\"\n")
		p.toAsmFile.WriteString(")\n")
		p.toAsmFile.WriteString("\n")
		p.toAsmFile.WriteString("var " + pk + "ToGoAsm = map[" + pk + ".Instr]Instr{\n")

		p.instrFile, err = os.Create("codegen/" + pk + "/instructions.go")
		if err != nil {
			panic(err)
		}
		p.instrFile.WriteString("package " + pk + "\n\n")
		if err != nil {
			panic(err)
		}
		p.instrFile.WriteString("type Instr int\n")
		p.instrFile.WriteString(`
// SSE2 types
type M128i [16]byte
type M128 [4]float32
type M128d [2]float64

// AVX
type M256 [8]float32
type M256i [32]byte
type M256d [4]float64

// AVX2
type M512 [16]float32
type M512i [64]byte
type M512d [8]float64

type Mmask8 uint8
type Mmask16 uint16
type Mmask32 uint32
type Mmask64 uint64

const (
	INVALID Instr = iota
`)

		packages[pk] = p
	}
	return p
}

func fixFuncName(in string) string {
	if strings.HasPrefix(in, "_m_") {
		in = strings.TrimPrefix(in, "_m_")
	}
	out := CamelCase(in)
	if strings.HasPrefix(out, "mm512") {
		out = strings.TrimPrefix(out, "m")
	}
	if strings.HasPrefix(out, "mm256") {
		out = strings.TrimPrefix(out, "m")
	}
	if strings.HasPrefix(out, "mm512") {
		out = strings.TrimPrefix(out, "m")
	}
	out = strings.TrimPrefix(out, "mm")
	out = strings.Title(out)
	return out
}

func parseTableSimd(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
	in.Performance = make(map[string]Timing)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return in
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			tns := strings.ToLower(string(tn))
			switch tns {
			case "tr":
				if tt == html.StartTagToken {
					n := 0
					p := Timing{}
					for {
						tt = z.Next()
						tn, _ = z.TagName()
						tns = strings.ToLower(string(tn))
						if tt == html.EndTagToken && tns == "tr" {
							break
						}
						if tt == html.StartTagToken && tns == "td" {
							switch n {
							case 0:
								p.Arch = getText(z)
							case 1:
								p.Latency, _ = strconv.ParseFloat(getText(z), 64)
							case 2:
								p.Throughput, _ = strconv.ParseFloat(getText(z), 64)
								in.Performance[p.Arch] = p
							}
							n++
						}
					}
				} else {
					panic("tr ended")
				}
			case "table":
				if tt == html.EndTagToken {
					return in
				} else {
					panic("table started")

				}
			}
		}
	}
	return in
}

func fixTypeSimd(s string) string {
	return fixType(s, "simd")
}

var wroteAsm = make(map[string]struct{})

func (in Intrinsic) FinishSimd() {
	if in.Name == "" {
		return
	}

	in.fixup()

	out := in.getFilesSimd().goFile

	if in.shouldSkip() {
		fmt.Fprintf(out, "\n// Skipped: %s. Contains pointer parameter.\n\n", in.OrgName)
		return
	}
	rework := in.shouldRework()

	//fmt.Println("Finished:", in)
	desc := in.Name + ": " + in.Description
	desc = strings.Replace(desc, "\r\n", "\n", -1)
	desc = strings.Replace(desc, "\r", "\n", -1)
	desc = WrapString(desc, 76)
	desc = strings.Replace(desc, "\n", "\n// ", -1)
	fmt.Fprintln(out, "\n//", desc, "\n//")
	op := strings.Replace(strings.TrimSpace(in.Operation), "\n", "\n//\t\t", -1)
	fmt.Fprint(out, "//\t\t", op, "\n//\n")
	fmt.Fprintln(out, "// Instruction:", "'"+in.Instruction+"'. Intrinsic:", "'"+in.OrgName+"'.")
	if len(in.CpuID) > 0 {
		fmt.Fprintln(out, "// Requires", in.CpuID+".")
	}
	if rework {
		fmt.Fprintln(out, "//\n// FIXME: Will likely need to be reworked (has pointer parameter).")
	}
	if in.hasImmediateX86() {
		fmt.Fprintln(out, "//\n// FIXME: Requires compiler support (has immediate)")
	}
	fmt.Fprint(out, "func ", in.Name, "(")
	params := []string{}
	for _, param := range in.Params {
		params = append(params, fmt.Sprint(param.Name, " ", param.Type))
	}

	//retparam := Param{Type: in.RetType}
	rettypeprint := in.RetType
	if strings.Contains(rettypeprint, ".") {
		rettypeprint = "(dst " + rettypeprint + ")"
	}

	fmt.Fprint(out, strings.Join(params, ", "), ") ", rettypeprint, " {\n")

	fmt.Fprintln(out, "\t"+`panic("unreachable")`)
	fmt.Fprintln(out, "}")
	fmt.Fprintln(out, "")

	out = in.getFilesSimd().toAsmFile
	if in.Instruction == "..." || in.Instruction == "" {
		fmt.Fprintln(out, "\t// Add Composite: "+in.Package+"."+in.Name)
	} else {
		fmt.Fprintln(out, "\t"+in.Package+"."+in.Name+":"+in.Instruction+",")
	}

	if in.Instruction != "..." && in.Instruction != "" {
		_, ok := wroteAsm[in.Package+in.Instruction]
		if !ok {
			out = in.getFilesSimd().instrFile
			fmt.Fprintln(out, "\t"+in.Instruction)
		}
		wroteAsm[in.Package+in.Instruction] = struct{}{}
	}
}

func (p Params) getAsmSimd() (string, int) {
	out := ""
	atbytes := 0
	for i, param := range p {
		size := param.getSizeX86()
		if size == 0 {
			return "\t// FIXME: Unimplemented. Unknown size of type " + param.Type + "\n", 0
		}
		reg := param.getReg(i)
		if reg == "" {
			return "\t// FIXME: Unimplemented. Unknown register for type " + param.Type + "\n", 0
		}
		postfix := param.getPostfix()
		if postfix == "" {
			out += "\t// FIXME: Unimplemented. Unknown MOVE postfix for type " + param.Type + "\n\t//"
		}
		if strings.Contains(param.Name, "imm") {
			out += "\t// FIXME: Immediate parameter should be removed (" + param.Name + " " + param.Type + ")\n"
		} else {
			out += "\tMOV" + postfix + " " + param.Name + "+" + strconv.Itoa(atbytes) + "(FP)," + reg + "\n"
		}
		if size < 4 {
			size = 4
		}
		atbytes += size
	}
	return out, atbytes
}
