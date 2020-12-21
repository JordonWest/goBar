
package menus

import ("fmt"
        )

func MainMenu() {
  var selection string
  fmt.Println("GOBAR MENU...")
  fmt.Println("What would you like to do?")
  fmt.Println("1: See a full menu")
  fmt.Println("2: See our hours")
  fmt.Scanln(&selection)
  if selection == "1" {
    fmt.Println(BeerMenu())
  }else if selection == "2" {
    fmt.Println(Hours())
  }else{
    fmt.Println("Please select a valid choice")
    MainMenu()
  }
}


func Hours() string {
  return "PLACEHOLDER -- HOURS -- PLACEHOLDER"
  }

func BeerMenu() string {
  return "PLACEHOLDER -- BEER MENU -- PLACEHOLDER"
  }
