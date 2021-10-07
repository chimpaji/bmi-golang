package main

import (
	"bufio"
	"os"
)

//cant use const because the functino before program start dont have any value now, unlike js
var reader = bufio.NewReader(os.Stdin)