package main

import "fmt"

func main() {
	//The family provided by fmt, however, are built to be in production code.
	//They report predictably to stdout, unless otherwise specified.
	//They are more versatile (fmt.Fprint* can report to any io.Writer, such as os.Stdout, os.Stderr, or even a net.Conn type.) and are not implementation specific.
	fmt.Println("Hola mundo")
	// println is a function built into the language. It is in the Bootstrapping
	//println("Hola mundo")
}
