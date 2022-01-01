package main

import "fmt"

func main() {
	followers := []string{"hiro", "taro", "hiro", "yoshi", "yoshi", "tatsu", "kousuke", "kousuke", "tatsu"}

	uniques := make([]string, 0, len(followers))
	m := make(map[string]bool)
	for _, v := range followers {
		if m[v] {
			continue
		}
		uniques = append(uniques, v)
		m[v] = true
	}
	fmt.Println(uniques)
}
