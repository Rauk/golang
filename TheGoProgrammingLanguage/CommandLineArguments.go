package main
import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args), os.Args[0])
}
