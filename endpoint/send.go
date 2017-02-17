package endpoint

import (
	"bufio"
	"errors"

	common "github.com/ekhabarov/go-common"
)

var ErrBufferIsNil = errors.New("endpoint: send: rw is nil")

func send(rw *bufio.ReadWriter, data string) error {
	if rw == nil {
		return ErrBufferIsNil
	}

	// Send data
	_, err := rw.WriteString(data + "\n")
	common.LogIf(err, "endpoint: send: write data:")

	err = rw.Flush()
	common.LogIf(err, "endpoint: send: flush:")

	return nil
}
