package version

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestFormat(t *testing.T) {
	format := Format("dev", "10/10/2020")
	assert.Equal(t, "dev 10/10/2020\n", format)
}

func TestNewCmdVersion(t *testing.T) {
	output := captureOutput(func() {
		cmd := NewCmdVersion("dev", "10/10/2020")
		cmd.Execute()
	})
	assert.Equal(t, "dev 10/10/2020\n\n", output)

}
