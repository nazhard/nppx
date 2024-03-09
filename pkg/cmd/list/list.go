package list

import (
	"fmt"
	"sync"
)

func Print(args []string) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			fmt.Println(a)
		}(arg)
	}

	wg.Wait()
}
