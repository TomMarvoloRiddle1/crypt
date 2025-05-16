package main

import (
	"testing"
)

func TestTempReadKey(t *testing.T) {
	key := string(tempReadKey())

	check := string([]byte{20, 10, 49, 17, 220, 71, 39, 226, 124, 182, 205, 91, 170, 88, 193, 222, 76, 153, 40, 123, 103, 59, 148, 214, 215, 51, 57, 213, 126, 65, 33, 71})
	if key != check {
		t.Errorf("%s", "wrong")
	}
}
