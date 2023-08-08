package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   [] Go Routines
   [] Channels
   [] Buffer
   [] Select
   [] Mutex
*/

func main(){
    waitgroups()

}

func goroutines(){
    go dosomething()
    for{
        time.Sleep(time.Second)
        fmt.Println("Also chilling while doing super heavy task")
    }
}

func dosomething(){
    for{
        fmt.Println("Doing Super Heavy Task ahhhh!")
        time.Sleep(time.Second)
    }
}

func channels(){
    ch := make(chan int)
    go func(){
        ch <- 42
    }()
    value := <-ch
    fmt.Println(value)
}

func bufferedChannels(){
    ch := make(chan int,3)
    go func(){
        ch<-1 
        ch<-2
        ch<-3
        ch<-4
        ch<-5
        ch<-6
        ch<-7
    }()
    
}
func waitgroups() {
	var wg sync.WaitGroup
	wg.Add(3)
	go testwaitgroup(&wg)
	go testwaitgroup1(&wg)
	go testwaitgroup2(&wg)
	wg.Wait() // Wait for all goroutines to finish
}

func testwaitgroup(wg *sync.WaitGroup) {
    var z time.Duration = 1
    defer wg.Done()
    time.Sleep(time.Second * z)
    fmt.Println("Finished")
}
func testwaitgroup1(wg *sync.WaitGroup) {
    var z time.Duration = 3
    defer wg.Done()
    time.Sleep(time.Second * z)
    fmt.Println("Finished")
}
func testwaitgroup2(wg *sync.WaitGroup) {
    var z time.Duration = 5
    defer wg.Done()
    time.Sleep(time.Second * z)
    fmt.Println("Finished")
}

