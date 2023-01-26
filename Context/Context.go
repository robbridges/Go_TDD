package Context

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
	Response string
	t        *testing.T
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprintf(w, data)
	}
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.Response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

//func (s *SpyStore) Cancel() {
//	s.Cancelled = true
//}
//
//func (s *SpyStore) AssertWasCancelled() {
//	s.t.Helper()
//	if !s.Cancelled {
//		s.t.Errorf("store was not cancelled")
//	}
//}
//
//func (s *SpyStore) AssertWasNotCancelled() {
//	s.t.Helper()
//	if s.Cancelled {
//		s.t.Errorf("store was cancelled when it should not have been")
//	}
//}
