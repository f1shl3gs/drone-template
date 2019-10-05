# drone-template-lib
library for generating templates in drone plugin developing

## Usage
Official Golang Library is used, if you are not familiar with that,
please visit https://godoc.org/text/template

```

import (
 "github.com/f1shl3gs/drone-template-lib/template"
)

func YourFunction() {
    key, err := template.Execute("the repo name is {{ .Repo.Name }}")
    // handle error
    // do something 
}
```

## functions
- `date`: convert unix timestamp to date, timezone(eg: 'Asia/Shanghai') and format is optional
    - timestamp: required, unix second
    - timezone:  optional, eg: 'Asia/Shanghai'
    - format:    optional, if `format` specified, timezone become required, empty value equal to "Local"
- `now`:  generate a date string of now
    - timezone:  optional, eg: 'Asia/Shanghai'
    - format:    optional, if `format` specified, timezone become required, empty value equal to "Local"  

### quote
line break might break json format(https://github.com/drone-plugins/drone-webhook/issues/34),
you can fix it by `quote` which is a built-in function of text/template  