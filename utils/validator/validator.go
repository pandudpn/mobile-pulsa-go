package validator

import (
	"context"
	"fmt"
	"strings"

	gpv "github.com/go-playground/validator/v10"
)

var (
	validator   = gpv.New()
	requiredTag = "required"
)

// ValidateRequired validate fields that are required in the struct with `valiate:"required"` tag
func ValidateRequired(ctx context.Context, value interface{}) error {
	err := validator.StructCtx(ctx, value)
	errs, ok := err.(gpv.ValidationErrors)
	if !ok {
		return err
	}

	var missingFields []string
	for _, e := range errs {
		if e.Tag() == requiredTag {
			missingFields = append(missingFields, e.Field())
		}
	}
	return errorMsgFromMissingFields(missingFields)
}

func errorMsgFromMissingFields(fields []string) error {
	return fmt.Errorf("missing required fields: %s", strings.Join(fields, "', '"))
}
