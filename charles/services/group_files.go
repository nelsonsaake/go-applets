package services

import (
	"fmt"
	"os"
	"path/filepath"
	"projects/applets/charles/helpers"
	"projects/applets/charles/nkspath"
)

func makeChunkDir(dir string, i int) string {
	return filepath.Join(dir, fmt.Sprint("_", i))
}

func preCheckChunkDirs(dir string, count int) {
	for i := 0; i < count; i++ {
		chunkDir := makeChunkDir(dir, i)
		if nkspath.IsExistingDir(chunkDir) {
			panic(fmt.Errorf(chunkDir, "already exists!"))
		}
	}
}

func pushChunksToDifferentDirs(dir string, chunks [][]string) {
	for i := 0; i < len(chunks); i++ {

		chunk := chunks[i]
		chunkDir := makeChunkDir(dir, i)

		for _, path := range chunk {

			base := filepath.Base(path)
			newPath := filepath.Join(chunkDir, base)

			err := os.MkdirAll(filepath.Dir(newPath), 0x6666)
			if err != nil {
				panic(err)
			}

			err = os.Rename(path, newPath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func GroupFiles(dir string, size int) [][]string {
	if !nkspath.IsExistingDir(dir) {
		return [][]string{}
	}

	files := nkspath.DirFiles(dir)

	chunks := helpers.Chunk(files, size)

	preCheckChunkDirs(dir, len(chunks))

	pushChunksToDifferentDirs(dir, chunks)

	return chunks
}
