
package main

import (
  "fmt"
//  "github.com/goBar/data"
)

type Beer struct{
  name, brand string
  abv, rating int
}

func (beer Beer) beer_stats() {
  return Beer

func show_inv(list []string) {
  for _, beer := range list {
    fmt.Printf("Beer name: %v\n", beer)
  }
}
func welcome() string{
  return "welcome to the beer store!"
}

func main() {
  inventory := []string{"beer 1", "beer 2", "beer 3"}
  fmt.Println("Hello world!")
  show_inv(inventory)
  fmt.Println(welcome())
}

