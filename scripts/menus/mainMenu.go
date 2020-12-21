
package menus

import ("fmt"
        "github.com/goBar/scripts/menus"
        )

func MainMenu() {
  var selection string
  fmt.Println("GOBAR MENU...")
  fmt.Println("What would you like to do?")
  fmt.Println("1: See a full menu")
  fmt.Println("2: See our hours")
  fmt.Scanln(&selection)
  if selection == "1" {
    menus.BeerMenu()
  }else if selection == "2" {
    menus.Hours()
  }else{
    fmt.Println("Please select a valid choice")
    MainMenu()
  }
}
