package main
import "fmt"
func main() {
  var b,t int

  fmt.Scan(&b,&t)

  met := (160 * 70) / 2
  areaF := ((b + t)/2) * 70

  if areaF > met {
    fmt.Println("1")
  }else if areaF < met {
    fmt.Println("2")
  } else {
    fmt.Println("0")
  }

}
