package compton

import (
	"bufio"
	"bytes"
	"io"
)

// separators must have len == 2
const sepLength = 2

var separators = []string{
	"{{",
	"}}",
	"/*",
	"*/",
}

type TokenWriter func(token string, w io.Writer) error

func scanSeparators(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	indexes := make([]int, len(separators))
	for i, sep := range separators {
		indexes[i] = bytes.Index(data, []byte(sep))
	}

	minIndex := len(data)
	for _, index := range indexes {
		if index >= 0 && index < minIndex {
			minIndex = index
		}
	}

	if minIndex >= 0 && minIndex < len(data) {
		return minIndex + sepLength, data[0:minIndex], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func WriteContents(r io.Reader, w io.Writer, wd TokenWriter) error {

	s := bufio.NewScanner(r)
	s.Split(scanSeparators)
	for s.Scan() {
		sbts := s.Bytes()
		bts := sbts

		if bytes.HasPrefix(sbts, []byte{'.'}) {
			if err := wd(string(sbts), w); err != nil {
				return err
			}
			continue
		}

		if _, err := w.Write(bts); err != nil {
			return err
		}
	}

	return nil
}
