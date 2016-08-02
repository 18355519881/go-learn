package main

import "fmt"

func main() {
    function1()
    //	fmt.Println("vim-go")
}
func function1(){
    fmt.Printf("In function1 at the top\n")
    defer function2()
    fmt.Printf("In function1 at the bottom!\n")
}
func function2(){
    fmt.Printf("function2 :defered until the end of the calling function!")
}
