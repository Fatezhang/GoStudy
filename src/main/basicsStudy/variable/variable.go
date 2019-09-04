package variable

import "fmt"

func Function() {
	var a = 0
	fmt.Println(a)
	var p *int = &a
	fmt.Printf("%x", p)
	fmt.Println()
	a = 100
	fmt.Printf("a = %d\n", a)
	*p = 250
	fmt.Printf("a = %d\n", a)
	fmt.Printf("%x", p)
	fmt.Println()
}
