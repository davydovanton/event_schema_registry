package schemaregistry

import (
	"fmt"
	"strings"
)

type Loader struct {
	schemasRootPath string
}

func NewLoader(schemasRootPath string) *Loader {
	return &Loader{
		schemasRootPath: schemasRootPath,
	}
}

func (l *Loader) SchemaPath(name string, version int) string {
	name = strings.ReplaceAll(name, ".", "/")
	return fmt.Sprintf("file://%s/%s/%d.json", l.schemasRootPath, name, version)
}
