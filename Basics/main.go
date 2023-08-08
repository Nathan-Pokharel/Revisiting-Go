package main

import "fmt"

func main(){
    // Conditionals()
    Data_Types()
}

/* Data Types */
func Data_Types(){
    boolean := true // Bool 
    var unsigned_int uint64 = 3
    var signed_int uint64=3
    const PI = 3.141592653
    bytes := []byte("This Is A Byte")
    bytetostring := string([]byte{102,97,108,99,111,110})
    fmt.Println(boolean,bytes,bytetostring,unsigned_int,signed_int,PI) 
}

/* Conditionals */
func Conditionals(){
    z := "UGABUGA"
    if z == "UGABUGA"{
        fmt.Println("True")

    }else if z == "BUGAUGA"{
        fmt.Println("False")

    }else{
        fmt.Println("False")
    }
    toswitch := 1 
    switch toswitch {
    case 3:
        fmt.Println("This is 3")

    case 2:
        fmt.Println("This is 2")

    default:
        fmt.Println("UGA BUGA")
    }
}
