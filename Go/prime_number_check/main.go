package main
import (
    "fmt"
    "math"
)

func isPrime(m int) bool {
    if m <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(m))); i++ {
        if m%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var n, m int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&m)
        if isPrime(m) {
            fmt.Println("Prime")
        } else {
            fmt.Println("Not prime")
        }
    }
}
