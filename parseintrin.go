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

var m64, m128, m256, m512, gen *os.File
var m64a, m128a, m256a, m512a, gena *os.File

const genimport = `import . "github.com/klauspost/intrinsics/x86"`

func main() {
	f, err := os.Open("allintrin.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m64, err = os.Create("x86/m64/m64.go")
	if err != nil {
		panic(err)
	}
	defer m64.Close()
	_, err = m64.WriteString("package m64\n\n" + genimport + "\n")
	if err != nil {
		panic(err)
	}
	m128, err = os.Create("x86/m128/m128.go")
	if err != nil {
		panic(err)
	}
	defer m128.Close()
	_, err = m128.WriteString("package m128\n\n" + genimport + "\n")
	if err != nil {
		panic(err)
	}
	m256, err = os.Create("x86/m256/m256.go")
	if err != nil {
		panic(err)
	}
	defer m256.Close()
	_, err = m256.WriteString("package m256\n\n" + genimport + "\n")
	if err != nil {
		panic(err)
	}
	m512, err = os.Create("x86/m512/m512.go")
	if err != nil {
		panic(err)
	}
	defer m512.Close()
	_, err = m512.WriteString("package m512\n\n" + genimport + "\n")
	if err != nil {
		panic(err)
	}
	gen, err = os.Create("x86/intrin.go")
	if err != nil {
		panic(err)
	}
	defer gen.Close()
	_, err = gen.WriteString("package x86\n\n")
	if err != nil {
		panic(err)
	}

	m64a, err = os.Create("x86/m64/m64_amd64.s")
	if err != nil {
		panic(err)
	}
	defer m64a.Close()
	m128a, err = os.Create("x86/m128/m128_amd64.s")
	if err != nil {
		panic(err)
	}
	defer m128a.Close()
	m256a, err = os.Create("x86/m256/m256_amd64.s")
	if err != nil {
		panic(err)
	}
	defer m256a.Close()
	m512a, err = os.Create("x86/m512/m512_amd64.s")
	if err != nil {
		panic(err)
	}
	defer m512a.Close()
	gena, err = os.Create("x86/intrin_amd64.s")
	if err != nil {
		panic(err)
	}

	/*	fmt.Println("package intrin")
		fmt.Println("")
		fmt.Println("import \"unsafe\"")*/
	parseHTML(f)
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

func getFile(pack string) *os.File {
	switch pack {
	case "m64":
		return m64
	case "m128":
		return m128
	case "m256":
		return m256
	case "m512":
		return m512
	}
	return gen
}

func getAsmFile(pack string) *os.File {
	switch pack {
	case "m64":
		return m64a
	case "m128":
		return m128a
	case "m256":
		return m256a
	case "m512":
		return m512a
	}
	return gena
}

func getPackage(i Intrinsic) string {
	in := i.OrgName
	if strings.HasPrefix(in, "_mm256") {
		return "m256"
	}
	if strings.HasPrefix(in, "_mm512") {
		return "m512"
	}
	if strings.HasPrefix(in, "_m_") {
		return "m64"
	}
	in = strings.TrimPrefix(in, "_mm")
	if strings.Contains(in, "_sd") {
		return "m128"
	}
	if strings.Contains(in, "_si128") {
		return "m128"
	}
	if strings.Contains(in, "_ss") {
		return "m128"
	}
	if strings.Contains(in, "_ps") {
		return "m128"
	}
	if strings.Contains(in, "_ph") {
		return "m128"
	}
	if strings.Contains(in, "_pd") {
		return "m128"
	}
	if strings.Contains(in, "_epi") {
		return "m128"
	}
	if strings.Contains(in, "_pi") {
		return "m64"
	}
	if strings.Contains(in, "_cvt") {
		return "m128"
	}
	if strings.Contains(in, "_si") {
		return "m64"
	}
	if strings.Contains(in, "_epu") {
		return "m128"
	}
	if strings.Contains(in, "_pu") {
		return "m64"
	}
	if strings.Contains(in, "_su") {
		return "m64"
	}
	if strings.Contains(i.CpuID, "AVX512") {
		return "m512"
	}
	return "intrin"
}

func fixFuncName(in string) string {
	if strings.HasPrefix(in, "_m_") {
		in = strings.TrimPrefix(in, "_m_")
	}
	out := CamelCase(in)
	if strings.HasPrefix(out, "mm512") {
		out = strings.TrimPrefix(out, "mm512")
	}
	if strings.HasPrefix(out, "mm256") {
		out = strings.TrimPrefix(out, "mm256")
	}
	if strings.HasPrefix(out, "mm512") {
		out = strings.TrimPrefix(out, "m256")
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
	r := CamelCase(s)
	if len(r) == 0 {
		return ""
	}
	rb := []byte(r)
	if rb[0] == 'm' {
		rb[0] = 'M'
	}
	r = string(rb)

	switch r {
	case "void":
		r = ""
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
		r = "float32"
	case "doubleConst":
		r = "float64"
	case "longLong":
		r = "int64"
	case "constMMCMPINTENUM":
		r = "uint8"
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
	in.Package = getPackage(*in)
	name := in.Name
	next := 1
	for {
		_, ok := usedNames[in.Package+name]
		if !ok {
			break
		}
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

func (in Intrinsic) Finish() {
	if in.Name == "" {
		return
	}

	in.fixup()
	out := getFile(in.Package)

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
	fmt.Fprint(out, "func ", in.Name, "(")
	params := []string{}
	for _, param := range in.Params {
		params = append(params, fmt.Sprint(param.Name, " ", param.Type))
	}
	fmt.Fprint(out, strings.Join(params, ", "), ") ", in.RetType, " {\n")

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
	out = getAsmFile(in.Package)
	fmt.Fprint(out, "// func ", in.AsmName, "(")
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n")
	fmt.Fprint(out, "TEXT ·"+in.AsmName+"(SB),7,$0\n")
	load, off := in.Params.getAsm()
	fmt.Fprintln(out, load)

	fmt.Fprintln(out, "\t//TODO: Code missing\n")

	// Attempts to write a return value
	if in.RetType != "" {
		retparam = Param{Type: in.RetType}
		if retparam.getSize() > 0 && retparam.getReg(0) == "R8" {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" $0, ret+"+strconv.Itoa(off)+"(FP)\n")
		} else if retparam.getSize() > 8 {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" "+retparam.getReg(0)+", ret+"+strconv.Itoa(off)+"(FP)\n")
		} else {
			fmt.Fprintf(out, "\t// Return size: %d\n", retparam.getSize())
		}

	}
	fmt.Fprint(out, "\tRET\n")
	fmt.Fprint(out, "\n")
}

func (p Param) getSize() int {
	switch p.Type {
	case "M64", "M64i":
		return 8
	case "M128", "M128i", "M128d":
		return 16
	case "M256", "M256i", "M256d":
		return 32
	case "M512", "M512i", "M512d":
		return 64
	case "int", "uint", "int64", "uint64", "float64", "unsafe.Pointer", "Mmask64":
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
	switch p.Type {
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
	}
	switch p.Type {
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
	switch p.Type {
	case "M64", "M64i":
		return "M" + strconv.Itoa(n)
	case "M128", "M128i", "M128d":
		return "X" + strconv.Itoa(n)
	case "M256", "M256i", "M256d":
		return "Y" + strconv.Itoa(n)
	case "M512", "M512i", "M512d":
		return "Z" + strconv.Itoa(n)
	case "int", "uint", "int64", "uint64", "float64", "unsafe.Pointer", "Mmask64":
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
			return "\t// Unimplemented. Unknown size of type " + param.Type, 0
		}
		reg := param.getReg(i)
		if reg == "" {
			return "\t// Unimplemented. Unknown register for type " + param.Type, 0
		}
		postfix := param.getPostfix()
		if postfix == "" {
			return "\t// Unimplemented. Unknown MOVE postfix for type " + param.Type, 0
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
