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

func Test_validateRuleOrRuleSet(t *testing.T) {
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
			name: "Rule passes",
			args: args{
				value: "Rule",
				field: "elementType",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "RuleSet passes",
			args: args{
				value: "RuleSet",
				field: "elementType",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "elementType",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'Rule' or 'RuleSet' not %s", "elementType", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateRuleOrRuleSet(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateRuleOrRuleSet() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateRuleOrRuleSet() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(t *testing.T) {
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
			name: "SuccessIfAllSucceed passes",
			args: args{
				value: "SuccessIfAllSucceed",
				field: "successCriteria",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "SuccessIfAnyOneSucceeds passes",
			args: args{
				value: "SuccessIfAnyOneSucceeds",
				field: "successCriteria",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "successCriteria",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'SuccessIfAllSucceed' or 'SuccessIfAnyOneSucceeds' not %s", "successCriteria", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
