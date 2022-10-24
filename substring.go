package main

import "fmt"

//A function which returns the length of longest substring.
//Which takes a String as input.
func SubString(inputString string) int {
	if len(inputString) == 0 {
		return 0
	}
	charLastIndex := map[rune]int{}
	currentSubStringLength, longestSubStringLength, start := 0, 0, 0
	for index, character := range inputString {
		if lastIndex, hasCharacter := charLastIndex[character]; !hasCharacter || lastIndex < index-currentSubStringLength {
			currentSubStringLength++
		} else {
			if currentSubStringLength > longestSubStringLength {
				longestSubStringLength = currentSubStringLength
			}
			start = lastIndex + 1
			currentSubStringLength = index - start + 1
		}
		charLastIndex[character] = index
	}
	if currentSubStringLength > longestSubStringLength {
		longestSubStringLength = currentSubStringLength
	}
	return longestSubStringLength
}

func main() {
	fmt.Println(SubString("abcabcbb"))
	fmt.Println(SubString("bbbbb"))
	fmt.Println(SubString("pwwkew"))
	fmt.Println(SubString("javaconceptoftheday"))
}
