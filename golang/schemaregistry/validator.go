package schemaregistry

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	loader *Loader
}

func NewValidator(path string) *Validator {
	return &Validator{
		loader: NewLoader(path),
	}
}

func (v *Validator) Validate(data json.RawMessage, name string, version int) error {
	schemaPath := v.loader.SchemaPath(name, version)
	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return err
	}
	dataLoader := gojsonschema.NewStringLoader(string(data))
	result, err := schema.Validate(dataLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		errs := "json schema validation failed"
		for _, desc := range result.Errors() {
			errs = fmt.Sprintf("%s - %s", errs, desc)
		}
		return errors.New(errs)
	}

	return nil
}
