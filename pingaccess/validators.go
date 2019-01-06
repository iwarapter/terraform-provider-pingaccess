package pingaccess

import (
	"fmt"
)

func validateWebOrAPI(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	if v != "Web" && v != "API" {
		errs = append(errs, fmt.Errorf("%q must be either 'Web' or 'API' not %s", field, v))
	}
	return
}
