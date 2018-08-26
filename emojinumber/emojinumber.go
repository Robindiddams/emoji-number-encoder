package emojinumber

import (
	"fmt"
	"strings"
)

// Encode will take a 36 bit number and
// return a 4 emoji string
func Encode(src int) string {
	return strings.Join([]string{
		string(encode[src>>0&0x1FF]),
		string(encode[src>>9&0x1FF]),
		string(encode[src>>18&0x1FF]),
		string(encode[src>>27&0x1FF]),
	}, "")
}

// ParseEmojis will parse out all emojis in a []rune,
// useful for large blocks of emojis since emojis can
// span many runes, requires an emoji set to reference
func ParseEmojis(r []rune) []int {
	// looks for
	reduceForEmoji := func(index int) (int, int) {
		for subset := len(r); subset >= 0; subset-- {
			found, emojiIndex := searchEmoji(r[index:subset])
			if found {
				return subset, emojiIndex
			}
		}
		return 0, 0
	}
	emojis := make([]int, 0)
	index := 0
	for {
		lastEmoji, emojiNum := reduceForEmoji(index)
		if lastEmoji == 0 {
			fmt.Println("this is weird")
			// TODO: investigate, skip for now
			index++
			continue
		}
		emojis = append(emojis, emojiNum)
		index = lastEmoji
		if index > 4 {
			break
		}
	}
	return emojis
}

func searchEmoji(maybe []rune) (bool, int) {
	for i, emoji := range encode {
		if string(maybe) == string(emoji) {
			return true, i
		}
	}
	return false, 0
}

// Decode will turn a string of 4 emojis into a 36 bit number
func Decode(src string) int {
	emojis := ParseEmojis([]rune(src))
	return emojis[3]<<27 | emojis[2]<<18 | emojis[1]<<9 | emojis[0]
}
