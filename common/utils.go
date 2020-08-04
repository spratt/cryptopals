// Package common contains various utility functions that don't belong
// elsewhere.
package common

// PanicIfNotNil takes an error and will panic if the given error is
// not nil.
func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
