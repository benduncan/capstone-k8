package main

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello world!"
	if got := SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
