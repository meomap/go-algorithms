// https://www.geeksforgeeks.org/longest-common-prefix-set-3-divide-and-conquer/
package longest

import "bytes"

// BruteForce iteratively scans through whole collection to find out longest common prefix
func BruteForce(in []string) (string, error) {
	out := bytes.NewBuffer(nil)
	lenIn := len(in)

	// examine chars from first word
	lenWord := len(in[0])
	var err error
	charIdx := 0

LOOP:
	for charIdx < lenWord {
		nextChar := in[0][charIdx]
		for i := 1; i < lenIn; i++ {
			nextWord := in[i]
			if len(nextWord) <= charIdx || nextWord[charIdx] != nextChar {
				// no more common prefix
				break LOOP
			}
		}
		if err = out.WriteByte(nextChar); err != nil {
			return "", err
		}
		charIdx++
	}
	return out.String(), nil
}

func mergeCommonPrefixes(left, right string) (string, error) {
	lenL, lenR := len(left), len(right)
	i := 0
	out := bytes.NewBuffer(nil)
	var err error
	for i < lenL && i < lenR && left[i] == right[i] {
		if err = out.WriteByte(left[i]); err != nil {
			return "", err
		}
		i++
	}
	return out.String(), nil
}

func longestCommonPrefix(in []string) (string, error) {
	lenIn := len(in)
	if lenIn == 1 {
		return in[0], nil
	}
	mid := lenIn / 2
	lc, err := longestCommonPrefix(in[:mid])
	if err != nil {
		return "", err
	}
	rc, err := longestCommonPrefix(in[mid:])
	if err != nil {
		return "", err
	}
	return mergeCommonPrefixes(lc, rc)
}

// DivideAndConquor implements strategy to merge two halves common longest prefix
func DivideAndConquor(in []string) (string, error) {
	return longestCommonPrefix(in)
}
