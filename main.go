package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	const layout = "Now, Monday Jan 02 15:04:05 JST 2006"
	fmt.Println(t.Format(layout))

}
