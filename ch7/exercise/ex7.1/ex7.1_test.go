//@author: hdsfade
//@date: 2020-11-06-08:40
package ex7_1

import "testing"

func TestWordCounter_Write(t *testing.T) {
	p := []byte("Hello World!\nGo programming language")
	var w WordCounter
	want := 5
	got, err := w.Write(p)
	if err != nil {
		t.Fatalf("expected nonerr, but get one: %v", err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLineCounter_Write(t *testing.T) {
	p := []byte("Hello World!\nGo programming language")
	var w LineCounter
	want := 2
	got, err := w.Write(p)
	if err != nil {
		t.Fatalf("expected nonerr, but get one: %v", err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
