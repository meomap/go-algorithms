// Package karatsuba provides Binary String Multiplication impelementation
package karatsuba

import (
	"fmt"
	"strconv"
)

func BruteForce(a, b []byte) int {
	binaryResult := []byte{}
	lenA, lenB := len(a), len(b)
	// start from second row
	for i := lenB - 1; i >= 0; i-- {

		// prepare current row sum placeholders
		placeHolders := lenB - (i + 1)
		lineSum := make([]byte, placeHolders)
		for k := 0; k < placeHolders; k++ {
			lineSum[k] = '0'
		}

		for j := lenA - 1; j >= 0; j-- {
			product := multipleBinary(b[i], a[j])
			lineSum = append([]byte{product}, lineSum...)
		}

		// aggregate line
		binaryResult = addBinaryChars(binaryResult, lineSum)
	}
	result, err := strconv.ParseInt(string(binaryResult), 2, 64)
	if err != nil {
		panic(fmt.Sprintf("Expect binary string result but got %x with error %q", binaryResult, err))
	}

	return int(result)
}

// http://www.binarymath.info/binary-multiplication.php
func multipleBinary(b1, b2 byte) byte {
	if b1 == '1' && b2 == '1' {
		return '1'
	}
	return '0'
}

func addBinares(b1, b2, b3 byte) []byte {
	out := addBinary(b1, b2)
	var carrier byte
	if len(out) == 2 {
		carrier, out[0] = out[0], out[1]
	}
	out = addBinary(out[0], b3)

	if carrier != 0 {
		return []byte{carrier, out[0]}
	}
	return out

}

// http://www.binarymath.info/binary-addition.php
func addBinary(b1, b2 byte) []byte {
	if b2 == 0 {
		// exception case when consume all list with extra carrier value
		return []byte{b1}
	}

	if b1 == '0' && b2 == '0' {
		return []byte{'0'}
	} else if b1 == '1' && b2 == '1' {
		return []byte{'1', '0'}
	}
	// either b1 or b2 equals 1 or 0
	return []byte{'1'}
}

func addBinaryChars(a, b []byte) (out []byte) {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return b
	}
	if lenB == 0 {
		return a
	}
	i, j := lenA, lenB
	var carrier byte
	for {
		i--
		j--
		// add column binary with carrier from last column
		sum := addBinares(a[i], b[j], carrier)
		carrier = 0
		if len(sum) == 2 {
			// move carrier to next column
			carrier, sum[0] = sum[0], sum[1]
		}
		// prepand column sum to final result
		out = append([]byte{sum[0]}, out...)
		if i == 0 || j == 0 {
			break
		}
	}
	if i == 0 && j > 0 {
		// consume all remaining from b
		for ; j > 0; j-- {
			j--
			sum := addBinary(b[j], carrier)
			carrier = 0
			out = append([]byte{sum[0]}, out...)
		}
	} else if j == 0 && i > 0 {
		for ; i > 0; i-- {
			sum := addBinary(b[j], carrier)
			carrier = 0
			out = append([]byte{sum[0]}, out...)
		}
	} else {
		// same length
		if carrier != 0 {
			out = append([]byte{carrier}, out...)
		}
	}
	return out
}
