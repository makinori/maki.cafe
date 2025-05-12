package render

import (
	"fmt"
	"io"

	"maragu.dev/gomponents"
)

func EscapedHTML(input string) string {
	var output string
	for i := range len(input) {
		output += fmt.Sprintf("&#%d;", input[i])
	}
	return output
}

type AttrRaw struct {
	Name  string
	Value string
}

// gomponents attr but no html escaping
func (a *AttrRaw) Render(w io.Writer) error {
	_, err := w.Write([]byte(" " + a.Name + `="` + a.Value + `"`))
	return err
}

func (a *AttrRaw) Type() gomponents.NodeType {
	return gomponents.AttributeType
}
