/*
Notes:

1. When the process ends, the variable create by os.Setenv no longer exists.

2. If you want to check the existence of am environment variable, you should use os.LookupEnv.
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

func TestEnvExist2(t *testing.T) {
	tests := []struct {
		Key    string
		Set    bool
		Answer bool
	}{
		{"IDK2", false, false},
		{"IDK2", true, true},
		{"IDK2", false, true},
	}

	for _, test := range tests {
		if test.Set {
			err := os.Setenv(test.Key, test.Key)
			if err != nil {
				t.Errorf("Failed to set environment variable %s: %s", test.Key, err.Error())
			}
		}

		result := envExist2(test.Key)
		if result != test.Answer {
			t.Errorf("For key %s (Set=%t), expected %t, but got %t", test.Key, test.Set, test.Answer, result)
		}
	}
}
