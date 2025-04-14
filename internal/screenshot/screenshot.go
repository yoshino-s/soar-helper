package screenshot

import (
	"bytes"
	"io"
)

func CreateScreenshot(
	command []string,
	cols int,
	result string,
	writer io.Writer,
) error {
	scaffold := NewImageCreator()

	scaffold.SetColumns(cols)
	scaffold.DrawShadow(true)
	scaffold.DrawDecorations(true)
	scaffold.ClipCanvas(false)

	scaffold.AddCommand(command...)

	scaffold.AddContent(bytes.NewReader([]byte(result)))
	scaffold.Write(writer)

	return nil
}
