package main

import (
	"fmt"
)

func main() {
	basic()
	fmt.Println("\n##############################################################################################\n")
	scan()
	fmt.Println("\n##############################################################################################\n")
	sscanf()
	fmt.Println("\n##############################################################################################\n")
	fprintf()
	fmt.Println("\n##############################################################################################\n")
	verbs()
	fmt.Println("\n##############################################################################################\n")
	structs()
	fmt.Println("\n##############################################################################################\n")
	numbers()
	fmt.Println("\n##############################################################################################\n")
	unicode()
	fmt.Println("\n##############################################################################################\n")
	castPointerToString()
	fmt.Println("\n##############################################################################################\n")
}

func basic() {
	message := "hello, fmt"
	fmt.Print("fmt.Print: hello, fmt\n")
	fmt.Println("fmt.Println: hello, fmt")
	fmt.Printf("fmt.Printf: %s\n", message)

	name := "Kim"
	message = fmt.Sprintf("Hello, %s", name)
	fmt.Println(message)
}

func scan() {
	fmt.Println("Scan reads from stdin")
	var name string
	var age int
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func sscanf() {
	fmt.Println("Sscanf supports formatted inputs")
	var first string
	var last string
	var age int
	_, _ = fmt.Sscanf("Kim Pettersen - 32", "%s %s - %d", &first, &last, &age)
	fmt.Printf("First Name: %s, Last Name: %s Age: %d\n", first, last, age)
}

type Database struct {
	N string
}

func (Database) Write(b []byte) (int, error) {
	fmt.Printf("Writing to database: %s", b)
	return 0, nil
}

func fprintf() {
	fmt.Println("Fprint/Fscan uses an io.Writer/io.Reader")
	var name, age = "Kim", 32
	database := Database{}
	_, _ = fmt.Fprintf(database, "Name: %s, Age: %d\n", name, age)
}

func verbs() {
	name := "Kim"
	var age int
	fmt.Printf("String: %s\n", name)
	fmt.Printf("Base 10 digit: %d\n", 17)

	fmt.Printf("String: %v\n", name)
	fmt.Printf("Undeclared variable: %v\n", age)
}

type person struct {
	name        string
	age         int
	nationality *string
}

func structs() {
	var p person
	fmt.Printf("Initialized: %v\n", p)

	p.name = "Kim"
	p.age = 32

	fmt.Printf("Declared: %v\n", p)
	fmt.Printf("Type of p: %T\n", p)

	fmt.Printf("Field names and values: %+v\n", p)
	fmt.Printf("Type, field names and values: %#v\n", p)
}

func numbers() {

	fmt.Printf("Base 2: %b\n", 255)
	fmt.Printf("Base 2: %b\n", 0xff)

	fmt.Printf("Base 10: %d\n", 0xff)
	fmt.Printf("Base 10: %d\n", 0b11111111)

	fmt.Printf("Base 16: %X\n", 255)
	fmt.Printf("Base 16: %x\n", 11111111)

	fmt.Printf("Scientific notation: %e\n", 0.00000000001876)

}
func unicode() {
	fmt.Printf("Unicode: %U\n", 32)
	fmt.Printf("Unicode: %U\n", '😀')
}

func castPointerToString() {
	fmt.Println("You should probably only use this when you print something")
	var pointerToString *string
	a := fmt.Sprintf("%v", pointerToString)
	fmt.Printf("pointerToString: %v\n", a)
}
