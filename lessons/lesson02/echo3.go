package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")
var newLineMark = flag.Bool("n", false, "new line mark")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
    if *newLineMark {
        fmt.Println()     
    }
}
