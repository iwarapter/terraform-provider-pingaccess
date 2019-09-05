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

func Test_validateAudience(t *testing.T) {
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
			name: "Normal value passes",
			args: args{
				value: "normal",
				field: "audience",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "Empty value fails",
			args: args{
				value: "",
				field: "audience",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be between 1 and 32 characters not %s", "audience", "0")},
		},
		{
			name: "Very long value fails",
			args: args{
				value: "12345678901234567890123456789012",
				field: "audience",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be between 1 and 32 characters not %s", "audience", "32")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateAudience(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateAudience() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateAudience() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateCookieType(t *testing.T) {
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
			name: "Encrypted passes",
			args: args{
				value: "Encrypted",
				field: "cookie_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "Signed passes",
			args: args{
				value: "Signed",
				field: "cookie_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "cookie_type",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'Encrypted' or 'Signed' not %s", "cookie_type", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateCookieType(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateCookieType() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateCookieType() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateOidcLoginType(t *testing.T) {
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
			name: "Code passes",
			args: args{
				value: "Code",
				field: "oidc_login_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "POST passes",
			args: args{
				value: "POST",
				field: "oidc_login_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "x_post passes",
			args: args{
				value: "x_post",
				field: "oidc_login_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "oidc_login_type",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'Code', 'POST' or 'x_post' not %s", "oidc_login_type", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateOidcLoginType(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateOidcLoginType() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateOidcLoginType() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateRequestPreservationType(t *testing.T) {
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
			name: "None passes",
			args: args{
				value: "None",
				field: "refresh_user_info_claims_interval",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "POST passes",
			args: args{
				value: "POST",
				field: "refresh_user_info_claims_interval",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "All passes",
			args: args{
				value: "All",
				field: "refresh_user_info_claims_interval",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "refresh_user_info_claims_interval",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'None', 'POST' or 'All' not %s", "refresh_user_info_claims_interval", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateRequestPreservationType(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateRequestPreservationType() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateRequestPreservationType() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateWebStorageType(t *testing.T) {
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
			name: "SessionStorage passes",
			args: args{
				value: "SessionStorage",
				field: "web_storage_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "LocalStorage passes",
			args: args{
				value: "LocalStorage",
				field: "web_storage_type",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "web_storage_type",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'SessionStorage' or 'LocalStorage' not %s", "web_storage_type", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateWebStorageType(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateWebStorageType() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateWebStorageType() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateListLocationValue(t *testing.T) {
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
			name: "FIRST passes",
			args: args{
				value: "FIRST",
				field: "list_value_location",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "LAST passes",
			args: args{
				value: "LAST",
				field: "list_value_location",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "list_value_location",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'FIRST' or 'LAST' not %s", "list_value_location", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateListLocationValue(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateListLocationValue() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateListLocationValue() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateClassNameValue(t *testing.T) {
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
			name: "CookieBasedClassName passes",
			args: args{
				value: cookieBasedClassName,
				field: "className",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "HeaderBasedClassName passes",
			args: args{
				value: headerBasedClassName,
				field: "className",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "className",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%[1]q must be either %[2]s or %[3]s not %[4]s", "className", cookieBasedClassName, headerBasedClassName, "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateClassNameValue(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateClassNameValue() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateClassNameValue() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
