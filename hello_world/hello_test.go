package hello_world

import "testing"

func Test_hello(t *testing.T) {
	greeting := Hello_World("Rob")
	wanted := "hello Rob"
	if greeting != wanted {
		t.Errorf("Bad return value, we wanted %s, we got %s", wanted, greeting)
	}
}
