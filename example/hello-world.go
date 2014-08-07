package main

import "fmt"
import "time"
import "github.com/r8k/ticker"

func print(args []interface{}) {
	fmt.Printf("Hello World !! Got %d args\n", len(args))
	for i, s := range args {
		fmt.Println(i, s)
	}
}

func main() {
	stop := ticker.TickEvery(print, time.Second*3, []interface{}{"a", 2})
	time.Sleep(time.Second * 10) // sleep for 10 seconds, to see 3 print invocations
	stop <- true                 // send a stop signal
	time.Sleep(time.Second * 2)  // sleep for 2 more seconds, to ensure, it stopped :)
}
