package main

import (
	"encoding/csv"
	"os"
)

func writ() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"daoa", "dapa", "dafa"})
	_ = writer.Write([]string{"dono", "budi", "dono"})

	writer.Flush()
}
