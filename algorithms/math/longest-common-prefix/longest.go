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

// DivideAndConquor implement strategy to merge to common longest prefix
func DivideAndConquor(in []string) (string, error) {
	return "", nil
}
