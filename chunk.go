package piscine

func chunkSlice(slice []int, chunkSize int) {

    if chunkSize == 0 {
        fmt.Println()
    } else {
        var chunks [][]int
        for i := 0; i < len(slice); i += chunkSize {
            end := i + chunkSize

            // necessary check to avoid slicing beyond
            // slice capacity
            if end > len(slice) {
                end = len(slice)
            }

            chunks = append(chunks, slice[i:end])
        }
        fmt.Println(chunks)
    }

}
