package main

import (
	"fmt"
	"testing"
)

func TestFilterGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	filter := f(getFilterFunction(listName, typeName))

	expectedRaw := fmt.Sprintf(`
        // Filter is a method on %[1]s that takes a function of type %[2]s -> bool returns a list of type %[1]s which contains all members from the original list for which the function returned true
        func (l %[1]s) Filter(f func(%[2]s) bool) %[1]s {
            l2 := []%[2]s{}
            for _, t := range l {
                if f(t) {
                    l2 = append(l2, t)
                }
            }
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if filter != expected {
		t.Fail()
	}
}
