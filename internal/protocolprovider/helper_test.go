package protocol

import (
	"math/big"
	"testing"

	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_marshalDynamicPseudoTypes(t *testing.T) {
	type testCase struct {
		tfval tftypes.Value
		input interface{}
	}
	cases := map[string]testCase{
		"string": {
			tfval: tftypes.NewValue(tftypes.String, "foo"),
			input: "foo",
		},
		"number": {
			tfval: tftypes.NewValue(tftypes.Number, big.NewFloat(123)),
			input: big.NewFloat(123),
		},
		"bool": {
			tfval: tftypes.NewValue(tftypes.Bool, true),
			input: true,
		},
		"object-bool-string-number": {
			tfval: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.Bool,
					"b": tftypes.String,
					"c": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.Bool, true),
				"b": tftypes.NewValue(tftypes.String, "bar"),
				"c": tftypes.NewValue(tftypes.Number, big.NewFloat(456)),
			}),
			input: map[string]interface{}{
				"a": true,
				"b": "bar",
				"c": big.NewFloat(456),
			},
		},
		"list-number": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Number,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Number, big.NewFloat(1)),
				tftypes.NewValue(tftypes.Number, big.NewFloat(2)),
				tftypes.NewValue(tftypes.Number, big.NewFloat(3)),
				tftypes.NewValue(tftypes.Number, big.NewFloat(4)),
			}),
			input: []*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3), big.NewFloat(4)},
		},
		"list-string": {
			tfval: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "foo"),
				tftypes.NewValue(tftypes.String, "bar"),
				tftypes.NewValue(tftypes.String, "baz"),
				tftypes.NewValue(tftypes.String, "quux"),
			}),
			input: []string{"foo", "bar", "baz", "quux"},
		},
		"tuple-bool-string-number": {
			tfval: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{
					tftypes.Bool,
					tftypes.String,
					tftypes.Number,
				},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.String, "test"),
				tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
			}),
			input: []interface{}{true, "test", big.NewFloat(1234)},
		},
		"map-string": {
			tfval: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.String, "foo"),
				"b": tftypes.NewValue(tftypes.String, "bar"),
				"c": tftypes.NewValue(tftypes.String, "baz"),
			}),
			input: map[string]string{
				"a": "foo",
				"b": "bar",
				"c": "baz",
			},
		},
		"object with nil": {
			input: map[string]interface{}{
				"description": nil,
				"path":        "/foo/bar",
			},
			tfval: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					//"description": tftypes.String,
					"path": tftypes.String,
				},
			}, map[string]tftypes.Value{
				//"description": tftypes.NewValue(nil ?),
				"path": tftypes.NewValue(tftypes.String, "/foo/bar"),
			}),
		},
	}
	for name, testCase := range cases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, v, err := marshal(testCase.input)
			require.Nil(t, err)
			assert.Equal(t, testCase.tfval, v)
		})
	}
}

func Test_maskConfigFromDescriptors(t *testing.T) {
	type testCase struct {
		descriptors             *models.DescriptorsView
		input, config, expected string
	}
	cases := map[string]testCase{
		"password is preserved": {
			descriptors: &models.DescriptorsView{
				Items: []*models.DescriptorView{
					{
						ClassName: String("something"),
						ConfigurationFields: []*models.ConfigurationField{
							{
								Name: String("password"),
								Type: String("CONCEALED"),
							},
						},
						Label: nil,
						Type:  nil,
					}}},
			input:    "{\"password\":{\"encryptedValue\":\"OBF:JWE:eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2Iiwia2lkIjoiQTJxVXBQRjc3bjJqRFNhRiJ9..sdvO4RzS3fbVEUPSK9mNvQ.R6q-ulIt-npaAotMBImmvQ.8Xw9sLg5qgN6xPQ_Bfvxng\"},\"username\":\"cheese\"}",
			config:   "{\"password\":{\"value\":\"secret\"},\"username\":\"cheese\"}",
			expected: "{\"password\":{\"value\":\"secret\"},\"username\":\"cheese\"}",
		},
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expected, maskConfigFromDescriptors(testCase.descriptors, "something", testCase.input, testCase.config))
		})
	}
}
