package main

import (
	"fmt"
  "bufio"
  "os"
  "strings"
)

func main() { 
  min, max := 1, 100
  guess := max / 2
  posibilities := 100
	reader := bufio.NewReader(os.Stdin)

  fmt.Println("Guessing a number from 1 to 100")
  fmt.Println("---------------------")

  for posibilities > 2 {
    fmt.Println("Is your number greater than", guess, "? (y / n)")
    fmt.Print("-> ")

    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if strings.TrimRight(text, "\n") == "y" {
      min = guess + 1 // 51
    } else if strings.TrimRight(text, "\n") == "n" {
      max = guess // 50
    } else {
      fmt.Println("use y or n, idiot")
      
      continue
    }

    guess = max - ((max - min) / 2)
    posibilities = (max - min) + 1

    fmt.Println("min", min)
    fmt.Println("max", max)
    fmt.Println("guess", guess)
    fmt.Println("pos", posibilities)
  }

  // only 2 posibilities
  fmt.Println("Is your number ", min, "? (y / n)")

  text, _ := reader.ReadString('\n')
  // convert CRLF to LF
  text = strings.Replace(text, "\n", "", -1)

  if strings.TrimRight(text, "\n") == "y" {
    fmt.Println("Thanks, bye.")
  } else {
    fmt.Println("Your number is", min + 1, ". gtfo")
  }
}
