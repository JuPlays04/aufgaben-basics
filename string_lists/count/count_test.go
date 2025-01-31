package count

import "fmt"

func ExampleCount() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}

	fmt.Println(Count(s1, "a"))
	fmt.Println(Count(s1, "b"))
	fmt.Println(Count(s1, "c"))

	// Output:
	// 3
	// 2
	// 1
}

func ExampleContainsDuplicate() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}
	s2 := []string{"a", "b", "c", "d"}

	fmt.Println(ContainsDuplicate(s1))
	fmt.Println(ContainsDuplicate(s2))

	// Output:
	// true
	// false
}
