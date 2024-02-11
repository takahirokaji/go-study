package main

import "fmt"

// 暗黙的な変数は関数外で定義できない
// x i5 := 150
// o var i5 = 160

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)

	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	//s3 = "Go"
	fmt.Println(i3, s3)

	i4 := 400
	fmt.Println(i4)

	i4 := 450
	fmt.Println()

	i4 = "Hello"
	//	compile erroe

}

// ここまではコメントアウトしつつ、ファイルを作成してきたけれど、次からは、makefileで自動生成するので急に話が飛ぶかもしれません
