package flowchart

import (
	"fmt"
	"testing"
)

func setup(tb testing.TB) func(tb testing.TB) {
	fmt.Println("Setup for unit tests.")
	Counter = 0

	return func(tb testing.TB) {
		fmt.Println("Tear down setup.")

	}
}
