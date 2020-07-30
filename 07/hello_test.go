package hello_test

import (
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi there yourself!"
	got := hello.ReturnGreeting("Hi there")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	want2 := "How are you yourself!"
	got2 := hello.ReturnGreeting("How are you")
	if want2 != got2 {
		t.Errorf("want %q, got %q", want2, got2)
	}
}
