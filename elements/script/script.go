package script

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
)

var (
	//go:embed "markup/script.html"
	markupScript []byte
)

var b64 = base64.RawURLEncoding

type ScriptElement struct {
	compton.BaseElement
	hash string
}

func computeSha256(reader io.Reader) (string, error) {
	h := sha256.New()
	var err error
	if _, err = io.Copy(h, reader); err == nil {
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	}
	return "", err
}

func (se *ScriptElement) Sha256() string {
	if se.hash != "" {
		return "sha256-" + b64.EncodeToString([]byte(se.hash))
	}
	return ""
}

func Script(code []byte) *ScriptElement {

	script := &ScriptElement{
		BaseElement: compton.BaseElement{
			TagName: atom.Script,
			Markup:  markupScript,
		},
	}

	if hash, err := computeSha256(bytes.NewReader(code)); err == nil {
		script.hash = hash
	}

	script.Append(els.Text(string(code)))
	return script
}

func ScriptAsync(code []byte) *ScriptElement {
	script := Script(code)
	script.SetAttribute("async", "")
	return script
}
