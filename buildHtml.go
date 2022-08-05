package markdownFolderMap

import (
	"bytes"
	"html/template"
)

func buildHtml(temple string, payload map[string]map[string]interface{}) (string, error) {
	var outBytes bytes.Buffer
	t1 := template.New("te")
	t1 = template.Must(t1.Parse(temple))
	err := t1.Execute(&outBytes, payload)
	return outBytes.String(), err
}
