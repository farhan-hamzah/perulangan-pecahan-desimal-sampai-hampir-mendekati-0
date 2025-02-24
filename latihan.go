package main
import (
    "fmt"
)
func main() {
    var num int
    result := 1.0 
    for {
        fmt.Scan(&num) 
        if num == 0 { 
            break
        }
        result *= 1.0 / float64(num)
    }
    fmt.Printf("%.5f\n", result)
}
