package main

import (
	"fmt"
	"strings"
)

type Result struct {
	Level   string
	Message string
}

func main() {
	logs := []string{
		"INFO|2024-01-01 10:00:00|Server started",
		"WARN|2024-01-01 10:01:00|Low memory",
		"ERROR|2024-01-01 10:02:00|Connection failed",
		"INFO|2024-01-01 10:03:00|User login",
		"BROKEN LINE WITHOUT PIPES",
		"ERROR|2024-01-01 10:04:00|Timeout",
		"ERROR|ONE LINE STRING",
	}

	var errors int
	results := make(chan Result)

	go func() {
		for _, i := range logs {
			if strings.Count(i, "|") != 2 {
				continue
			}
			lvl := strings.Index(i, "|")
			mes := strings.LastIndex(i, "|")
			results <- Result{
				Level:   i[:lvl],
				Message: i[mes:],
			}
		}
		close(results)
	}()

	for i := range results {
		if strings.Contains(i.Level, "ERROR") {
			errors++
		}
	}

	fmt.Println(float32(errors) / float32(len(logs)))
}
