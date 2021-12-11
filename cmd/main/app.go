package main

import (
	"bufio"
	"github.com/kodaf1/finite-state-machine/internal/fms"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	data, err := os.Open("test.txt")
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

		state, err := fms.Run(cmd)
		log.Printf(
			"%d test returns %t where last state is %v",
			i,
			err == nil && state.IsFinal(),
			reflect.TypeOf(state))
	}
}
