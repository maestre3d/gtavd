package xmlutil

import (
	"encoding/xml"
	"os"
)

func WriteFile(filePath string, v any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = file.WriteString("\n")
		_ = file.Close()
	}()

	_, _ = file.WriteString(xml.Header)
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	return encoder.Encode(v)
}
