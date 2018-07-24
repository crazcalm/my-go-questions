/*
reference links:

1. https://stackoverflow.com/questions/31595791/how-to-test-panics
2. https://gobyexample.com/panic
*/

package panic

import (
	"testing"
)

func TestWillPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	//Call func
	willPanic()
}
