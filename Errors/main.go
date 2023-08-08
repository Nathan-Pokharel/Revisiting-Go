package main
import (
	"errors"
	"fmt"
)

/*
   [*] Errors
   [*] Panics
   [*] Defer
   [*] Recover
*/
func main(){
    result,err := erroz(1,0)
    if err != nil{
            fmt.Println("Error:",err)
    }
    fmt.Println(result)
    Defer()
}


func Defer(){
    defer Recover()
    fmt.Println("Its gon' panic after this statement")
    panic("Something Went Off the Charts Wrong")

}

func Recover(){
    if r:=recover();r !=nil{
        fmt.Println("Recovered From Panic")
    }

}

func erroz(a,b int) (int,error){
    if b ==0 {
        return 0,errors.New("Division By 0")
    }
    return a / b,nil

}
