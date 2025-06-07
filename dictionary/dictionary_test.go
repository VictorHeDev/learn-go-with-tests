package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test": "this is a test",
	}
	got := Search(dictionary, "test")
	// got := dictionary["test"]
	want := "this is a test"

	if got != want {
		t.Errorf("got %q want %q, %q", got, want, "test should return the value for the key 'test'")
	}
}
