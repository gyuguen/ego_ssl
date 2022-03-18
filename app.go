package main

import (
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	schemaURI := "https://json.schemastore.org/github-issue-forms.json"
	jsonInput := []byte(`{
		"name": "This is a name",
		"description": "This is a description, man",
		"body": [{ "type": "markdown", "attributes": { "value": "val1" } }]
	}`)

	fmt.Println("execute validation")
	err := validateJSONSchema(jsonInput, schemaURI)
	if err != nil {
		panic(err)
	}
	fmt.Println("finish validation")
}

// ValidateJSONSchema performs the JSON Schema validation: https://json-schema.org/
// This fetches the schema definition via http(s) or local filesystem.
// If jsonInput is not a valid JSON or if jsonInput doesn't conform to the desired JSON schema, an error is returned.
//
// TODO: accept io.Reader instead of []byte
func validateJSONSchema(jsonInput []byte, desiredSchemaURI string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(desiredSchemaURI)
	docLoader := gojsonschema.NewBytesLoader(jsonInput)

	result, err := gojsonschema.Validate(schemaLoader, docLoader)
	if err != nil {
		return fmt.Errorf("failed to validate JSON schema: %w", err)
	}

	if !result.Valid() {
		var sb strings.Builder
		for _, err := range result.Errors() {
			sb.WriteString("\n\t")
			sb.WriteString(err.String())
		}
		return fmt.Errorf("JSON doc doesn't conform to the desired JSON schema: %s", sb.String())
	}

	return nil
}
