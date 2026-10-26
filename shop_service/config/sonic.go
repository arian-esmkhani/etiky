package config

import (
	"github.com/bytedance/sonic"
)

// SetupSonic initializes sonic with custom configuration if needed.
// Currently uses defaults; reserved for future tuning (e.g., JSON validation, number handling).
func SetupSonic() {
}

func MarshalToString(v interface{}) (string, error) {
	return sonic.MarshalString(v)
}

func UnmarshalFromString(data string, v interface{}) error {
	return sonic.UnmarshalString(data, v)
}

const EnablePrefork = false