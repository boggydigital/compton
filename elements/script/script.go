package script

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
)

var (
	//go:embed "markup/script.html"
	markupScript []byte
)

var b64 = base64.StdEncoding

type ScriptElement struct {
	compton.BaseElement
	hash []byte
}

func computeSha256(reader io.Reader) ([]byte, error) {
	h := sha256.New()
	var err error
	if _, err = io.Copy(h, reader); err == nil {
		return h.Sum(nil), nil
	}
	return nil, err
}

func (se *ScriptElement) Sha256() string {
	if len(se.hash) > 0 {
		return "sha256-" + b64.EncodeToString(se.hash)
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
