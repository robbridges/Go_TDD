package Select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(slow, fast string) (winner string, error error) {
	select {
	case <-ping(slow):
		return slow, nil
	case <-ping(fast):
		return fast, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("Timed out waiting for %s, and %s", slow, fast)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
