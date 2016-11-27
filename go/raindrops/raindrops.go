package raindrops

import "fmt"

const testVersion = 2

func Convert(a int) string {
	song := ""
	if (a % 3) == 0 {
		song = "Pling"
	}
	if (a % 5) == 0 {
		song = fmt.Sprintf("%vPlang", song)
	}
	if (a % 7) == 0 {
		song = fmt.Sprintf("%vPlong", song)
	}
	if len(song) == 0 {
		song = fmt.Sprintf("%v", a)
	}
	return song
}
