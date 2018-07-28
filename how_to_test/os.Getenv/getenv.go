package getenv

import (
	"os"
	"strings"
)

func envValue(env, defaultValue string) (result string, err error) {
	result = os.Getenv(env)
	if strings.EqualFold(result, "") {
		err = os.Setenv(env, defaultValue)
		if err != nil {
			return
		}
		result = os.Getenv(env)
	}
	return
}

func envExist(env string) (result bool) {
	value := os.Getenv(env)
	if !strings.EqualFold(value, "") {
		result = true
	}
	return
}

func envExist2(env string) (result bool) {
	_, result = os.LookupEnv(env)
	return
}
