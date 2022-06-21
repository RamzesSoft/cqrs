// Package cqrs provides a CQRS reference implementation.
//
// The implementation follows as much as possible the classic reference implementation
// m-r by Greg Young.
//
// The implementation differs in a number of respects because the original is written
// in C# and uses Generics where generics are not available in Go.
// This implementation instead uses interfaces to deal with types in a generic manner
// and used delegate functions to instantiate specific types.
package cqrs

import (
	"reflect"
)

// TypeWithPackage is a convenience function that returns the name of a type
//
// This is used so commonly throughout the code that it is better to
// have this convenience function and also allows for changing the scheme
// used for the type name more easily if desired.
func TypeWithPackage(i interface{}) string {
	return reflect.TypeOf(i).Elem().String()
}

// Int returns a pointer to int.
//
// There are a number of places where a pointer to int
// is required such as expectedVersion argument on the repository
// and this helper function makes keeps the code cleaner in these
// cases.
//func Int(i int) *int {
//	return &i
//}
