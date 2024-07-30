package compton

import (
	"bufio"
	"bytes"
	"io"
)

var (
	pfxCurlyBraces = []byte("{{")
	sfxCurlyBraces = []byte("}}")
)

type TokenWriter func(token string, w io.Writer) error

func scanCurlyBraces(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	oi, ci := bytes.Index(data, pfxCurlyBraces), bytes.Index(data, sfxCurlyBraces)

	if oi >= 0 || ci >= 0 {
		i := min(oi, ci)
		if oi < 0 {
			i = ci
		}
		if ci < 0 {
			i = oi
		}
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func WriteContents(r io.Reader, w io.Writer, wd TokenWriter) error {

	s := bufio.NewScanner(r)
	s.Split(scanCurlyBraces)
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

//type stringWriter string
//
//func (sw stringWriter) write(_ string, w io.Writer) error {
//	_, err := io.WriteString(w, string(sw))
//	return err
//}
//
//func WriteString(r io.Reader, w io.Writer, s string) error {
//	return WriteContents(r, w, stringWriter(s).write)
//}
