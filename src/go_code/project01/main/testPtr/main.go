package main

import "fmt"

func main() {
   var c1 chan int = make(chan int, 1)
   var c2 chan int = make(chan int, 1)
   var c3 chan int = make(chan int, 1)
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received %v ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent %v", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received %v", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}