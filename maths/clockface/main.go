package main

import (
	"os"
	"time"

	clockface "github.com/pmuens/learn-go-with-tests/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
