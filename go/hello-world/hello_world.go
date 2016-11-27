package hello

import "fmt"

const testVersion = 2

func HelloWorld(name string) string {
	if len(name) <= 0 {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
