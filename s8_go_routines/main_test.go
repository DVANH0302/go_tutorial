package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_PrintSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		printSomething("test")
	}()

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "test") {
		t.Errorf("Expected to find test but it is not there")
	}
}
