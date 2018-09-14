package main

import(
	"testing"
	"time"
)

func Test10sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(10 * time.Second)
}

func Test20sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(20 * time.Second)
}

func Test30sSleep(t *testing.T) {
	t.Parallel()
	time.Sleep(30 * time.Second)
}
