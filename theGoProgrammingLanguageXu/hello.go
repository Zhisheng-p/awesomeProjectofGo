package main

import "fmt"

func main()  {
	fmt.Println("hello world。 你好，世界！")
	fmt.Println("-----------------------")
	/*
	 * 2.3 类型
	 */
	var v1 bool
	v1 = true
	// v2 会被推导为 bool 类型
	v2 := (1 == 2)

	fmt.Println("v1 is: ", v1)
	fmt.Println("v2 is: ", v2)

	// bool 类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	var v3 bool
	v3 = (1 != 0)
	fmt.Println("v3 is: ", v3)

	// v5 将被自动推导为 int 类型，和 int32 是不同类型，需要强转
	var v4 int32
	v5 := 64
	v4 = int32(v5)
	fmt.Println("v4 is: ", v4)
	fmt.Println("v5 is: ", v5)

	v6, v7 := 1, 2
	if v6 == v7 {
		fmt.Println("v6 and v7 are equal.")
	} else {
		fmt.Println("v6 and v7 are not equal.")
	}

	// 不同类型的整型不能直接比较，但各种类型的整型变量都可以直接与字面常量（literal）进行比较
	var v8 int32
	var v9 int64
	v8, v9 = 1, 2
	// 编译不通过
	//if v8 == v9 {
	//	fmt.Println("v8 and v9 are equals.")
	//}
	// 编译通过
	if v8 == 1 || v9 == 2 {
		fmt.Println("v8 is 1, v9 is 2.")
	}


}
