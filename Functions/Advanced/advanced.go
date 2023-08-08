package Advanced

import "fmt"
func Advanced(){
    add := func(a,b int) int{
        return a+b
    }
    fmt.Println(functionsasparams(add,4,5))
}

func functionsasparams(operation func(int,int)int,a,b int)int{
    return operation(a,b)
}

func Returnafunc(multiplier int) func(int)int {
    return func(i int) int {
        return i * multiplier
    } 
}

