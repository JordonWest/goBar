
package main

import (
  "fmt"
//  "github.com/goBar/data"
)

type Beer struct{
  name, brand string
  abv, rating, stock int
}

func (beer *Beer) take_one() {
  beer.stock = beer.stock - 1
}

func show_inv(list []string) {
  for _, beer := range list {
    fmt.Printf("Beer name: %v\n", beer)
  }
}
func welcome() string{
  return "welcome to the beer store!"
}

func main() {
  //inventory := []string{"beer 1", "beer 2", "beer 3"}
  corova := Beer{"Batman", "Corova", 3, 5, 100}
  fmt.Println(corova)
  corova.take_one()
  fmt.Println(corova)
}

