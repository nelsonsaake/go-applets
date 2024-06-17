package helpers

import "math"

func Chunk(s []string, size int) [][]string {

	chunks := [][]string{}

	for i := 0; i < len(s); i += size {
		in := int(math.Min(float64(i+size), float64(len(s))))
		chunks = append(chunks, s[i:in])
	}

	return chunks
}
