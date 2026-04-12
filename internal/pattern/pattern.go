package pattern

// FromDigest generates a 5x5 grid based on the provided digest data.
func FromDigest(data [16]byte) [5][5]bool {
	var grid [5][5]bool

	nibbleIndex := 0
	for col := 2; col >= 0; col-- {
		mirrorCol := 4 - col

		for row := 0; row < 5; row++ {
			b := data[nibbleIndex/2]

			shift := 4
			if nibbleIndex%2 == 1 {
				shift = 0
			}

			paint := ((b >> shift) & 1) == 0
			grid[row][col] = paint
			grid[row][mirrorCol] = paint

			nibbleIndex++
		}
	}

	return grid
}
