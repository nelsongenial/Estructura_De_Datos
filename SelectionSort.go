package main
import (
    "fmt"
    "math/rand"
	"time"
)

func cambiazo(a, b int) (int, int){
    return b, a
    }
    
func selection(arr []int) {
    n := len(arr) - 1
    if n <= 0 {
        return 
        }
    minimo :=0
    cont :=0
    for i:=0; i < n-1 ; i++{
        minimo = arr[i]
        cont = i
        for j:=i; j<=n; j++{
            if arr[j] < minimo{
                minimo = arr[j]
                cont =j
                }
            }
        arr[i], arr[cont] = cambiazo(arr[i], arr[cont])
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
    selection(x)
    fin := time.Now()
    fmt.Printf("selection se demoró %v", fin.Sub(ini))
    fmt.Println("")
    //fmt.Println(x)
    
    y := make([]int, N)
    for j:=0; j < N;j++{
        y[j] = rand.Int() % N
    }
    ini =time.Now()
    selection(y)
    fin = time.Now()
    fmt.Printf("selection se demoró %v", fin.Sub(ini))
    fmt.Println("")
    
    z := make([]int, N)
    for o:=0; o < N;o++{
        z[o] = rand.Int() % N
    }
    ini =time.Now()
    selection(z)
    fin = time.Now()
    fmt.Printf("selection se demoró %v", fin.Sub(ini))
    fmt.Println("")
    }