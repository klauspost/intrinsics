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

func parsex86() {
	genimport = `import . "github.com/klauspost/intrinsics/x86"`
	f, err := os.Open("allintrin.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	parseHTMLX86(f)
	for _, pk := range packages {
		pk.goFile.Close()
		pk.asmFile.Close()
	}
}

func parseHTMLX86(f io.Reader) error {
	z := html.NewTokenizer(f)
	in := NewIntrinsic()
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			in.FinishX86()
			return z.Err()
		case html.TextToken:
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			switch strings.ToLower(string(tn)) {
			case "span", "div":
				if tt == html.StartTagToken {
					in = parseDivX86(z, in)
				}
			case "table":
				in = parseTableX86(z, in)
			}
		}
	}
	in.FinishX86()
	return nil
}

func parseDivX86(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
	more := true
	var k, v []byte
	for more {
		k, v, more = z.TagAttr()
		//		fmt.Println("attr:", string(k))
		switch string(k) {
		case "class":
			val := string(v)
			if strings.Contains(val, "intrinsic") {
				in.FinishX86()
				return NewIntrinsic()
			}
			switch val {
			case "cpuid":
				in.CpuID = getText(z)
			case "instruction":
				in.Instruction = strings.ToUpper(getText(z))
			case "rettype":
				in.RetType = fixTypeX86(getText(z))
			case "param_type":
				in.cParam = &Param{Type: fixTypeX86(getText(z))}
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
				in.Name = fixFuncNameX86(in.OrgName)
			case "operation":
				in.Operation = strings.TrimSpace(getText(z))
			default:
				//fmt.Println("unparsed class:", string(v))

			}
		}
	}
	return in
}

func (i Intrinsic) getFilesX86() pack {
	pk := i.getPackage()
	p, ok := packages[pk]
	if !ok {
		var err error
		err = os.MkdirAll("x86/"+pk, 0777)
		if err != nil {
			panic(err)
		}
		p.goFile, err = os.Create("x86/" + pk + "/" + pk + ".go")
		if err != nil {
			panic(err)
		}
		p.goFile.WriteString("// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!\n")
		p.goFile.WriteString("//\n// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!\n")
		p.goFile.WriteString("//\n// See https://github.com/klauspost/intrinsics\n")
		p.goFile.WriteString("package " + pk + "\n\n")
		p.goFile.WriteString(`import "github.com/klauspost/intrinsics/x86"` + "\n\n")

		p.goFile.WriteString(`var _ = x86.M64{}  // Make sure we use x86 package` + "\n\n")

		p.asmFile, err = os.Create("x86/" + pk + "/" + pk + "_amd64.s")
		if err != nil {
			panic(err)
		}
		packages[pk] = p
	}
	return p
}

func fixFuncNameX86(in string) string {
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

func parseTableX86(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
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

func fixTypeX86(s string) string {
	return fixType(s, "x86")
}

func (in Intrinsic) hasImmediateX86() bool {
	for _, param := range in.Params {
		if strings.Contains(param.Name, "imm") {
			return true
		}
	}
	return false
}

func (in Intrinsic) FinishX86() {
	if in.Name == "" {
		return
	}

	in.fixup()

	out := in.getFilesX86().goFile

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

	retparam := Param{Type: in.RetType}
	rettypeprint := in.RetType
	if strings.Contains(rettypeprint, ".") {
		rettypeprint = "(dst " + rettypeprint + ")"
	}

	fmt.Fprint(out, strings.Join(params, ", "), ") ", rettypeprint, " {\n")

	// We have a receiver pointer,
	if rework {
		fmt.Fprintln(out, "\t// FIXME: Rework to avoid possible return value as parameter.")
		if strings.HasPrefix(in.RetType, "x86.") {
			fmt.Fprint(out, "\treturn "+in.RetType+"{}")
		} else if in.RetType != "" {
			fmt.Fprint(out, "\treturn 0")
		}
		fmt.Fprint(out, "\n}\n")
		return
	}

	if in.RetType != "" {
		fmt.Fprint(out, "\treturn "+in.RetType+"("+in.AsmName+"(")
	} else {
		fmt.Fprint(out, "\t"+in.AsmName+"(")
	}
	for i, param := range in.Params {
		pn := param.Name
		if param.getNativeX86() != "" {
			// cast
			pn = param.getNativeX86() + "(" + pn + ")"
		}
		fmt.Fprint(out, pn)
		if i != len(in.Params)-1 {
			fmt.Fprint(out, ", ")
		}
	}
	if in.RetType != "" {
		fmt.Fprint(out, ")")
	}
	fmt.Fprintln(out, ")\n}")
	fmt.Fprintln(out, "")

	// ASM declaration
	fmt.Fprint(out, "func ", in.AsmName, "(")
	params = []string{}
	for _, param := range in.Params {
		typ := param.Type
		ot := param.getNativeX86()
		if ot != "" {
			typ = ot
		}
		params = append(params, fmt.Sprint(param.Name, " ", typ))
	}
	if retparam.getNativeX86() != "" {
		retparam = Param{Type: retparam.getNativeX86()}
	}
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n\n")

	//emptyType(in.RetType)
	// ASM
	out = in.getFilesX86().asmFile
	fmt.Fprint(out, "// func ", in.AsmName, "(")
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n")
	fmt.Fprint(out, "TEXT ·"+in.AsmName+"(SB),7,$0\n")
	load, off := in.Params.getAsmX86()
	fmt.Fprintln(out, load)
	retparam = Param{Type: in.RetType}

	fmt.Fprint(out, "\t// TODO: Code missing")
	if in.Instruction != "..." && in.Instruction != "" {
		if len(params) == 1 {
			if retparam.getReg(0) == "X0" && in.Params[0].getReg(0) == "X0" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[0].getReg(0))
			} else if retparam.getReg(0) == "M0" && in.Params[0].getReg(0) == "M0" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[0].getReg(0))
			} else if in.Params[0].getReg(0) != "" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s", in.Instruction, in.Params[0].getReg(0))
			} else {
				fmt.Fprintf(out, " - uses instrunction: %s", in.Instruction)
			}
		} else if len(params) == 2 {
			fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[1].getReg(1))
		} else {
			fmt.Fprintf(out, " - uses instrunction: %s", in.Instruction)
		}
	}
	fmt.Fprintln(out, "\n")

	// Attempts to write a return value
	if in.RetType != "" {
		// Guess a return reg
		rreg := retparam.getReg(0)

		if len(params) > 0 {
			rtry := retparam.getReg(len(params) - 1)
			if rtry != "" {
				rreg = rtry
			}
		}

		if retparam.getSizeX86() > 0 && retparam.getReg(0) == "R8" {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" $0, ret+"+strconv.Itoa(off)+"(FP)\n")
		} else if retparam.getSizeX86() > 8 {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" "+rreg+", ret+"+strconv.Itoa(off)+"(FP)\n")
		} else {
			fmt.Fprintf(out, "\t// Return size: %d\n", retparam.getSizeX86())
		}

	}
	fmt.Fprint(out, "\tRET\n")
	fmt.Fprint(out, "\n")
}

func (p Param) getSizeX86() int {
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "M64", "M64i":
		return 8
	case "M128", "M128i", "M128d":
		return 16
	case "M256", "M256i", "M256d":
		return 32
	case "M512", "M512i", "M512d":
		return 64
	case "int", "uint", "int64", "uint64", "float64", "uintptr", "Mmask64":
		return 8
	case "int32", "uint32", "float32", "Mmask32":
		return 4
	case "int16", "uint16", "Mmask16":
		return 2
	case "int8", "uint8", "Mmask8", "byte":
		return 1
	}
	return 0
}

func (p Param) getNativeX86() string {
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "M128":
		return "[4]float32"
	case "M256":
		return "[8]float32"
	case "M512":
		return "[16]float32"
	case "M128i":
		return "[16]byte"
	case "M256i":
		return "[32]byte"
	case "M512i":
		return "[64]byte"
	case "M128d":
		return "[2]float64"
	case "M256d":
		return "[4]float64"
	case "M512d":
		return "[8]float64"
	case "Mmask64":
		return "uint64"
	case "Mmask32":
		return "uint32"
	case "Mmask16":
		return "uint16"
	case "Mmask8":
		return "uint8"
	}
	return ""
}

func (p Params) getAsmX86() (string, int) {
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
