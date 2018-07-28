/*
Notes:

1. When the process ends, the variable create by os.Setenv no longer exists.

2. If you want to check the existence of am environment variable, you should use os.LookupEnv.
*/

package getenv

import (
	"os"
	"strings"
	"testing"
)

func TestEnvValueExplicit(t *testing.T) {
	tests := []struct {
		Key     string
		Value   string
		Default string
		Answer  string
	}{
		{"NAME", "", "default1", "default1"},
		{"NAME", "default1", "default2", "default1"},
	}

	for i, test := range tests {

		//Need to keep original value so that I can clean up after the test
		oldValue := os.Getenv(test.Key)

		//Ensuring that the environmental variable exists
		err := os.Setenv(test.Key, test.Value)
		if err != nil {
			t.Fatalf("Failed to set environment variable: %s", err.Error())
		}

		result, err := envValue(test.Key, test.Default)
		if err != nil {
			t.Errorf("Case (%d): envValue returned an err: %s", i, err.Error())
		}

		if !strings.EqualFold(result, test.Answer) {
			t.Errorf("Case (%d): Expected %s to be equal to %s, but got %s", i, test.Key, test.Answer, result)
		}

		err = os.Setenv(test.Key, oldValue)
		if err != nil {
			t.Logf("Case (%d): Failed to set environment variable back to %s: %s", i, oldValue, err.Error())
		}
	}
}

func TestEnvValueImplicit(t *testing.T) {
	tests := []struct {
		Key     string
		Default string
		Answer  string
	}{
		{"NAME", "default1", "default1"},
		{"NAME", "default2", "default1"},
	}

	for i, test := range tests {
		result, err := envValue(test.Key, test.Default)
		if err != nil {
			t.Errorf("envValue returned an err: %s", err.Error())
		}

		if !strings.EqualFold(result, test.Answer) {
			t.Errorf("Case (%d): Expected %s to be equal to %s, but got %s", i, test.Key, test.Answer, result)
		}
	}
}

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
