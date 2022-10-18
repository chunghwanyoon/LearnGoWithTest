package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertNoError(t, err)
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		assertError(t, err, ErrorNotFound)
		assertString(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add to dictionary", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := dict.Search("test")
		assertNoError(t, err)
		assertString(t, got, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should find added word:", err)
	}
}
