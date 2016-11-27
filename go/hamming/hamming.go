package hamming

import (
	"errors"
)

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("The genomes have a diferent lenght")
	}
	count := 0
	for pos, char := range a {
		if string(char) != string(b[pos]) {
			count++
		}
	}
	return count, nil
}
