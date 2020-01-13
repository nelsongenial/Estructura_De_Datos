package main

import (
    "fmt"
    "math/rand"
	"time"
)

type tNodo struct{
    valor int
    next *tNodo
    }
    
type tLista struct{
    head *tNodo
    tail *tNodo
    }

func NuevaLista() *tLista{
    return new(tLista)
    }

func (t *tLista) agregar (x int){
    if t.head == nil{
        t.head = &tNodo{
            valor: x,
            next: nil,
            }
        t.tail = t.head
        return
        }
    t.tail.next = &tNodo{
        valor: x,
        next: nil,
        }
    t.tail = t.tail.next
    return
    }

//Por el momento ListSort no retorna nada porque como no se usaria no podria compilarlo
func ListSort(arr []int) {
    QuickSort(arr)
    r := NuevaLista()
    for i:=0; i < len(arr);i++{
        r.agregar(arr[i])
        }
    return 
    }

func med3(a []int) int {
	p1 := 0 //rand.Int() % len(a) // 0
	p2 := len(a)/2 //rand.Int() % len(a) // len(a) / 2
	p3 := len(a)-1 //rand.Int() % len(a) // len(a) - 1
	if a[p1] < a[p2]{
	if a[p2] < a[p3]{
			return p2
		}

	if a[p3] < a[p1]{
   return p1
		}
		
   return p3
	}else{
  if a[p1] < a[p3]{
  return p1
		}

  if a[p3] < a[p2]{
 return p2
	}

  return p3
	}
}


func pivotear(a []int) int{
  p := med3(a)
	a[0], a[p] = a[p], a[0]
	areaMenores := 1

  areaMayores := len(a) - 1
	
for areaMenores <= areaMayores {
		
if a[areaMenores] < a[0] {
			
areaMenores++
		} else{
  a[areaMayores], a[areaMenores] = a[areaMenores], a[areaMayores]

   areaMayores--
	}
	}
	
a[0], a[areaMayores] = a[areaMayores], a[0]
	
return areaMayores
}

func QuickSort(arr []int) {
	if len(arr) <= 1{
		return
	}
	
p := pivotear(arr)
	
if p > 0 {
		
QuickSort(arr[0 : p-1])
	}
	
QuickSort(arr[p+1:])
}


func main(){
    N := 100000*10
    x := make([]int, N)
    for i:=0; i < N;i++{
        x[i] = rand.Int() % N
    }
    //fmt.Println(x)
    ini :=time.Now()
    ListSort(x)
    fin := time.Now()
    fmt.Printf("ListSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
    //fmt.Println(X)
    
    y := make([]int, N)
    for j:=0; j < N;j++{
        y[j] = rand.Int() % N
    }
    ini =time.Now()
    ListSort(y)
    fin = time.Now()
    fmt.Printf("ListSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
    
    z := make([]int, N)
    for o:=0; o < N;o++{
        z[o] = rand.Int() % N
    }
    ini =time.Now()
    ListSort(z)
    fin = time.Now()
    fmt.Printf("ListSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
}
