package main

import (
	"bufio"
	"os"
)

func buwr() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("heloo world\n")
	_, _ = writer.WriteString("selamat malam\n")
	writer.Flush()
}
