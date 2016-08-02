package main

import "fmt"

func main() {
    x := min(1,2,3,4,0)
    fmt.Printf("The minimum is : %d \n",x)
    arr := []int{9,7,5,4}
    x = min(arr...)
    fmt.Printf("The minimum in the array is : %d \n",x)
    //	fmt.Println("vim-go")
}
func min(a ...int) int{
    if(len(a) == 0){
        return 0
    }
    min := a[0]
    for _, v := range a{
        if(v < min){
            min = v
        }
    }
    return min
}
