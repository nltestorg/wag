package swagger

import (
	"fmt"
	"strings"

	"github.com/go-openapi/spec"
)

// SingleSchemaedBodyParameter returns true if the operation has a single,
// schema'd body input. It also returns the name of the model (without "models.").
func SingleSchemaedBodyParameter(op *spec.Operation) (bool, string) {
	if len(op.Parameters) == 1 && op.Parameters[0].ParamProps.Schema != nil {
		typ, err := TypeFromSchema(op.Parameters[0].ParamProps.Schema, false)
		if err != nil {
			panic(err)
		}
		return true, typ
	}
	return false, ""
}

// Interface returns the interface for an operation
func Interface(op *spec.Operation) string {
	capOpID := Capitalize(op.ID)

	// Don't add the input parameter argument unless there are some arguments.
	// If a method has a single input parameter, and it's a schema, make the
	// generated type for that schema the input of the method.
	// TODO: If a method has a single input parameter, and it's a simple type (int/string),
	// make that the input of the method.
	// If a method has multiple input parameters, wrap them in a struct.
	input := ""
	if ssbp, opModel := SingleSchemaedBodyParameter(op); ssbp {
		input = fmt.Sprintf("i *models.%s", opModel)
	} else if len(op.Parameters) > 0 {
		input = fmt.Sprintf("i *models.%sInput", capOpID)
	}

	if NoSuccessType(op) {
		return fmt.Sprintf("%s(ctx context.Context, %s) error", capOpID, input)
	}

	successCodes := SuccessStatusCodes(op)
	successType := ""

	if len(successCodes) == 1 {
		singleSchema := op.Responses.StatusCodeResponses[successCodes[0]].Schema
		var err error
		successType, err = TypeFromSchema(singleSchema, true)
		if err != nil {
			panic(fmt.Errorf("Could not convert operation to type %s", err))
		}
		// Make non-arrays pointers
		if singleSchema.Ref.String() != "" {
			successType = "*" + successType
		}
	} else {
		successType = fmt.Sprintf("models.%sOutput", capOpID)
	}

	return fmt.Sprintf("%s(ctx context.Context, %s) (%s, error)",
		capOpID, input, successType)
}

// InterfaceComment returns the comment for the interface for the operation
func InterfaceComment(method, path string, op *spec.Operation) string {

	capOpID := Capitalize(op.ID)
	comment := fmt.Sprintf("// %s makes a %s request to %s.\n", capOpID, method, path)
	if op.Description != "" {
		comment += "// " + op.Description
	}
	return comment
}

// OutputType returns the output type for a given status code of an operation
func OutputType(op *spec.Operation, statusCode int) string {
	successCodes := SuccessStatusCodes(op)
	if len(successCodes) == 1 && statusCode < 400 {
		successType, err := TypeFromSchema(op.Responses.StatusCodeResponses[successCodes[0]].Schema, true)
		if err != nil {
			panic(fmt.Errorf("Could not convert operation to type %s", err))
		}
		return successType
	}
	return fmt.Sprintf("models.%s%dOutput", Capitalize(op.ID), statusCode)
}

// SUccessStatusCodes returns a slice of all the success status codes for an operation
func SuccessStatusCodes(op *spec.Operation) []int {
	successCodes := make([]int, 0)
	for statusCode, _ := range op.Responses.StatusCodeResponses {
		if statusCode < 400 {
			successCodes = append(successCodes, statusCode)
		}
	}
	return successCodes
}

// NoSuccessType returns true if the operation has no-success response type. This includes
// either no 200-399 response code or a 200-399 response code without a schema.
func NoSuccessType(op *spec.Operation) bool {
	successCodes := SuccessStatusCodes(op)
	if len(successCodes) > 1 {
		return false
	}
	if len(successCodes) == 0 {
		return true
	}
	return op.Responses.StatusCodeResponses[successCodes[0]].Schema == nil
}

// TypeFromSchema returns the type of a Swagger schema as a string. If includeModels is true
// then it returns models.TYPE
func TypeFromSchema(schema *spec.Schema, includeModels bool) (string, error) {
	// We support three types of schemas
	// 1. An empty schema, which we represent by the empty struct
	// 2. A schema with one element, the $ref key
	// 3. A schema with two elements. One a type with value 'array' and another items field
	// referencing the $ref
	if schema == nil {
		return "struct{}", nil
	} else if schema.Ref.String() != "" {
		ref := schema.Ref.String()
		if !strings.HasPrefix(ref, "#/definitions/") {
			return "", fmt.Errorf("schema.$ref has undefined reference type. Must start with #/definitions")
		}
		def := ref[len("#/definitions/"):]
		if includeModels {
			def = "models." + def
		}
		return def, nil
	} else {
		schemaType := schema.Type
		if len(schemaType) != 1 || schemaType[0] != "array" {
			return "", fmt.Errorf("Two element schemas must have a 'type' field with the value 'array'")
		}
		items := schema.Items
		if items == nil || items.Schema == nil {
			return "", fmt.Errorf("Two element schemas must have a '$ref' field in the 'items' descriptions")
		}
		ref := items.Schema.Ref.String()
		if !strings.HasPrefix(ref, "#/definitions/") {
			return "", fmt.Errorf("schema.$ref has undefined reference type. Must start with #/definitions")
		}
		def := ref[len("#/definitions/"):]
		if includeModels {
			def = "models." + def
		}
		return "[]" + def, nil
	}
}
