package main

import (
 "fmt"
 "os"
)

func main() {
    var goos string = os.Getenv("GOOS");
    fmt.Printf("The operating system is : %s\n",goos);
    path := os.Getenv("PATH");
    fmt.Printf("PATH is %s\n",path);
	fmt.Println("vim-go")
}
