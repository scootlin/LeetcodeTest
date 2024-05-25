// flipAndInvertImage flips and inverts a given image represented as a 2D array of integers.
// It takes a 2D array A as input and returns the modified array after flipping and inverting each row.
package leetcodetest

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		last := len(row) - 1
		for j := 0; j <= last/2; j++ {
			row[j], row[last-j] = row[last-j], row[j]
			row[j] = 1 - row[j]
			if j != last-j {
				row[last-j] = 1 - row[last-j]
			}
		}
	}
	return A
}
