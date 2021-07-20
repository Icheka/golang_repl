package repl

import (
	"bufio"
	"fmt"
	"io"
)

type repl struct {
	prompt    string
	exitWords []string
}

func NewRepl(prompt string, exitWords []string) *repl {
	return &repl{prompt, exitWords}
}

func (r *repl) Start(i io.Reader, o io.Writer) {
	// create a new Scanner using the bufio package
	scanner := bufio.NewScanner(i)

	for {
		// print the prompt to StdErr
		fmt.Fprintf(o, r.prompt)

		// scanner.Scan() reads the next token in the input
		// returns false if it has reached the end of the input
		// (or if there was an error, but this is a less-likely situation)
		// if this was unsuccessful - because there is nothing else to scan, for example,
		// stop attempting any more scans (return)
		isScannedNext := scanner.Scan()
		if !isScannedNext {
			return
		}

		// the scanned input is available via the Text() method
		// it returns a 'line' of input
		line := scanner.Text()

		for _, exitWord := range r.exitWords {
			if line == exitWord {
				io.WriteString(o, "Quiting REPL...\n")
				return
			}
		}

		// returnMessage := fmt.Sprintf("Go_REPL says back:%s", line)
		io.WriteString(o, fmt.Sprintf("Go_REPL says back:  %s \n\n", line))
	}
}
