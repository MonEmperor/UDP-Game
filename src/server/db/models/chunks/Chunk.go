package chunks

import "server/conf"

type Chunk struct {
	X int
	Y int
}

func ToChunkCoords(pos []int) []int { // converts block coordinates to chunk coordinates
	chunkCoords := make([]int, 2)
	for i := range pos {
		chunkCoords[i] = int(pos[i] / conf.CHUNK_SIZE)
		if chunkCoords[i] < 0 {
			chunkCoords[i] -= 1
		}
	}

	return chunkCoords
}

func ChunksInRenderDist(curCunk []int) [][]int { // returns all chunks within render distance
	chunks := [][]int{}
	i := 0
	for x := curCunk[0] - conf.RENDER_DISTANCE; x < curCunk[0]+conf.RENDER_DISTANCE; x++ {
		for y := curCunk[1] - conf.RENDER_DISTANCE; y < curCunk[1]+conf.RENDER_DISTANCE; y++ {
			chunks = append(chunks, []int{x, y})
			i++
		}
	}
	return chunks
}