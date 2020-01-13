package main
import (
    "fmt"
    "math/rand"
	"time"
)

func cambiazo(a, b int) (int, int){
    return b, a
    }

func bubble(arr []int){
    n := len(arr) - 1
    if n <= 0 {
        return 
        }
    for j:=0; j < n; j++{
        for i:=0; i < n; i++ {
            if arr[i] > arr[i+1]{
                arr[i], arr[i+1] = cambiazo(arr[i], arr[i+1])
                }
        }
    }
    return 
    }
   
   //time.Now()
   //rand.Int() % N
    
func main(){
    N := 1000000
    x := make([]int, N)
    for i:=0; i < N;i++{
        x[i] = rand.Int() % N
    }
    //fmt.Println(x)
    ini :=time.Now()
    bubble(x)
    fin := time.Now()
    fmt.Printf("Bubble se demoró %v", fin.Sub(ini))
    fmt.Println("")
    //fmt.Println(x)
    
    y := make([]int, N)
    for j:=0; j < N;j++{
        y[j] = rand.Int() % N
    }
    ini =time.Now()
    bubble(y)
    fin = time.Now()
    fmt.Printf("Bubble se demoró %v", fin.Sub(ini))
    fmt.Println("")
    
    z := make([]int, N)
    for o:=0; o < N;o++{
        z[o] = rand.Int() % N
    }
    ini =time.Now()
    bubble(z)
    fin = time.Now()
    fmt.Printf("Bubble se demoró %v", fin.Sub(ini))
    fmt.Println("")
    }