package main

import (
	"bufio"
	"github.com/kodaf1/finite-state-machine/internal/fms"
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
	"github.com/kodaf1/finite-state-machine/internal/storage"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

const PROCS = 4

func main() {
	runtime.GOMAXPROCS(PROCS) // control amount of procs. default - # of logical processor you have.

	data, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(data)

	for i := 0; ; i++ {
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
		maxChunks := 1
		if len(cmd) >= PROCS {
			maxChunks = PROCS
		}

		for j := 0; j < maxChunks; j++ {
			wg.Add(1)
			go func(cmd string, chunkId int) {
				resultStorage.Store(chunkId, fms.PartialRun(cmd, chunkId == 0))
				wg.Done()
			}(cmd[j*len(cmd)/maxChunks:(j+1)*len(cmd)/maxChunks], j)
		}

		wg.Wait()

		// sync
		state := states.GetStartState()
		for j := 0; j < maxChunks; j++ {
			chunkResult, _ := resultStorage.Load(j)
			newState, ok := chunkResult.(storage.Data)[state]
			if !ok {
				state = states.Error{}
				break
			}
			state = newState
		}
		log.Printf("%s => %t", cmd, state.IsFinal())
	}
}
