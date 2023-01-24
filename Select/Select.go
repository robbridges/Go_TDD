package Select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(slow, fast string) (winner string, error error) {
	return ConfigurableRacer(slow, fast, tenSecondTimeout)
}

func ConfigurableRacer(slow, fast string, duration time.Duration) (winner string, error error) {
	select {
	case <-ping(slow):
		return slow, nil
	case <-ping(fast):
		return fast, nil
	case <-time.After(duration):
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
