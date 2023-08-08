package main
import(
    "fmt"
    "github.com/Nathan-Pokharel/Learning_functions/Anonymous"
    "github.com/Nathan-Pokharel/Learning_functions/Advanced"
)

func main(){
    simple(3)
    fmt.Println(Anonymous.Anonymous())
    Advanced.Advanced()
    a := Advanced.Returnafunc(2)
    fmt.Println(a(3))
    fmt.Println(variadicfunc(1,2,3,4,5,6,7,7,8,8,8))

}
func simple(a int){
    fmt.Println("Simple Function Call with parameters")
    fmt.Println(a)
}
func variadicfunc(numbers ...int)int {
    total := 0
    for _,num := range numbers{
        total += num
    }
    return total
}
