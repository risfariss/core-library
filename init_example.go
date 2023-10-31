package core_library

import (
	"fmt"
	"io"
	"os"
)

/*
todo copy paste function here for testing
*/

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	_, _ = io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

	// Output:
	// Kim is 22 years old.}
}
