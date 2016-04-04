package main
import "fmt"

type obj struct{
	member1 int64
	member2 string
}
func main() {
	obj1 := obj{2, "sd"}
	fmt.Printf("%+v\n", obj1)

}
