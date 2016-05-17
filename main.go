// go run main.go
// go build main.go -> golang-brno.exe

package main

import "fmt"

func main() {

  slice := []int{1, 2, 3, 6, 12, 31, 45}

  fmt.Println(slice)

  for index, value := range slice {
    fmt.Println("slice[%d] = %d", index, value)
  };

  printSlice(slice)
}

func printSlice(slice []int) {
  for _, value := range slice {
    fmt.Printf("%dn", value)
  }
}