package getenv

import (
	"os"
	"strings"
)

func envExist(env string) (result bool) {
	value := os.Getenv(env)
	if !strings.EqualFold(value, "") {
		result = true
	}
	return
}
