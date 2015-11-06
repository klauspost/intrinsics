package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

type Param struct {
	Name  string
	Type  string
	Const bool
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
	goFile      *os.File
	toInstrFile *os.File
	asmFile     *os.File
}

var genimport = `import . "github.com/klauspost/intrinsics/x86"`

func main() {
	f, err := os.Open("allintrin.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var arm = flag.Bool("arm", false, "simd")
	var x86 = flag.Bool("x86", false, "intel intrinsics")
	var simd = flag.Bool("simd", false, "Dump tables for gensimd")
	flag.Parse()
	if !*arm && !*x86 && !*simd {
		fmt.Println("Please specify at least one flag")
	}
	if *arm {
		parsearm()
	}
	if *x86 {
		parsex86()
	}
	if *simd {
		parsesimd()
	}
}

func NewIntrinsic() *Intrinsic {
	return &Intrinsic{}
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
		case "imm8":
			in.Params[i].Type = "byte"
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

func (p Param) getPostfix() string {
	size := p.getSizeX86()
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
}

func fixType(s string, importname string) string {
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
		r = importname + "." + string(rb)
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
