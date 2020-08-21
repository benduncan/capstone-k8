package main

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hola world!"
	if got := SayHello(); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
