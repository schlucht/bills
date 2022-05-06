package main

import (
  "fmt"

  "github.com/schlucht/bills/model"
  )

func main() {
  b := Bill {
    Text: "Hallo",
    Amount: 12.3,
    Ballance: 1235.5896,
  }
  fmt.Println(b.ToString())
  fmt.Println("Hier ist die zweite Meldung")
}