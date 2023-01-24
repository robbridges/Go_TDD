package Select

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Happy path, slow channel should lose", func(t *testing.T) {
		slowServer := makeDelayedSever(20 * time.Millisecond)
		fastServer := makeDelayedSever(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Errorf("We got an error when we weren't meant to")
		}

		if got != want {
			t.Errorf("The fast URL should have won")
		}
	})
	t.Run("Returns an error if a server doesn't respond within set duration seconds", func(t *testing.T) {
		slowServer := makeDelayedSever(3 * time.Second)
		fastServer := makeDelayedSever(3 * time.Second)

		errorMessage := fmt.Sprintf("Timed out waiting for %s, and %s", slowServer.URL, fastServer.URL)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 2*time.Second)
		if err == nil {
			t.Errorf("This should have returned an error")
		}
		if err.Error() != errorMessage {
			t.Errorf("We got the wrong error message, we got %s we wanted, %s", errorMessage, err.Error())
		}
	})
}

func makeDelayedSever(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
