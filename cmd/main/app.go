package main

import (
	"bufio"
	"github.com/kodaf1/finite-state-machine/internal/fms"
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
	"github.com/kodaf1/finite-state-machine/internal/storage"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const ChunksCount = 10

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(data)

	for i := 1; ; i++ {
		cmd, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("EndOfFile")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		cmd = strings.TrimSpace(cmd)

		var wg sync.WaitGroup

		var resultStorage sync.Map

		maxChunks := getMaxChunks(len(cmd))
		chunkSize := len(cmd) / maxChunks

		wg.Add(maxChunks)
		for j := 0; j < maxChunks; j++ {
			from := j * chunkSize
			to := (j + 1) * chunkSize
			if j == maxChunks-1 {
				to = len(cmd)
			}
			go func(cmd string, chunkId int) {
				resultStorage.LoadOrStore(chunkId, fms.PartialRun(cmd, chunkId == 0))
				wg.Done()
			}(cmd[from:to], j)
		}

		wg.Wait()

		// sync
		state := states.GetStartState()
		for j := 0; j < maxChunks; j++ {
			chunkResult, _ := resultStorage.Load(j)

			newState, ok := chunkResult.(storage.Data)[state]
			if !ok {
				state = states.GetErrorState()
				break
			}
			state = newState
		}
		log.Printf("%d => %s : %t", i, states.GetStateName(state), state.IsFinal())
	}
}

func getMaxChunks(cmdLen int) int {
	maxChunks := 1
	if cmdLen >= ChunksCount {
		maxChunks = ChunksCount
	}
	if cmdLen >= ChunksCount*20 {
		maxChunks = ChunksCount * 2
	}
	return maxChunks
}
