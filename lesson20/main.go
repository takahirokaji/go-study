package main

import "fmt"

func main() {
	var x interface{}
	//すべての方を汎用的に表すもの
	// すべての型と互換性がある
	// 初期値としてnilを持っている

	// 初期値は nil
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	//x = 2
	// これはだめです
	//fmt.Println(x + 3)
}
