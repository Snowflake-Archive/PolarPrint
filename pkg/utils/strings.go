package utils

import "strings"

func UnpackPrinterArray(data string) []string {
	return strings.Split(data, ";")
}

func PackPrinterArray(data []string) string {
	return strings.Join(data, ";")
}
