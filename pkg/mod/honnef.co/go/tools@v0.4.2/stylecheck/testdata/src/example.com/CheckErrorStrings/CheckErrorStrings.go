// Package pkg ...
package pkg

import "errors"

func fn() {
	errors.New("a perfectly fine error")
	errors.New("Not a great error")       //@ diag(`error strings should not be capitalized`)
	errors.New("also not a great error.") //@ diag(`error strings should not end with punctuation or newlines`)
	errors.New("URL is okay")
	errors.New("SomeFunc is okay")
	errors.New("URL is okay, but the period is not.") //@ diag(`error strings should not end with punctuation or newlines`)
	errors.New("T must not be nil")
	errors.New("Foo() failed")
	errors.New("Foo(bar) failed")
	errors.New("Foo(bar, baz) failed")
	errors.New("P384 is a nice curve")
}

func Write() {
	errors.New("Write: this is broken")
}

type T struct{}

func (T) Read() {
	errors.New("Read: this is broken")
	errors.New("Read failed")
}

func fn2() {
	// The error string hasn't to be in the same function
	errors.New("Read failed")
}
