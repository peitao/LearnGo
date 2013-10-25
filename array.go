package main

import "fmt"

func change( a []int ){
	a[0] = 100
}
func change_fix( a[6]int ){
	a[0] = 123
}
func main(){

	var myArray []int = []int{1,2,3,4,5,6}
	var myArray_fix_len [6]int = [6]int{1,2,3,4,5,6}


	change ( myArray )
	change_fix(myArray_fix_len)
	fmt.Println("changing []")
	for _, v:= range myArray{

		fmt.Print( v, " " )
	}
	fmt.Println("\nchanging [6]")
	for _,v:= range myArray_fix_len{

		fmt.Print( v, " ")
	}
	fmt.Println()
}

