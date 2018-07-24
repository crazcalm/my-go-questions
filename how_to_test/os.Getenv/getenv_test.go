/*
When the process ends, the variable create by os.Setenv no longer exists.
*/

package getenv

import (
	"os"
	"testing"
)

func TestEnvExist(t *testing.T) {
	tests := []struct {
		Key    string
		Set    bool
		Answer bool
	}{
		{"IDK", false, false},
		{"IDK", true, true},
		{"IDK", false, true},
	}

	for _, test := range tests {
		if test.Set {
			err := os.Setenv(test.Key, test.Key)
			if err != nil {
				t.Errorf("Failed to set environment variable %s: %s", test.Key, err.Error())
			}
		}

		result := envExist(test.Key)
		if result != test.Answer {
			t.Errorf("For key %s (Set=%t), expected %t, but got %t", test.Key, test.Set, test.Answer, result)
		}
	}
}
