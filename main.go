// go run main.go
// go build main.go -> golang-brno.exe

package main

import "fmt"

func main() {

  // a := 6
  // var b int = 10

  // c := a + b

  slice := []int{1, 2, 3, 6, 12, 31, 45}

  for i := 0; i < len(slice); i++ {
    fmt.Println(slice[i])
  }

  fmt.Println(c)
}