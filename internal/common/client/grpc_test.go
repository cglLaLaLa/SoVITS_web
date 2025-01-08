package client

import (
	"fmt"
	"testing"
)

func TestWaitFor(t *testing.T) {
	fmt.Print(waitFor("127.0.0.1:12888", 10000000000))
}
