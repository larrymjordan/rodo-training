package main

import (
	"challenge2/ipScanner"
	"fmt"
)

const iP = "186.159.114.6"

func main() {
	fmt.Printf("%v\n", ipScanner.ScanIP(iP))
}
