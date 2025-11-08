package util

import (
	"encoding/binary"
	"fmt"
	"io"
)

func WriteBigEndianUint16(w io.Writer, b uint16) error {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, b)

	if _, err := w.Write(bytes); err != nil {
		return fmt.Errorf("write %d to %s as big endian 2 bytes: %w", b, w, err)
	}

	return nil
}
