package main

	
import (
    "fmt"
    "os"
	"strconv"
)

func Soma(x int, y int) int {
	return x + y
}

func main() {
	arg1, err := strconv.Atoi(os.Args[1])
	arg2, err := strconv.Atoi(os.Args[2])
	
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }	
	
	fmt.Println(Soma(arg1, arg2))
}
