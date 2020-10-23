// Program yang menggunakan variabel dan kondisional
package main

import "fmt"

func main() {
    var a = 2
    var b = 4
    if a%2 == 0 {
        fmt.Println(a, "genap")
    } else {
        fmt.Println(a, "ganjil")
    }
    
    if b%a == 0 {
        fmt.Println(b + " dapat dibagi " + a)
    }

    if num := 9; num < 0 {
        fmt.Println(num, "negatif")
    } else if num < 10 {
        fmt.Println(num, "memiliki 1 digit")
    } else {
        fmt.Println(num, "has banyak digit")
    }
}
/* Akan menghasilkan:
2 genap
4 dapat dibagi 2
9 memiliki 1 digit
*/
