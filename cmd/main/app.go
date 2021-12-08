package main

import (
	"bufio"
	"fmt"
	"github.com/kodaf1/finite-state-machine/internal/fms"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmd, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	cmd = strings.TrimSpace(cmd)

	fmt.Println(fms.Run(cmd))
}
