package main
import "fmt"
func main() {
   var nome string
   var class int

   fmt.Scan(&nome, &class)

   if class < 12 {
    fmt.Println(nome, "eh crianca")
   } else if class >= 12 && class < 18{
    fmt.Println(nome, "eh jovem")
   } else if class >= 18 && class < 65{
    fmt.Println(nome, "eh adulto")
   } else if class >= 65 && class < 1000 {
    fmt.Println(nome, "eh idoso")
   } else {
    fmt.Println(nome, "eh mumia")
   }
}
