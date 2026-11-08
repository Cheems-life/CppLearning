package main

import (
	"fmt"

	_case "example.com/day02/case"
)

func main() {
	a := 10
	b := 20
	fmt.Println(_case.SumCase(a, b))

	_case.ReferenceCase(&a, &b)
	fmt.Println(a, b)

	fmt.Println(_case.G_val)
	_case.ScopeCase(a, b)
	fmt.Println(_case.G_val)

	user := _case.NewUser("nick", 18)

	fmt.Println(user.GetName(), user.GetAge())
}
