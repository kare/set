package main // import "kkn.fi/set/genset"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type setTemplate struct {
	ShortTypeName string
	TypeName      string
	Type          string
}

var templates = []setTemplate{
	{"String", "StringSet", "string"},
	// same size as uint
	{"Int", "IntSet", "int"},
	// either 32 or 64 bits
	//{"Uint", "UintSet", "uint"},

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	//{"Uint8", "Uint8Set", "uint8"},
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	//{"Uint16", "Uint16Set", "uint16"},
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	//{"Uint32", "Uint32Set", "uint32"},
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	//{"Uint64", "Uint64Set", "uint64"},

	// int8        the set of all signed  8-bit integers (-128 to 127)
	//{"Int8", "Int8Set", "int8"},
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	//{"Int16", "Int16Set", "int16"},
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	//{"Int32", "Int32Set", "int32"},
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	//{"Int64", "Int64Set", "int64"},

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	//{"Float32", "Float32Set", "float32"},
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	//{"Float64", "Float64Set", "float64"},

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	//{"Complex64", "Complex64Set", "complex64"},
	// complex128  the set of all complex numbers with float64 real and imaginary parts
	//{"Complex128", "Complex128Set", "complex128"},

	// byte        alias for uint8
	//{"Byte", "ByteSet", "byte"},
	// rune        alias for int32
	//{"Rune", "RuneSet", "rune"},
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("genset: ")

	tmpl := template.New("Set")
	listTmpl, err := ioutil.ReadFile("set.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	t, err := tmpl.Parse(string(listTmpl))
	if err != nil {
		log.Fatal(err)
	}
	for _, template := range templates {
		var (
			name             = fmt.Sprintf("%v.go", template.Type)
			mode             = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			perm os.FileMode = 0644
		)
		file, err := os.OpenFile(name, mode, perm)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if err := t.Execute(file, template); err != nil {
			log.Fatal(err)
		}
	}
}
