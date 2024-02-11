package main

import "fmt"

func main() {
	// 配列型は他の言語とのちがいとしては、あとから要素数を変化することができない
	//	一般的な配列みたいなとことはsliceでやる

	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)
	// これを実行すると出てくるのが... [3]int つまり、これはintを3つ持つ配列という意味

	var arr2 [3]string = [3]string{"A", "B", "C"}
	//var arr2a [3]string = [3]string{"A", "B"} これは3番目に空文字が入る
	// これはだめです
	//var arr2 [3]string = [3]string{"A", "B", "C", "D"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 自動で要素数がわかる
	// これはなんか使えそう
	arr4 := [...]string{"C", "D"}
	fmt.Printf("%T\n", arr4)

	// 要素数の取り出し
	// index番号を指定
	fmt.Println(arr2[2])

	//
}
