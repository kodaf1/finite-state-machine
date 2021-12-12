package main

import (
	"bufio"
	"github.com/kodaf1/finite-state-machine/internal/fms"
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
	"io"
	"log"
	"os"
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

		result, err := fms.PartialRun(cmd)
		presentationMap := make(map[string]string)
		for k, v := range result {
			presentationMap[states.GetStateName(k)] = states.GetStateName(v)
		}
		log.Printf("%s => %v", cmd, presentationMap)
		//log.Printf(
		//	"%d test returns %t where last state is %v",
		//	i,
		//	err == nil && state.IsFinal(),
		//	reflect.TypeOf(state))
	}
}
