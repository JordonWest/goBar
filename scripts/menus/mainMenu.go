
package menus

import ("fmt"
        "time"
        )
func Hours() string {
  day := (time.Now().Weekday()).String()
  hours := map[string]string{"Sunday":"closed", "Saturday":"1000 to 2200"}
  otherDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
  for _, day := range otherDays {
    hours[day] = "1000 to 1600"
  }
  return "Our hours today are: " + hours[day]
}

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


func BeerMenu() string {
  
  return "PLACEHOLDER -- BEER MENU -- PLACEHOLDER"
  }

