package easy

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	// 使用 map 存储字符及其索引
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0 // 滑动窗口的起始位置

	for end := 0; end < len(s); end++ {
		if index, exists := charIndex[s[end]]; exists && index >= start {
			// 如果字符重复，更新滑动窗口的起始位置
			start = index + 1
		}

		// 更新字符的位置
		charIndex[s[end]] = end
		// 计算窗口长度并更新最大值
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func LengthOfLongest(s string) int {
	CharMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if charIndex, exists := CharMap[s[end]]; exists && charIndex >= start {
			start = charIndex + 1
		}
		CharMap[s[end]] = end

		currentLength := end - start + 1

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	expected := 3
	result := lengthOfLongestSubstring(s)
	fmt.Println("The length of the longest substring without repeating characters is:", lengthOfLongestSubstring(s))
	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}
