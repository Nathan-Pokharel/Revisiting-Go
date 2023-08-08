package main
import "fmt"
/*
    [*] Arrays
    [*] Slices
    [*] Maps
    [*] Make()
*/
func main(){
    Slices()
}

func Arrays(){
    //var Array [5]int
    //Array = [5]int{1,2,3,4,5}
    var array = [5]int{1,2,3,4,5}
    fmt.Println(array[2])
    fmt.Println(len(array))
    for index,value := range array{
        fmt.Println(index,value)
    }
    // Arrays Are Values
    var arr1 = [2]int{1,2}
    var arr2 = arr1 
    arr2[1] = 3
    fmt.Println(arr2,arr1)
    arrayasparams([5]int{1,2,3,4,5})
    
    var matrix = [3][3]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println(matrix)
    for i:=0;i<3;i++{
        for j:=0;j<3;j++{
            fmt.Println(matrix[i][j])
        }
    }
}

func ArraysWithEllipses(){
    fmt.Println("---- ArraysWithEllipses --- ")
    numbers := [...]int{1,2,3,4,5}
    fmt.Println(len(numbers))
    fmt.Println("---- --- ")
}

func arrayasparams(array [5]int){
    fmt.Println("---- Arrays As Params --- ")
    fmt.Println(array)
    fmt.Println("---- --- ")
    ArraysWithEllipses()
}

func Slices(){
    //var slice []int
    var numbers = []int{1,2,4,5,6}
    fmt.Println(len(numbers),cap(numbers))
    numbers = append(numbers, 1,2,3,4)
    numbers = append(numbers,[]int{5,6,7,8,9,10}...)
    fmt.Println(numbers)
    numbers = []int{2,4,7,1,2,4,7}
    fmt.Println(numbers[2:],numbers[:3])
    // Removing Elements with Append
    numbers=append(numbers[:2],numbers[3:]...)
   // Arrays Create Copies But Slices Dont unless copy() is used.

}

func Maps(){
    age := make(map[string]int) // make(map[KeyType]ValueType)
    age["Alice"]=20
    age["Bob"]=30
    fmt.Println(age["Bob"])
    fmt.Println(age["Candice"]) // Returns 0 since Candice doesnt exist
    delete(age,"Bob")
    //Check the Existence of a key ,
    value,exists:=age["Alice"]
    if exists{
        fmt.Println(value)
    }else{
        fmt.Println("Key Dont Be existing My Guy")
    }
    for key,value := range age{
        fmt.Println(key,value)
    }
}

func Make(){
    makeslice := make([]int,5,10) //make a slice of type []int with length 5 and capacity 10
    makemaps := make(map[string]int) // make a map with keyType string and value type int
    makechannels := make(chan int) // make a channel of type int 
    fmt.Println(makeslice,makemaps,makechannels)
}
