package util

import "fmt"

func ConvertInterfaceToString(i interface{}) string {
	if i != nil {
		return fmt.Sprint(i)
	}
	return ""
}
