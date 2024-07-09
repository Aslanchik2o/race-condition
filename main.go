package main
//состояние гонки 

import (
	"fmt"
	"time"
)

func main() {
	countr := 0
	for i := 0; i < 1000; i++ {
		go func() {
			countr++
		}()
	}
	time.Sleep(time.Second)

	fmt.Println(countr)
}