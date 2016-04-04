package main
import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB]) // ""
	fmt.Printf("%+v\n", symbol)
	fmt.Println(symbol)
//	fmt.Println(iota)


}
