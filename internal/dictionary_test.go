package internal

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, d Dictionary, word string, definition string) {
	t.Helper()
	res, _ := d.Search(word)

	if res != definition {
		t.Errorf("got word '%s' want '%s'", res, definition)
	}
}

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("first test", func(t *testing.T) {
		dictionary := Dictionary{word: definition}

		got, err := dictionary.Search(word)
		want := "this is just a test"

		assertStrings(t, got, want)
		assertError(t, err, nil)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "someword"
		dictionary := Dictionary{}

		_, err := dictionary.Search(word)
		want := ErrNotFound

		assertStrings(t, err.Error(), want.Error())
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("adding a word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if want != got {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("adding a word that already exists", func(t *testing.T) {
		dictionary := Dictionary{}
		err1 := dictionary.Add(word, definition)
		assertError(t, err1, nil)

		err2 := dictionary.Add(word, definition)

		want := ErrAlreadyExists
		assertError(t, err2, want)

	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}
