package pingaccess

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_validateWebOrAPI(t *testing.T) {
	type args struct {
		value interface{}
		field string
	}
	tests := []struct {
		name      string
		args      args
		wantWarns []string
		wantErrs  []error
	}{
		{
			name: "Web passes",
			args: args{
				value: "Web",
				field: "default_auth_type_override",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "API passes",
			args: args{
				value: "API",
				field: "default_auth_type_override",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "default_auth_type_override",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'Web' or 'API' not %s", "default_auth_type_override", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateWebOrAPI(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateWebOrAPI() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateWebOrAPI() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
