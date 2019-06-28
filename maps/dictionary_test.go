package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("can get definition of known word", func(t *testing.T) {
		def, _ := d.Search("test")
		assertEquals(t, "this is just a test", def)
	})

	t.Run("expect error when word undefined", func(t *testing.T) {
		word := "foo"
		def, err := d.Search(word)

		assertEquals(t, "", def)
		assertError(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{}

	word := "test"
	def  := "this is just a test"

	t.Run("new word", func(t *testing.T) {
		err := d.Add(word, def)

		assertErrorIsNil(t, err)
		assertDefinition(t, d, word, def)
	})

	t.Run("existing word", func(t * testing.T) {
		_ = d.Add(word, def)

		err := d.Add(word, "something else")

		assertError(t, ErrWordExists, err)
		assertDefinition(t, d, word, def)
	})
}

func TestUpdate(t *testing.T) {
	d := Dictionary{}

	word   := "test"
	def    := "this is just a test"
	newDef := "this is just a test, only a test"

	t.Run("update new word should fail", func(t *testing.T) {
		updateError := d.Update(word, def)

		_, searchError := d.Search(word)

		assertError(t, ErrWordDoesNotExist, updateError)
		assertError(t, ErrNotFound, searchError)
	})

	t.Run("Update existing word", func(t *testing.T) {
		_ = d.Add(word, def)
		err := d.Update(word, newDef)

		assertErrorIsNil(t, err)
		assertDefinition(t, d, word, newDef)
	})
}

func TestHas(t *testing.T) {
	goodWord := "test"
	badWord  := "abcdefg"

	d := Dictionary{}

	_ = d.Add(goodWord, "this is a test")

	assertWordExists(t, d, goodWord)
	assertWordNotExists(t, d, badWord)
}

func TestDelete(t *testing.T) {
	//t.Error("https://github.com/quii/learn-go-with-tests/blob/master/maps.md#write-the-test-first-6");

	d := Dictionary{}

	word := "test"
	def  := "this is a test"

	t.Run("can not delete words that do not exist", func(t *testing.T) {
		d.Delete(word)

		assertWordNotExists(t, d, word)
	})

	t.Run("delete existing word", func (t *testing.T) {
		_ = d.Add(word, def)

		d.Delete(word)

		assertWordNotExists(t, d, word)
	})
}

func assertWordExists(t *testing.T, d Dictionary, word string) {
	t.Helper()

	if !d.Has(word) {
		t.Errorf("d.Has(%s) should be true but was false", word)
	}
}

func assertWordNotExists(t *testing.T, d Dictionary, word string) {
	t.Helper()

	if d.Has(word) {
		t.Errorf("d.Has(%s) should be false but was true", word)
	}
}

func assertEquals(t *testing.T, expected string, actual string) {
	t.Helper()

	if actual != expected {
		message := "Strings are not the same\n" +
			"expected\t'%s'\n" +
			"received\t'%s'"

		t.Errorf(message,	expected,	actual)
	}
}

func assertError(t *testing.T, expected, received error) {
	t.Helper()
	assertEquals(t, expected.Error(), received.Error())
}

func assertErrorIsNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}

func assertDefinition(t *testing.T, d Dictionary, word, expected string) {
	t.Helper()
	actual, err := d.Search(word)

	assertErrorIsNil(t, err)
	assertEquals(t, expected, actual)
}