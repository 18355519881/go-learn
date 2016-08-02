package main

import "fmt"

func main() {
    var a int = 30
    var b int = 20
    printValues(cal1(a,b))
    printValues(cal2(a,b))

//	fmt.Println("vim-go")
}
func printValues(sum int,mult int,sub int){
    fmt.Printf("sum = %d , mult = %d ,sum = %d \n",sum,mult,sub)
}
func cal1(a int,b int) (int,int,int){
    return a+b,a*b,a-b
}
func cal2(a int,b int)(sum int,mult int,sub int){
    sum = a + b
    mult = a * b
    sub = a - b
    return 
}
