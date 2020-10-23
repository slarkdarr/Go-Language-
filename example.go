// Program yang menampilkan "Hello, World!" di layar
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

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

// Program yang menggunakan loop
package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
/* Akan menghasilkan:
1
2
3
7
8
9
loop
1
3
5
*/
