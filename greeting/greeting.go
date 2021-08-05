package greeting

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	if name == "" {
		name = "my friend"
	}
	result := fmt.Sprintf("Hello ,%s.", name)
	if strings.ToUpper(name) == name {
		result = strings.ToUpper(result)
	}

	return result
}
