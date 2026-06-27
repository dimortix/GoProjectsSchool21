package main

import "testing"

func TestTopKFrequent_NormalCase(t *testing.T) {
	words := []string{"aa", "bb", "cc", "aa", "cc", "cc", "cc", "aa", "ab", "ac", "bb"}
	k := 3

	got := TopKFrequent(words, k)
	want := []string{"cc", "aa", "bb"}

	if len(got) != len(want) {
		t.Fatalf("expected %d words, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at position %d expected %q, got %q", i, want[i], got[i])
		}
	}
}

func TestTopKFrequent_EmptyWords(t *testing.T) {
	words := []string{}
	k := 3

	got := TopKFrequent(words, k)

	if len(got) != 0 {
		t.Fatalf("expected empty result for empty input, got %v", got)
	}
}

func TestTopKFrequent_KGreaterThanUniqueWords(t *testing.T) {
	words := []string{"aa", "bb", "aa"}
	k := 10

	got := TopKFrequent(words, k)
	want := []string{"aa", "bb"}

	if len(got) != len(want) {
		t.Fatalf("expected %d words, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at position %d expected %q, got %q", i, want[i], got[i])
		}
	}
}

