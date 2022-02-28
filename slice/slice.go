package slice

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("I am the initializer")
}

// return true if a slice contain an element
func ContainInt(source []int, element int) bool {
	for _, val := range source {
		if val == element {
			return true
		}
	}
	return false
}

// return true if all elements are not nil
func NotNil(items ...interface{}) bool {
	for _, item := range items {
		val := reflect.ValueOf(item)
		if val.IsNil() {
			return false
		}
	}
	return true
}
