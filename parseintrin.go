package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

type Param struct {
	Name string
	Type string
}

type Params []Param

type Intrinsic struct {
	OrgName     string
	Name        string
	AsmName     string
	Instruction string
	Description string
	RetType     string
	Params      Params
	CpuID       string
	Operation   string
	Performance Perf
	Package     string
	cParam      *Param
}

type Timing struct {
	Arch       string
	Latency    float64
	Throughput float64
}

type Perf map[string]Timing

//var m64, m128, m256, m512, gen *os.File
//var m64a, m128a, m256a, m512a, gena *os.File
type pack struct {
	goFile  *os.File
	asmFile *os.File
}

const genimport = `import . "github.com/klauspost/intrinsics/x86"`

// Downloaded from
// https://software.intel.com/sites/landingpage/IntrinsicsGuide/#techs=MMX,SSE,SSE2,SSE3,SSSE3,SSE4_1,SSE4_2,AVX,AVX2,FMA,KNC,SVML,Other&avx512techs=AVX512F,AVX512BW,AVX512CD,AVX512DQ,AVX512ER,AVX512IFMA52,AVX512PF,AVX512VBMI

func main() {
	f, err := os.Open("allintrin.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	parseHTML(f)
	for _, pk := range packages {
		pk.goFile.Close()
		pk.asmFile.Close()
	}
}

func NewIntrinsic() *Intrinsic {
	return &Intrinsic{}
}

func parseHTML(f io.Reader) error {
	z := html.NewTokenizer(f)
	in := NewIntrinsic()
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			in.Finish()
			return z.Err()
		case html.TextToken:
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			switch strings.ToLower(string(tn)) {
			case "span", "div":
				if tt == html.StartTagToken {
					in = parseDiv(z, in)
				}
			case "table":
				in = parseTable(z, in)
			}
		}
	}
	in.Finish()
	return nil
}

func parseDiv(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
	more := true
	var k, v []byte
	for more {
		k, v, more = z.TagAttr()
		//		fmt.Println("attr:", string(k))
		switch string(k) {
		case "class":
			val := string(v)
			if strings.Contains(val, "intrinsic") {
				in.Finish()
				return NewIntrinsic()
			}
			switch val {
			case "cpuid":
				in.CpuID = getText(z)
			case "instruction":
				in.Instruction = strings.ToUpper(getText(z))
			case "rettype":
				in.RetType = fixType(getText(z))
			case "param_type":
				in.cParam = &Param{Type: fixType(getText(z))}
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

func (i Intrinsic) getPackage() string {
	pk := strings.ToLower(i.CpuID)
	if pk == "" {
		pk = "misc"
	}
	if pk == "sse4.1" || pk == "sse4.2" {
		pk = "sse4"
	}
	return pk
}

var packages = make(map[string]pack)

func (i Intrinsic) getFiles() pack {
	pk := i.getPackage()
	p, ok := packages[pk]
	if !ok {
		var err error
		err = os.MkdirAll("x86/"+pk, 666)
		if err != nil {
			panic(err)
		}
		p.goFile, err = os.Create("x86/" + pk + "/" + pk + ".go")
		if err != nil {
			panic(err)
		}
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

func parseTable(z *html.Tokenizer, in *Intrinsic) *Intrinsic {
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

func getText(z *html.Tokenizer) string {
	tt := z.Next()
	switch tt {
	case html.ErrorToken:
		panic(z.Err())
	case html.TextToken:
		return string(z.Text())
	}
	return ""
}

func getTextR(z *html.Tokenizer) string {
	r := ""
	depth := 1
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			panic(z.Err())
		case html.TextToken:
			r += string(z.Text())
		case html.StartTagToken:
			tn, _ := z.TagName()
			tns := strings.ToLower(string(tn))
			switch tns {
			case "div":
				r += "\r"
				depth++
			case "span":
				r += "'"
				depth++
			}
		case html.EndTagToken:
			tn, _ := z.TagName()
			tns := strings.ToLower(string(tn))
			switch tns {
			case "div":
				depth--
			case "span":
				r += "'"
				depth--
			}
		}
		if depth == 0 {
			return r
		}
	}
	return r
}

func fixType(s string) string {
	pointer := false
	if strings.Contains(s, "*") {
		pointer = true
	}
	r := CamelCase(s)
	if len(r) == 0 {
		return ""
	}
	rb := []byte(r)
	if rb[0] == 'm' {
		rb[0] = 'M'
		r = "x86." + string(rb)
	}

	switch r {
	case "void":
		if pointer {
			r = "uintptr"
		} else {
			r = ""
		}
	case "char", "charConst":
		r = "byte"
	case "unsignedChar":
		r = "uint8"
	case "unsignedShort":
		r = "uint16"
	case "sizeT", "constInt", "intConst":
		r = "int"
	case "int64Const":
		r = "int"
	case "unsignedInt64":
		r = "uint64"
	case "unsigned", "constUnsignedInt":
		r = "uint"
	case "unsignedInt", "unsignedLong":
		r = "uint32"
	case "unsignedInt32":
		r = "uint32"
	case "voidConst":
		r = "uintptr"
	case "constVoid":
		r = "uintptr"
	case "mem_addr":
		r = "uintptr"
	case "float", "constFloat":
		r = "float32"
	case "double", "constDouble":
		r = "float64"
	case "short":
		r = "int16"
	case "floatConst":
		r = "uintptr"
	case "doubleConst":
		r = "uintptr"
	case "longLong":
		r = "int64"
	case "constMMCMPINTENUM":
		r = "uint8"
	}

	if r == "uintptr" || r == "" {
		return r
	}

	if pointer {
		r = "*" + r
	}
	return r

}

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}

func (p Params) HasParam(s string) bool {
	for _, param := range p {
		if param.Name == s {
			return true
		}
	}
	return false
}

var usedNames = make(map[string]struct{})

func (in *Intrinsic) fixup() {
	in.Package = in.getPackage()
	name := in.Name
	next := 2
	for {
		_, ok := usedNames[in.Package+name]
		if !ok {
			break
		}
		fmt.Println("Warning: Name conflict:", in.Package+"/"+name+"(...)")
		name = fmt.Sprintf("%s%d", in.Name, next)
		next++

	}
	usedNames[in.Package+name] = struct{}{}
	in.Name = name
	a := []rune(name)
	a[0] = unicode.ToLower(a[0])

	in.AsmName = string(a)
	for i, param := range in.Params {
		switch param.Name {
		case "type":
			in.Params[i].Name = "typ"
		case "func":
			in.Params[i].Name = "fnc"
		}
		if param.Type == "" {
			in.Params[i].Type = "uintptr"
		}
	}

}

func (in Intrinsic) shouldSkip() bool {
	if in.RetType == "uintptr" {
		return true
	}
	if strings.HasPrefix(in.RetType, "*") {
		return true
	}
	for _, param := range in.Params {
		if param.Type == "uintptr" {
			return true
		}
	}
	return false
}

func (in Intrinsic) shouldRework() bool {
	if strings.HasPrefix(in.RetType, "*") {
		return true
	}
	for _, param := range in.Params {
		if strings.HasPrefix(param.Type, "*") {
			return true
		}
	}
	return false
}

func (in Intrinsic) hasImmediate() bool {
	for _, param := range in.Params {
		if strings.Contains(param.Name, "imm") {
			return true
		}
	}
	return false
}

func (in Intrinsic) Finish() {
	if in.Name == "" {
		return
	}

	in.fixup()

	out := in.getFiles().goFile

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
	if in.hasImmediate() {
		fmt.Fprintln(out, "//\n// FIXME: Requires compiler support (has immediate)")
	}
	fmt.Fprint(out, "func ", in.Name, "(")
	params := []string{}
	for _, param := range in.Params {
		params = append(params, fmt.Sprint(param.Name, " ", param.Type))
	}
	fmt.Fprint(out, strings.Join(params, ", "), ") ", in.RetType, " {\n")

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
		if param.getNative() != "" {
			// cast
			pn = param.getNative() + "(" + pn + ")"
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
		ot := param.getNative()
		if ot != "" {
			typ = ot
		}
		params = append(params, fmt.Sprint(param.Name, " ", typ))
	}
	retparam := Param{Type: in.RetType}
	if retparam.getNative() != "" {
		retparam = Param{Type: retparam.getNative()}
	}
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n\n")

	//emptyType(in.RetType)
	// ASM
	out = in.getFiles().asmFile
	fmt.Fprint(out, "// func ", in.AsmName, "(")
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n")
	fmt.Fprint(out, "TEXT ·"+in.AsmName+"(SB),7,$0\n")
	load, off := in.Params.getAsm()
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

		if retparam.getSize() > 0 && retparam.getReg(0) == "R8" {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" $0, ret+"+strconv.Itoa(off)+"(FP)\n")
		} else if retparam.getSize() > 8 {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" "+rreg+", ret+"+strconv.Itoa(off)+"(FP)\n")
		} else {
			fmt.Fprintf(out, "\t// Return size: %d\n", retparam.getSize())
		}

	}
	fmt.Fprint(out, "\tRET\n")
	fmt.Fprint(out, "\n")
}

func (p Param) getSize() int {
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

func (p Param) getNative() string {
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
	case "uintptr":
		return "uintptr"
	}
	return ""
}

func (p Param) getPostfix() string {
	size := p.getSize()
	switch size {
	case 1:
		return "B"
	case 2:
		return "W"
	case 4:
		return "L"
	case 8:
		return "Q"
	case 16:
		return "OU"
	}
	return ""
}

func (p Param) getReg(n int) string {
	if n > 7 {
		return ""
	}
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "M64", "M64i":
		return "M" + strconv.Itoa(n)
	case "M128", "M128i", "M128d":
		return "X" + strconv.Itoa(n)
	case "M256", "M256i", "M256d":
		return "Y" + strconv.Itoa(n)
	case "M512", "M512i", "M512d":
		return "Z" + strconv.Itoa(n)
	case "int", "uint", "int64", "uint64", "float64", "uintptr", "Mmask64":
		return "R" + strconv.Itoa(8+n)
	case "int32", "uint32", "float32", "Mmask32":
		return "R" + strconv.Itoa(8+n)
	case "int16", "uint16", "Mmask16":
		return "R" + strconv.Itoa(8+n)
	case "int8", "uint8", "Mmask8", "byte":
		return "R" + strconv.Itoa(8+n)
	}
	return ""
}

func (p Params) getAsm() (string, int) {
	out := ""
	atbytes := 0
	for i, param := range p {
		size := param.getSize()
		if size == 0 {
			return "\t// FIXME: Unimplemented. Unknown size of type " + param.Type + "\n", 0
		}
		reg := param.getReg(i)
		if reg == "" {
			return "\t// FIXME: Unimplemented. Unknown register for type " + param.Type + "\n", 0
		}
		postfix := param.getPostfix()
		if postfix == "" {
			return "\t// FIXME: Unimplemented. Unknown MOVE postfix for type " + param.Type + "\n", 0
		}

		out += "\tMOV" + postfix + " " + param.Name + "+" + strconv.Itoa(atbytes) + "(FP)," + reg + "\n"
		if size < 4 {
			size = 4
		}
		atbytes += size
	}
	return out, atbytes
}

func emptyType(s string) string {
	if strings.HasPrefix(s, "int") {
		return "0"
	}
	if strings.HasPrefix(s, "uint") {
		return "0"
	}
	if strings.HasPrefix(s, "Mmask") {
		return s + "(0)"
	}
	if strings.HasPrefix(s, "float") {
		return "0.0"
	}
	switch s {
	case "string":
		return `""`
	case "byte":
		return "0"
	case "unsafe.Pointer":
		return s + "(uintptr(0))"
	default:
		return s + "{}"
	}
}

// WrapString wraps the given string within lim width in characters.
//
// Wrapping is currently naive and only happens at white-space. A future
// version of the library will implement smarter wrapping. This means that
// pathological cases can dramatically reach past the limit, such as a very
// long word.
func WrapString(s string, lim uint) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer
	var skipline bool

	for _, char := range s {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+uint(spaceBuf.Len()) > lim {
					current = 0
				} else {
					current += uint(spaceBuf.Len())
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
			} else {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
			skipline = false
		} else if unicode.IsSpace(char) {
			if current == 0 && spaceBuf.Len() > 2 {
				skipline = true
			}

			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}

			spaceBuf.WriteRune(char)
		} else {

			wordBuf.WriteRune(char)

			if current+uint(spaceBuf.Len()+wordBuf.Len()) > lim && uint(wordBuf.Len()) < lim && !skipline {
				buf.WriteRune('\n')
				current = 0
				spaceBuf.Reset()
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+uint(spaceBuf.Len()) <= lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}

	return buf.String()
}
