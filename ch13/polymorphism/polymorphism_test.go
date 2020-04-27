package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {

}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	gopro := &GoProgrammer{}
	javapro := new(JavaProgrammer)
	writeFirstProgram(gopro)
	writeFirstProgram(javapro)
}