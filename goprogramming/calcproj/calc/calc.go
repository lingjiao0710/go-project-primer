package main

import (
	"fmt"
	"github.com/lingjiao0710/test/goprogramming/calcproj/simplemath"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("使用方法: calc command [arguments] ...")
	fmt.Println("命令是: ")
	fmt.Println("add 两数相加")
	fmt.Println("sqrt 平方根")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	fmt.Println(args, len(args))

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("使用方法: calc add int1 int2")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("使用方法: calc add int1 int2")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("结果: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("使用方法: calc sqrt int")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("使用方法: calc sqrt int")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("结果: ", ret)
	default:
		Usage()
	}
}
