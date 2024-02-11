package main

import "fmt"

func main() {
	//	byte型の話はしっかりやろうかな
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// http://www3.nit.ac.jp/~tamura/ex2/ascii.html
	// 文字列はasciiで表すことができるので、以下の文字列HIになる
	fmt.Println(string(byteA))

	// 文字列 to byte

	c := []byte("HI") // {72, 73}
	fmt.Println(c)

	fmt.Println(string(c))
}
