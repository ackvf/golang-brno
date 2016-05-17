// go build main.go -> golang-brno.exe
package main

import "fmt"

func main() {

  a := 6
  var b int = 10

  c := a + b

  if c == 15 {
    fmt.Println("yes it's 15")
  } else {
    fmt.Println("no it's not")
  }

  fmt.Println(c)
}