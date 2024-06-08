package main

import (
	"fmt"
	"time"
)

func main() {
	var n int

	fmt.Println("Enter the number of elements to print (n):")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			time.Sleep(time.Second * 1) // Adjust time.Second for desired interval
			fmt.Println(i)
		}(i)
	}

	// Keep the main program running to prevent goroutines from exiting prematurely
	select {}
}
