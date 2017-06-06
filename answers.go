// Exercise: Loops and Functions

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 20; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

// Exercise: Slices

func Pic(dx, dy int) [][]uint8 {
	// creates a slice of of length dy with slices that contain uint8 types
	ret := make([][]uint8, dy)

	// at each slice within ret, make a slice with length dx
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			// Within each element inside this matrix, create a value
			ret[i][j] = uint8(i^j + 2*i*j)
		}
	}
	return ret
}

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	splits := strings.Fields(s)
	
	for i := 0; i < len(splits); i++ {
		fmt.Printf("%q", splits[i]);
		counts[splits[i]] = strings.Count(s, splits[i])
	}
	return counts
}