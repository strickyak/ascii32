// +build main

// echo ' 3 8 + ! ' | go run ascii32.go
package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/chzyer/readline"

	ascii32 "github.com/strickyak/ascii32"
)

var FlagI = flag.Bool("i", false, "run interactive shell even if command line scripts are given")

func main() {
	flag.Parse()

	__ := ascii32.New__(emit)

	for _, filename := range flag.Args() {
		bb, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Cannot read file %q: %v", filename, err)
		}
		__.RunProgram(string(bb))
	}

	if *FlagI || flag.NArg() == 0 {

		home := os.Getenv("HOME")
		if home == "" {
			home = "."
		}
		rl, err := readline.NewEx(&readline.Config{
			Prompt:          " ok ",
			HistoryFile:     filepath.Join(home, ".ascii32.history"),
			InterruptPrompt: "*SIGINT*",
			EOFPrompt:       "*EOF*",
			// AutoComplete:    completer,
			// HistorySearchFold:   true,
			// FuncFilterInputRune: filterInput,
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		for {
			os.Stderr.Write([]byte{'\n'})
			line, err := rl.Readline()
			if err == readline.ErrInterrupt {
				if len(line) == 0 {
					break
				} else {
					continue
				}
			} else if err == io.EOF {
				break
			}
			TryRunProgram(__, line)
		}
	}
}

func TryRunProgram(__ *ascii32.A__, line string) {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("*** ERROR *** %v", r)
		}
	}()
	__.RunProgram(line)
}

func emit(r rune) {
	buf := []byte(string([]rune{r}))
	_, err := os.Stdout.Write(buf)
	if err != nil {
		log.Fatalf("cannot emit: %v", err)
	}
}
