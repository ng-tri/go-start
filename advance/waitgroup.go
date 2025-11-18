package advance

import (
	"fmt"
	"sync"
	"time"
)

func RunWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		count1("sheep")
	}()
	wg.Wait()

	fmt.Println("Done")
}

func count1(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(thing, i)
		time.Sleep(500 * time.Millisecond)
	}
}
