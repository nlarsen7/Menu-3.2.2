//Nicholas Larsen
//4/7/2020
//a simple menu of food items
package main

import "fmt"

func menu() {
  fmt.Println("scallops")
  fmt.Println("steak")
  fmt.Println("salmon")
  fmt.Println("hamburger")
  fmt.Println("salad")
  fmt.Println("chicken")
  fmt.Println("pizza")
  //the menu
  
}
func main() {
menu()
var user string
fmt.Println("choose what you want from the menu")
fmt.Scanln(&user)
fmt.Println("you have chosen",user,"as your meal")
//prints out the menu and has the user choose frim the items
}