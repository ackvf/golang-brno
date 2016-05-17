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

  length, _ := printSlice(slice) // ignore the second value
  fmt.Printf("Length is: %d", length)
}

func printSlice(slice []int) (int, string) { //declare two return values here
  for _, value := range slice {
    fmt.Printf("%d\n", value)
  }

  return len(slice), "This is slice"
}