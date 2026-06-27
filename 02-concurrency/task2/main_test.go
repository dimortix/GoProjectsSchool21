package main

import "testing"

func collect(ch <-chan int) []int {
	var res []int
	for v := range ch {
		res = append(res, v)
	}
	return res
}

func TestGeneratorAscending(t *testing.T) {
	got := collect(generator(2, 5))
	want := []int{2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("len: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}

func TestSquarePipeline(t *testing.T) {
	got := collect(square(generator(3, 5)))
	want := []int{9, 16, 25}
	if len(got) != len(want) {
		t.Fatalf("len: want %d, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}
