package main
import (
    "fmt"
    "math/rand"
	"time"
)

func cambiazo(a, b int) (int, int){
    return b, a
    }

func insertion(arr []int){
    n := len(arr) - 1
    if n <= 0 {
        return 
        }
    for i:=1;i<=n;i++{
        for j:=i; j > 0; j-- {
            if arr[j] < arr[j-1]{
                arr[j], arr[j-1] = cambiazo(arr[j], arr[j-1])
                }
            }
        }
    }
    
func main(){
    N := 1000000
    x := make([]int, N)
    for i:=0; i < N;i++{
        x[i] = rand.Int() % N
    }
    //fmt.Println(x)
    ini :=time.Now()
    insertion(x)
    fin := time.Now()
    fmt.Printf("insertion se demoró %v", fin.Sub(ini))
    fmt.Println("")
    //fmt.Println(x)
    
    y := make([]int, N)
    for j:=0; j < N;j++{
        y[j] = rand.Int() % N
    }
    ini =time.Now()
    insertion(y)
    fin = time.Now()
    fmt.Printf("insertion se demoró %v", fin.Sub(ini))
    fmt.Println("")
    
    z := make([]int, N)
    for o:=0; o < N;o++{
        z[o] = rand.Int() % N
    }
    ini =time.Now()
    insertion(z)
    fin = time.Now()
    fmt.Printf("insertion se demoró %v", fin.Sub(ini))
    fmt.Println("")
    }