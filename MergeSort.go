package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

n1 := len(arr) / 2
a1 := make([]int, n1)
for i := 0; i < n1; i++ {
		a1[i] = arr[i]
	}
	n2 := len(arr) - n1
	a2 := make([]int, n2)
	
for i := 0; i < n2; i++{
		a2[i] = arr[i+n1]
	}

	MergeSort(a1)
	MergeSort(a2)
    m := Merge(a1, a2)
	
    for i := range m{
		arr[i] = m[i]
	}
}

func Merge(a1, a2 []int) []int{
	result := make([]int, len(a1)+len(a2))
	
pos1 := 0
pos2 := 0
pos := 0
	
for pos1 < len(a1) && pos2 < len(a2){
	if a1[pos1] < a2[pos2] {
		result[pos] = a1[pos1]
        pos1++
	}else{
		result[pos] = a2[pos2]
		pos2++
		}
	pos++
	}

	for pos1 < len(a1) {
		result[pos] = a1[pos1]
	    pos1++
	    pos++
	}
	
for pos2 < len(a2){
	result[pos] = a2[pos2]
	pos2++
		
pos++
	}
	
return result
}

func SelectionSort(arr []int) {
 	for i := 0; i < len(arr); i++ {
// 
		min := i
// 	
	for j := i + 1; j < len(arr); j++ {
// 	
		if arr[j] < arr[min] {
// 
				min = j
// 	
		}
 		}
// 		
arr[min], arr[i] = arr[i], arr[min]
// 
	}
}

func main() {
	N := 1000000
    x := make([]int, N)
    for i:=0; i < N;i++{
        x[i] = rand.Int() % N
    }
    //fmt.Println(x)
    ini :=time.Now()
    MergeSort(x)
    fin := time.Now()
    fmt.Printf("MergeSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
    //fmt.Println(x)
    
    y := make([]int, N)
    for j:=0; j < N;j++{
        y[j] = rand.Int() % N
    }
    ini =time.Now()
    MergeSort(y)
    fin = time.Now()
    fmt.Printf("MergeSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
    
    z := make([]int, N)
    for o:=0; o < N;o++{
        z[o] = rand.Int() % N
    }
    ini =time.Now()
    MergeSort(z)
    fin = time.Now()
    fmt.Printf("MergeSort se demoró %v", fin.Sub(ini))
    fmt.Println("")
}
