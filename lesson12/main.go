package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)

	//【uint型(+の整数）、complex(複素数型)】
	//他にもuint(+の整数) complex(複素数型）などの型があるが、あまり使用頻度は無いので今回は紹介だけにします。
	//
	//var u8 uint = 255   //uint型
	//
	//var c64 complex64 = -5 + 12i //complex型
}
