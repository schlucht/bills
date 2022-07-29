package main

import (
	"fmt"
	"log"

	"github.com/schlucht/bills/src/model"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	bills, err := model.New("data/bisSeptember.csv")

	log.Println(err)
	fmt.Println(bills.String())
	mb := bills.FilterMonth()

	for key, value := range mb {
		fmt.Println(key, value.String())
	}
	fmt.Println("Hier ist die zweite Meldung")
}
