// go run main.go
// go build main.go -> golang-brno.exe

package warmup

import "fmt"


type Workshop struct {
  Name      string
  Atendees  int
  Location  Location
}

type Location struct {
  Name  string
  Lat   float32
  Lon   float32
}

func (this *Location) GetName() string {
  return "---" + this.Name + "---"
}

func main() {

  loc := Location{"Brno", 40.13, 51.13}
  fmt.Println(loc.GetName())

  workshop := Workshop{"Workshop 101", 12, Location{"Brno", 31.1, 31.5}}
  workshop2 := Workshop{Name: "Golang 101"}

  fmt.Printf("Workshop: %v\n", workshop)
  fmt.Printf("Workshop2: %v\n", workshop2)

  // fmt.Println(slice)

  // for index, value := range slice {
  //   fmt.Println("slice[%d] = %d", index, value)
  // };

  // length, _ := printSlice(slice) // ignore the second value
  // fmt.Printf("Length is: %d", length)
}

func printSlice(slice []int) (int, string) { //declare two return values here
  for _, value := range slice {
    fmt.Printf("%d\n", value)
  }

  return len(slice), "This is slice"
}