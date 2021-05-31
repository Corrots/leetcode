package main

import "fmt"

func main() {
	//fmt.Println(Anonymous()) // defer1 in value is 1, defer2 value is 2, 0
	//fmt.Println(HasName())   //
	fmt.Println(Test1())
	fmt.Println(Test2())
	fmt.Println(Test3())
}

func Anonymous() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2 value is ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1 in value is ", i)
	}()

	return i
}

func HasName() (j int) {
	defer func() {
		j++
		fmt.Println("defer2 in value", j)
	}()

	defer func() {
		j++
		fmt.Println("defer1 in value", j)
	}()

	return j
}

func Test1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}

func Test2() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)
	return 2
}

func Test3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}
