package pattern

type BitReader func(int) bool

// FromDigest generates a 5x5 grid based on the provided digest data.
func FromDigest(data [16]byte) [5][5]bool {
	var grid [5][5]bool

	bitIndex := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < (5+1)/2; col++ {
			byteIndex := bitIndex / 8
			shift := uint(bitIndex % 8)
			on := (data[byteIndex]>>shift)&1 == 1

			grid[row][col] = on
			grid[row][5-1-col] = on
			bitIndex++
		}
	}

	return grid
}
