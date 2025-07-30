package store

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	s.Set("A", "B", "C")
}

func TestGet(t *testing.T) {
	s := New()
	s.Set("A", "B", "C")
	value, err := s.Get("A", "B")

	if err != nil {
		t.Fatalf("%q", err)
	}

	if value != "C" {
		t.Fatalf("incorrect value was returned")
	}
}

func TestGetBad(t *testing.T) {
	s := New()
	_, err := s.Get("A", "B")

	if err == nil {
		t.Fatalf("somehow got value before setting")
	}

	s.Set("A", "D", "C")
	_, err = s.Get("A", "B")
	if err == nil {
		t.Fatalf("somehow got uninitialized field")
	}
}

func TestDel(t *testing.T) {
	s := New()
	s.Set("A", "B", "C")
	s.Del("A", "B")

	_, err := s.Get("A", "B")
	if err == nil {
		t.Fatalf("item should not exist after deletion")
	}
}

func TestDelBeforeSet(t *testing.T) {
	s := New()

	// First part tests before setting any value
	deleted, err := s.Del("A", "B")

	if deleted {
		t.Fatalf("somehow deleted non-existent field")
	}

	if err == nil {
		t.Fatalf("error was returned incorrectly")
	}

	// Second part examines deleting an uninitialized key
	s.Set("A", "B", "C")
	deleted, err = s.Del("A", "D")

	if deleted {
		t.Fatalf("somehow deleted non-existent field")
	}

	if err == nil {
		t.Fatalf("error was returned incorrectly")
	}
}
