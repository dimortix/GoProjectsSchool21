package main

import "testing"

func TestCache_SetGetEvictionAndMoveToFront(t *testing.T) {
	c := NewCache[int](2)

	c.Set(1, 10)
	c.Set(2, 20)

	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("Get(1) expected (10, true), got (%v, %v)", v, ok)
	}

	c.Set(3, 30)

	if _, ok := c.Get(2); ok {
		t.Fatalf("expected key=2 to be evicted")
	}
	if v, ok := c.Get(1); !ok || v != 10 {
		t.Fatalf("expected key=1 to remain with value 10, got (%v, %v)", v, ok)
	}
	if v, ok := c.Get(3); !ok || v != 30 {
		t.Fatalf("expected key=3 to be present with value 30, got (%v, %v)", v, ok)
	}
}

func TestCache_Clear(t *testing.T) {
	c := NewCache[string](3)
	c.Set(1, "a")
	c.Set(2, "b")
	c.Clear()

	if _, ok := c.Get(1); ok {
		t.Fatalf("expected cache to be cleared")
	}
	if _, ok := c.Get(2); ok {
		t.Fatalf("expected cache to be cleared")
	}
}

func TestCache_ZeroCapacity(t *testing.T) {
	c := NewCache[int](0)
	c.Set(1, 10)
	if _, ok := c.Get(1); ok {
		t.Fatalf("expected nothing to be stored when capacity is 0")
	}
}
