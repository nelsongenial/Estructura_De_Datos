package main

import (
    "fmt"
    "strings"
)
    
type tnodo struct{
    valor string
    derecha *tnodo
    izquierda *tnodo
    }

type tree struct{
    head *tnodo
    }
    
func nuevoTree() *tree{
    return new(tree)
    }



func (t *tnodo) insertarOperador(ope string, n *int){
    if t == nil{
        return
        }
    if t.derecha == nil{
            if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
                t.derecha = &tnodo{
                    valor: ope,
                    derecha: nil,
                    izquierda: nil,
                }
                *n = 0
                return
                }
        }
    if t.izquierda == nil{
            if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
                t.izquierda = &tnodo{
                    valor: ope,
                    derecha: nil,
                    izquierda: nil,
                }
                *n = 0
                return
                }
        }
    t.derecha.insertarOperador(ope, n)
    if *n == 0{return}
    t.izquierda.insertarOperador(ope, n)
    if *n == 0{return}
    }

func (t *tnodo) insertarValor(ope string, n *int){
    if t == nil{
        return
        }
    t.derecha.insertarValor(ope, n)
    if *n == 0{return}
    if t.derecha == nil{
            if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
                t.derecha = &tnodo{
                    valor: ope,
                    derecha: nil,
                    izquierda: nil,
                }
                *n = 0
                return
                }
        }
    if t.izquierda == nil{
            if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
                t.izquierda = &tnodo{
                    valor: ope,
                    derecha: nil,
                    izquierda: nil,
                }
                *n = 0
                return
                }
        }
    //t.derecha.insertarValor(ope, n)
    t.izquierda.insertarValor(ope, n)
    if *n == 0{return}
    }

func (t *tree) crearTree(sumas []string, n int){
    //t= nuevoTree()
    contador := 1
    //var aux []string
    if t.head == nil{
            t.head = &tnodo{
                valor: sumas[n-1], //len(sumas-1)
                derecha: nil,
                izquierda: nil,
                }
            }
    for i:=0; i < len(sumas)-1;i++{
        switch sumas[n-2-i]{//len(sumas-i-2))
            case string('/'),string('+'),string('*'),string('-'):
                contador = 1
                t.head.insertarOperador(sumas[n-i-2], &contador)
            default:
                contador = 1
                t.head.insertarValor(sumas[n-i-2], &contador) //aux = append(aux,sumas[n-2-i])//t.head.insertarOperador(sumas[n-2-i])
            }
        }
    }
    
func (t *tnodo) imprimir(){
    if t == nil{
        return
        }
    if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
        fmt.Print("(")
        }
    t.izquierda.imprimir()
    fmt.Print(t.valor)
    t.derecha.imprimir()
    if t.valor == string('/') || t.valor == string('+') || t.valor == string('*') || t.valor == string('-'){
        fmt.Print(")")
        }
    }
    
func main(){
    var suma string
    fmt.Println("Ingrese una suma Polaca:")
    fmt.Scan(&suma)
    //fmt.Println(suma)
    sumas := strings.Split(suma, "")
    n := len(sumas)
    //fmt.Println(n)
    t := nuevoTree()
    t.crearTree(sumas,n)
    //fmt.Println("uguu, aqui empieza")
    t.head.imprimir()
    fmt.Println("  ")
    fmt.Println("Se dibujo exitosamente")
    }