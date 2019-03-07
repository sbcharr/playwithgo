package interfaces

import (
	"io"
	"os"
)

// PipeExample ises io interface
func PipeExample() error {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("hello world!"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil

}
