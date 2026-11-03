package piscine

import (
	piscine "command-line-arguments/home/bocal/revision/retainfirsthalf.go"
	"fmt"
)

func RetainFirstHalf(str string) string {
	if len(str) <= 1 {
		return str
	}
	return str[:len(str)/2]
}

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}
