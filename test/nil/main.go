package main

import "fmt"

func main() {
	//list := map[string]struct{}{}
	//list := make(map[string]struct{})
	var list map[string]struct{} // map未正常初始化，写数据会panic

	fmt.Println(list["cc"]) // 可读数据

	list["abc"] = struct{}{}

	fmt.Println(list)
}
