package main

import(
	"fmt"
)


func boring(msg string) <-chan string{
	c:= make(chan string)
	go func(){
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			// time.Sleep(sleepLen)
		}
	}()
return c
}


// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c:= make(chan string)
// 	go func() {for { c <- <-input1}}()
// 	go func() {for { c <- <-input2} }()
// 	return c
// }

func fanIn(input1, input2 <-chan string) <-chan string {
	c:= make(chan string)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <-s
			}
		}
	}()
	return c
}
func main(){
	// joe:= boring("Joe")
	// ann:= boring("Ann")
	c:= fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i <10; i ++ {
		// fmt.Println(<-joe)
		// fmt.Println(<-ann)
		fmt.Println(<-c)
	}

	fmt.Println("Your're both boring; I'm leaving.")
}