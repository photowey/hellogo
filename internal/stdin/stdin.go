package stdin

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func Run() {
	flag.Parse()
	data := flag.Args()
	if !terminal.IsTerminal(0) {
		b, err := io.ReadAll(os.Stdin)
		if err == nil {
			data = append(data, string(b))
		}
	}

	fmt.Println(strings.Join(data, " "))
}
