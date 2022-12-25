package b

import "fmt"

func main() {
	var a = true
	defer func() {
		fmt.Println("1") //程序运行时经过该defer，所以会在程序结束后输出
	}()
	if a {
		fmt.Println("2")
		return //程序运行到这里就结束了，所以不会运行下面的defer
	}
	defer func() {
		fmt.Println("3") // 因为程序已经return，故不会运行该defer
	}()
}
