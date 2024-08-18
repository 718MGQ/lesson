package ptr

import "fmt"

func IntPtr() {

	var a int = 10

	var pa *int

	// &取地址
	pa = &a

	a = 20

	fmt.Println(*pa)
}

func ChangeA(a *int) {
	if a != nil {
		*a = 5
	}
}
