package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func Test_validateWebOrAPI(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "Web passes",
			value:         "Web",
			expectedDiags: nil,
		},
		{
			name:          "API passes",
			value:         "API",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'Web' or 'API' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateWebOrAPI(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateRuleOrRuleSet(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "Rule passes",
			value:         "Rule",
			expectedDiags: nil,
		},
		{
			name:          "RuleSet passes",
			value:         "RuleSet",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'Rule' or 'RuleSet' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateRuleOrRuleSet(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "SuccessIfAllSucceed passes",
			value:         "SuccessIfAllSucceed",
			expectedDiags: nil,
		},
		{
			name:          "SuccessIfAnyOneSucceeds passes",
			value:         "SuccessIfAnyOneSucceeds",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'SuccessIfAllSucceed' or 'SuccessIfAnyOneSucceeds' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateSuccessIfAllSucceedOrSuccessIfAnyOneSucceeds(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateAudience(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "Normal value passes",
			value:         "normal",
			expectedDiags: nil,
		},
		{
			name:          "Empty value failes",
			value:         "",
			expectedDiags: diag.Errorf("must be between 1 and 32 characters not 0"),
		},
		{
			name:          "Very long value fails",
			value:         "12345678901234567890123456789012",
			expectedDiags: diag.Errorf("must be between 1 and 32 characters not 32"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateAudience(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateCookieType(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "Encrypted passes",
			value:         "Encrypted",
			expectedDiags: nil,
		},
		{
			name:          "Signed passes",
			value:         "Signed",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'Encrypted' or 'Signed' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateCookieType(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateOidcLoginType(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "Code passes",
			value:         "Code",
			expectedDiags: nil,
		},
		{
			name:          "POST passes",
			value:         "POST",
			expectedDiags: nil,
		},
		{
			name:          "x_post passes",
			value:         "x_post",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'Code', 'POST' or 'x_post' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateOidcLoginType(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateRequestPreservationType(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "None passes",
			value:         "None",
			expectedDiags: nil,
		},
		{
			name:          "POST passes",
			value:         "POST",
			expectedDiags: nil,
		},
		{
			name:          "All passes",
			value:         "All",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'None', 'POST' or 'All' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateRequestPreservationType(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateWebStorageType(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "SessionStorage passes",
			value:         "SessionStorage",
			expectedDiags: nil,
		},
		{
			name:          "LocalStorage passes",
			value:         "LocalStorage",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'SessionStorage' or 'LocalStorage' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateWebStorageType(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateListLocationValue(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "FIRST passes",
			value:         "FIRST",
			expectedDiags: nil,
		},
		{
			name:          "LAST passes",
			value:         "LAST",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'FIRST' or 'LAST' not other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateListLocationValue(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}
