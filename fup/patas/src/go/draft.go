package main
import "fmt"
func main() {
   var B, C, animais int

   fmt.Scan(&B, &C, &animais)

   animal := make([]string, animais)
   total := 0
   for i := 0; i < animais; i++{
      fmt.Scan(&animal[i])

      if animal[i] == "v" {
         total += 4
      } else if animal[i] == "g"{
         total += 2
      } else {
         total += 4
      }
   }

      fmt.Println(total)
      d1 := total - B
      d2 := total - C

      if d1 < 0 {
         d1 *= -1
      }
      if d2 < 0 {
         d2 *= -1
      }  
       
      if d2 > d1{
      fmt.Println("Chico Bento")
   } else if d1 > d2 {
      fmt.Println("Cebolinha")
   }else {
      fmt.Println("empate")
   }
}

