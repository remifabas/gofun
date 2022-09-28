package goroutines

import (
	"fmt"
	"time"
)

func Laroutourne() {
	for i := 0; i < 5; i++ {
		fmt.Println("laroutourne encore")
		time.Sleep(1000 * time.Millisecond)
	}
}
