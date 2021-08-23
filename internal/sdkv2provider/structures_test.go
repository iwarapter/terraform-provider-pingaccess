package sdkv2provider

import (
	"encoding/json"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
)

func testHiddenFieldView() map[string]interface{} {
	return map[string]interface{}{
		"encrypted_value": "atat",
		"value":           "atat",
	}
}

func Test_weCanFlattenHiddenFieldView(t *testing.T) {
	initialHiddenFieldView := &models.HiddenFieldView{
		Value:          String("atat"),
		EncryptedValue: String("atat"),
	}

	output := []map[string]interface{}{{"encrypted_value": "atat", "value": "atat"}}

	flattened := flattenHiddenFieldView(initialHiddenFieldView)

	equals(t, output, flattened)
}

func Test_expandHiddenFieldView(t *testing.T) {
	expanded := []interface{}{testHiddenFieldView()}
	expandHiddenFieldView := expandHiddenFieldView(expanded)

	equals(t, "atat", *(*expandHiddenFieldView).Value)
	equals(t, "atat", *(*expandHiddenFieldView).EncryptedValue)
}

func testOAuthClientCredentials() map[string]interface{} {
	return map[string]interface{}{
		"client_id": "atat",
		"client_secret": []interface{}{
			map[string]interface{}{
				"encrypted_value": "atat",
				"value":           "atat",
			},
		},
	}
}

func Test_weCanFlattenOAuthClientCredentials(t *testing.T) {
	initialOAuthClientCredentialsView := &models.OAuthClientCredentialsView{
		ClientId: String("atat"),
		ClientSecret: &models.HiddenFieldView{
			Value:          String("atat"),
			EncryptedValue: String("atat"),
		},
	}

	output := []map[string]interface{}{{"client_id": "atat", "client_secret": []map[string]interface{}{{"value": "atat", "encrypted_value": "atat"}}}}

	flattened := flattenOAuthClientCredentialsView(initialOAuthClientCredentialsView)

	equals(t, output, flattened)
}

func Test_expandOAuthClientCredentials(t *testing.T) {
	expanded := []interface{}{testOAuthClientCredentials()}
	expandOAuthClientCredentialsView := expandOAuthClientCredentialsView(expanded)

	equals(t, "atat", *(*expandOAuthClientCredentialsView).ClientId)
	equals(t, "atat", *(*expandOAuthClientCredentialsView).ClientSecret.Value)
}

func testPolicyItem() map[string]interface{} {
	return map[string]interface{}{
		"id":   "1334",
		"type": "Rule",
	}
}

func Test_weCanFlattenPolicy(t *testing.T) {
	initialPolicyItem := []*models.PolicyItem{
		{
			Id:   json.Number("1334"),
			Type: String("Rule"),
		},
	}

	output := []interface{}{map[string]interface{}{"id": "1334", "type": "Rule"}}

	flattened := flattenPolicyItem(initialPolicyItem)

	equals(t, output, flattened)
}

func Test_expandPolicyItem(t *testing.T) {
	expanded := []interface{}{testPolicyItem()}
	expandPolicyItem := expandPolicyItem(expanded)

	equals(t, "1334", expandPolicyItem[0].Id.String())
	equals(t, "Rule", *(*expandPolicyItem[0]).Type)
}

func testPolicy() []interface{} {
	return []interface{}{map[string]interface{}{"api": []interface{}{map[string]interface{}{"id": "1334", "type": "Rule"}}}}
}

func Test_expandPolicy(t *testing.T) {
	expandPolicyItem := expandPolicy(testPolicy())

	api := *(expandPolicyItem)["API"]
	equals(t, "1334", api[0].Id.String())
}

func Test_maskConfigFromConfigurationField(t *testing.T) {
	tests := []struct {
		name           string
		field          *models.ConfigurationField
		config         string
		originalConfig string
		want           string
	}{
		{
			name:           "we can mask a password",
			field:          &models.ConfigurationField{Name: String("password"), Type: String("CONCEALED")},
			config:         "{\"library\":\"foo\",\"password\":{\"encryptedValue\":\"OBF:JWE:eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2Iiwia2lkIjoiVGhNT3FrdTZ0cHZjdFNEaCJ9..tyXmXFR76lYnHrfwLkiLPw.sFqoXnRjyVllwKv80rHTbQ.4vuVPhY_l5KXA-w_bZbrcQ\"},\"slotId\":\"1234\"}",
			originalConfig: "\t  {\n\t\t\"slotId\": \"1234\",\n\t\t\"library\": \"foo\",\n\t\t\"password\": \"top_secret\"\n\t  }\n",
			want:           "{\"library\":\"foo\",\"password\":\"top_secret\",\"slotId\":\"1234\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskConfigFromConfigurationField(tt.field, String(""), tt.originalConfig, tt.config); got != tt.want {
				t.Errorf("maskConfigFromConfigurationField() = %v, want %v", got, tt.want)
			}
		})
	}
}
