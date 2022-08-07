package main

import "fmt"

func main() {

	strs := []string{"flower", "fkow"}
	lesslet := strs[0]

	lenght := 0
	for i := 0; i < len(strs); i++ {
		if len(lesslet) >= len(strs[i]) {
			lesslet = strs[i]

			lenght = len(strs[i])
		}

	}

	var prefix string
	for x := 0; x <= len(strs) || x <= len(lesslet); x++ {
		if lenght < 0 {
			break
		}
		var count int
		for i := range strs {

			if lesslet[0:(lenght)] == strs[i][0:(lenght)] {
				fmt.Println("เช็ค", lesslet[0:(lenght)])
				count += 1
			}
			if count == len(strs) {
				prefix = (lesslet[0:(lenght)])
				fmt.Println("เช็ค prefix", prefix)
			}

		}
		if prefix != "" {
			break
		}
		lenght -= 1
	}
	if prefix != "" {
		fmt.Println(prefix)
	} else {
		fmt.Println("")
	}
}
