package main

//this is an import package, allows multiple imports to the program
import (
	"fmt"
)

/*
 * The main function, when part of the main package, identifies the entry point of an application
 */
func main() {
	fmt.Println("Hello, World!")
}

/*
 * -Whitespace is not 'necessary' in Go like python but formatting for readability is strongly encouraged. Try:
 *  remove the tab inside main where print function is called.
 * -Some syntax is enforced and will give an error (Syntax error). Try: move the opening brace of the main function
 *  to a new line. (errors because of automatic semicolon insertion)
 * -If an imported package isn't used, an error will be returned. Try: Comment out hello world print statement
 * -Go enforces automatic semicolon insertion. Can include or exclude but like Js, may cause some issues if you try
 *  to run multiple statements in one line w/o semicolons.
 */
