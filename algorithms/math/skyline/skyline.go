// Package skyline provides Skyline Problem solution
// https://www.geeksforgeeks.org/divide-and-conquer-set-7-the-skyline-problem/
package skyline

type building struct {
	left, ht, right int
}

type strip struct {
	left, ht int
}

// BruteForce iterates over each building and scans through the rest to find out
// next overlapping strip
// Complexity: O(nˆ2)
func BruteForce(in []building) []strip {
	lenIn := len(in)
	out := []strip{}

	for i := 0; i < lenIn; i++ {
		b := in[i]
		s := queryStrip(b.left, b.ht, b.ht, in)
		if s != nil {
			out = append(out, *s)
		}
		s = queryStrip(b.right, 0, b.ht, in)
		if s != nil {
			out = append(out, *s)
		}
	}

	// in-place sort by strip x coordinate
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out)-i-1; j++ {
			// last i elements already sorted out
			if out[j].left > out[j+1].left {
				out[j], out[j+1] = out[j+1], out[j]
			}
		}
	}
	return out
}

func queryStrip(x, h, lh int, in []building) *strip {
	for j := 0; j < len(in); j++ {
		other := in[j]
		if x == other.left || x == other.right {
			// skip it self
			continue
		}
		if other.ht < h || other.right < x {
			// not overlap strip or other cannot raise strip height, continue looking
			continue
		}
		if other.left > x {
			// out of range, not inspect any further
			break
		}
		if other.ht > h {
			if other.left < x && x < other.right && lh < other.ht {
				// skip as completely inside other
				return nil
			}
			// raise strip height bar
			h = other.ht
		}
	}
	return &strip{x, h}
}
