package lib

import (
	"os"
	"encoding/json"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func ToJson(resp string) map[string]interface{}{
	data := map[string]interface{}{}
	json.Unmarshal([]byte(resp), &data)
	return data
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}