package template

import (
	"strings"
	"text/template"

	"github.com/kelseyhightower/envconfig"
)

var (
	functions = map[string]interface{}{
		"date": date,
		"env":  env,
		"now":  now,
	}
)

func Execute(text string) (string, error) {
	tmpl, err := template.New("").Funcs(functions).Parse(text)
	if err != nil {
		return "", err
	}

	meta := &Metadata{}
	err = envconfig.Process("", meta)
	if err != nil {
		return "", err
	}

	builder := &strings.Builder{}
	err = tmpl.Execute(builder, meta)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}
