package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"reflect"
	"strings"
)

// MaskSecrets recursively masks fields with the `mask:"true"` tag
func MaskSecrets(config interface{}) string {
	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	masked := maskValue(v)
	marshaled, err := yaml.Parser().Marshal(masked)
	if err != nil {
		return ""
	}

	return string(marshaled)
}

// maskValue applies masking to primitive types with custom rules
func maskValue(v reflect.Value) map[string]interface{} {
	masked := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		if fieldType.Tag.Get("secret") == "true" {
			masked[fieldName] = applyMask(fmt.Sprintf("%v", field.Interface()))
		} else if field.Kind() == reflect.Struct {
			masked[fieldName] = maskValue(field)
		} else {
			masked[fieldName] = field.Interface()
		}
	}
	return masked
}

// applyMask applies custom masking rules based on value length
func applyMask(value string) string {
	length := len(value)
	if length == 0 {
		return ""
	}

	switch {
	case length < 4:
		return value[:1] + strings.Repeat("*", 4) + value[length-1:]
	case length <= 10:
		return value[:1] + strings.Repeat("*", 4) + value[length-1:]
	case length <= 20:
		return value[:2] + strings.Repeat("*", 4) + value[length-2:]
	default:
		return value[:3] + strings.Repeat("*", 4) + value[length-3:]
	}
}
