package main

import (
	"fmt"
	"testing"
)

func TestFilterGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getFilterFunction(listName, typeName, "", ""))

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

	if result != expected {
		t.Fail()
	}
}

func TestPFilterGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getPFilterFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // PFilter is similar to the Filter method except that the filter is applied to all the elements in parallel. The order of resulting elements cannot be guaranteed. 
        func (l %[1]s) PFilter(f func(%[2]s) bool) %[1]s {
            wg := sync.WaitGroup{}
            mutex := sync.Mutex{}
            l2 := []%[2]s{}
            for _, t := range l {
                wg.Add(1)
                go func(t %[2]s){
                    if f(t) {
                        mutex.Lock()
                        l2 = append(l2, t)
                        mutex.Unlock()
                    }            
                    wg.Done()
                }(t)
            }
            wg.Wait()
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestEachGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getEachFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // Each is a method on %[1]s that takes a function of type %[2]s -> void and applies the function to each member of the list and then returns the original list.
        func (l %[1]s) Each(f func(%[2]s)) %[1]s {
            for _, t := range l {
                f(t) 
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestEachIGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getEachIFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // EachI is a method on %[1]s that takes a function of type (int, %[2]s) -> void and applies the function to each member of the list and then returns the original list. The int parameter to the function is the index of the element.
        func (l %[1]s) EachI(f func(int, %[2]s)) %[1]s {
            for i, t := range l {
                f(i, t) 
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestDropWhileGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getDropWhileFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // DropWhile is a method on %[1]s that takes a function of type %[2]s -> bool and returns a list of type %[1]s which excludes the first members from the original list for which the function returned true
        func (l %[1]s) DropWhile(f func(%[2]s) bool) %[1]s {
            for i, t := range l {
                if !f(t) {
                    return l[i:]
                }
            }
            var l2 %[1]s
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestTakeWhileGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getTakeWhileFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // TakeWhile is a method on %[1]s that takes a function of type %[2]s -> bool and returns a list of type %[1]s which includes only the first members from the original list for which the function returned true
        func (l %[1]s) TakeWhile(f func(%[2]s) bool) %[1]s {
            for i, t := range l {
                if !f(t) {
                    return l[:i]
                }
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestTakeGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getTakeFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // Take is a method on %[1]s that takes an integer n and returns the first n elements of the original list. If the list contains fewer than n elements then the entire list is returned.
        func (l %[1]s) Take(n int) %[1]s {
            if len(l) >= n {
                return l[:n]
            }
            return l
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestDropGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getDropFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // Drop is a method on %[1]s that takes an integer n and returns all but the first n elements of the original list. If the list contains fewer than n elements then an empty list is returned.
        func (l %[1]s) Drop(n int) %[1]s {
            if len(l) >= n {
                return l[n:]
            }
            var l2 %[1]s
            return l2
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestReduceGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getReduceFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // Reduce is a method on %[1]s that takes a function of type (%[2]s, %[2]s) -> %[2]s and returns a %[2]s which is the result of applying the function to all members of the original list starting from the first member
        func (l %[1]s) Reduce(t1 %[2]s, f func(%[2]s, %[2]s) %[2]s) %[2]s {
            for _, t := range l {
                t1 = f(t1, t)
            }
            return t1
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestReduceRightGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getReduceRightFunction(listName, typeName, "", ""))

	expectedRaw := fmt.Sprintf(`
        // ReduceRight is a method on %[1]s that takes a function of type (%[2]s, %[2]s) -> %[2]s and returns a %[2]s which is the result of applying the function to all members of the original list starting from the last member
        func (l %[1]s) ReduceRight(t1 %[2]s, f func(%[2]s, %[2]s) %[2]s) %[2]s {
            for i := len(l) - 1; i >= 0; i-- {
                t := l[i]
                t1 = f(t, t1)
            }
            return t1
        }
        `, listName, typeName)

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestMapGeneration1(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "string", ""
	result := f(getMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // Map is a method on stringList that takes a function of type string -> string and applies it to every member of stringList
        func (l stringList) Map(f func(string) string) stringList {
            l2 := make(stringList, len(l))
            for i, t := range l {
                l2[i] = f(t)
            }
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestMapGeneration2(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "int"
	result := f(getMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // MapInt is a method on stringList that takes a function of type string -> int and applies it to every member of stringList
        func (l stringList) MapInt(f func(string) int) intList {
            l2 := make(intList, len(l))
            for i, t := range l {
                l2[i] = f(t)
            }
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestMapGeneration3(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "I"
	result := f(getMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // MapI is a method on stringList that takes a function of type string -> int and applies it to every member of stringList
        func (l stringList) MapI(f func(string) int) intList {
            l2 := make(intList, len(l))
            for i, t := range l {
                l2[i] = f(t)
            }
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestPMapGeneration1(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "string", ""
	result := f(getPMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // PMap is similar to Map except that it executes the function on each member in parallel.
        func (l stringList) PMap(f func(string) string) stringList {
            wg := sync.WaitGroup{}
            l2 := make(stringList, len(l))
            for i, t := range l {
                wg.Add(1)
                go func(i int, t string) {
			l2[i] = f(t)
			wg.Done()
		}(i, t)
            }
            wg.Wait()
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestPMapGeneration2(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "int"
	result := f(getPMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // PMapInt is similar to MapInt except that it executes the function on each member in parallel.
        func (l stringList) PMapInt(f func(string) int) intList {
            wg := sync.WaitGroup{}
            l2 := make(intList, len(l))
            for i, t := range l {
                wg.Add(1)
                go func(i int, t string) {
			l2[i] = f(t)
			wg.Done()
		}(i, t)
            }
            wg.Wait()
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestPMapGeneration3(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "I"
	result := f(getPMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // PMapI is similar to MapI except that it executes the function on each member in parallel.
        func (l stringList) PMapI(f func(string) int) intList {
            wg := sync.WaitGroup{}
            l2 := make(intList, len(l))
            for i, t := range l {
                wg.Add(1)
                go func(i int, t string) {
			l2[i] = f(t)
			wg.Done()
		}(i, t)
            }
            wg.Wait()
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestAllGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getAllFunction(listName, typeName, "", ""))

	expectedRaw := `
        // All is a method on stringList that returns true if all the members of the list satisfy a function or if the list is empty.
        func (l stringList) All(f func(string) bool) bool {
            for _, t := range l {
                if !f(t) {
                    return false
                }
            }
            return true
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestAnyGeneration(t *testing.T) {
	listName, typeName := "stringList", "string"
	result := f(getAnyFunction(listName, typeName, "", ""))

	expectedRaw := `
        // Any is a method on stringList that returns true if at least one member of the list satisfies a function. It returns false if the list is empty.
        func (l stringList) Any(f func(string) bool) bool {
            for _, t := range l {
                if f(t) {
                    return true
                }
            }
            return false
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestFilterMapGeneration(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "int"
	result := f(getFilterMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // FilterMapInt is a method on stringList that applies the filter(s) and map to the list members in a single loop and returns the resulting list.
        func (l stringList) FilterMapInt(fMap func(string) int, fFilters ...func(string) bool) intList {
            l2 := intList{}
            for _, t := range l {
                pass := true
                for _, f := range fFilters {
                    if !f(t){
                        pass = false
                        break
                    }
                }
                if pass {
                    l2 = append(l2, fMap(t))
                }
            }
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}

func TestPFilterMapGeneration(t *testing.T) {
	listName, typeName, targetType, targetTypeName := "stringList", "string", "int", "int"
	result := f(getPFilterMapFunction(listName, typeName, targetType, targetTypeName))

	expectedRaw := `
        // PFilterMapInt is similar to FilterMapInt except that it executes the method on each member in parallel.
        func (l stringList) PFilterMapInt(fMap func(string) int, fFilters ...func(string) bool) intList {
            l2 := intList{}
            mutex := sync.Mutex{}
            wg := sync.WaitGroup{}
            wg.Add(len(l))
            
            for _, t := range l {
                go func(t string){
                    pass := true
                    for _, f := range fFilters {
                        if !f(t) {
                            pass = false
                            break
                        }
                    }
                    if pass {
                        mutex.Lock()
                        l2 = append(l2, fMap(t))
                        mutex.Unlock()
                    }
                    wg.Done()
                }(t)
            }
            wg.Wait()
            return l2
        }
        `

	expected := f(expectedRaw)

	if result != expected {
		t.Fail()
	}
}
