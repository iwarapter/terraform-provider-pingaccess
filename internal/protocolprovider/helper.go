package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"

	"github.com/tidwall/sjson"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingaccess-sdk-go/v62/pingaccess/models"
	"github.com/tidwall/gjson"
)

const unexpectedConfigurationFormat = "Unexpected configuration format"

func dynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{schemaDataSourceMistmatchDiagnostic(err)}
	}
	if !val.Type().Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  unexpectedConfigurationFormat,
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{schemaDataSourceMistmatchDiagnostic(err)}
	}
	return values, nil
}

func resourceDynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}
	}
	if !val.Type().Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  unexpectedConfigurationFormat,
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}
	}
	return values, nil
}

func valuesFromTypeConfigRequest(req *tfprotov5.ValidateResourceTypeConfigRequest, typ tftypes.Type) (*tfprotov5.ValidateResourceTypeConfigResponse, map[string]tftypes.Value) {
	val, err := req.Config.Unmarshal(typ)
	if err != nil {
		return &tfprotov5.ValidateResourceTypeConfigResponse{Diagnostics: []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}}, nil
	}
	if !val.Type().Is(typ) {
		return &tfprotov5.ValidateResourceTypeConfigResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  unexpectedConfigurationFormat,
					Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.",
				},
			},
		}, nil
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return &tfprotov5.ValidateResourceTypeConfigResponse{Diagnostics: []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}}, nil
	}
	return nil, values
}

func readResourceChangeError(err error) *tfprotov5.ReadResourceResponse {
	return &tfprotov5.ReadResourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  unexpectedConfigurationFormat,
				Detail:   err.Error(),
			},
		},
	}
}

func planResourceChangeError(err error) *tfprotov5.PlanResourceChangeResponse {
	return &tfprotov5.PlanResourceChangeResponse{Diagnostics: []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}}
}

func applyResourceChangeError(err error) *tfprotov5.ApplyResourceChangeResponse {
	return &tfprotov5.ApplyResourceChangeResponse{Diagnostics: []*tfprotov5.Diagnostic{schemaResourceMistmatchDiagnostic(err)}}
}

func importResourceError(detail string) *tfprotov5.ImportResourceStateResponse {
	return &tfprotov5.ImportResourceStateResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Error importing resource",
				Detail:   detail,
			},
		},
	}
}

func schemaResourceMistmatchDiagnostic(err error) *tfprotov5.Diagnostic {
	return &tfprotov5.Diagnostic{
		Severity: tfprotov5.DiagnosticSeverityError,
		Summary:  unexpectedConfigurationFormat,
		Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
	}
}

func schemaDataSourceMistmatchDiagnostic(err error) *tfprotov5.Diagnostic {
	return &tfprotov5.Diagnostic{
		Severity: tfprotov5.DiagnosticSeverityError,
		Summary:  unexpectedConfigurationFormat,
		Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
	}
}

func stateEncodingDiagnostic(err error) *tfprotov5.Diagnostic {
	return &tfprotov5.Diagnostic{
		Severity: tfprotov5.DiagnosticSeverityError,
		Summary:  "Error encoding state",
		Detail:   fmt.Sprintf("Error encoding state: %s", err.Error()),
	}
}

func marshal(i interface{}) (tftypes.Type, tftypes.Value, error) {
	v := reflect.ValueOf(i)
	//primitve check by type
	switch i.(type) {
	case *big.Float:
		return tftypes.Number, tftypes.NewValue(tftypes.Number, i), nil
	case string:
		return tftypes.String, tftypes.NewValue(tftypes.String, i), nil
	case bool:
		return tftypes.Bool, tftypes.NewValue(tftypes.Bool, i), nil
	}
	//non-primitive
	switch v.Kind() {
	case reflect.Slice:
		t := reflect.TypeOf(i)
		if t.Elem().Kind() == reflect.Interface {
			//tuple
			return marshalTuple(i)
		} else {
			//list
			return marshalSlice(i)
		}
	case reflect.Map:
		t := reflect.TypeOf(i)
		if t.Elem().Kind() == reflect.Interface {
			//object
			return marshalObject(i)
		} else {
			//map
			return marshalMap(i)
		}
	}
	return nil, tftypes.Value{}, nil
}

//handles object(map[string]interface where subtypes are different)
func marshalObject(i interface{}) (tftypes.Type, tftypes.Value, error) {
	vals := map[string]tftypes.Value{}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Map {
		return nil, tftypes.Value{}, fmt.Errorf("should only be called with map")
	}
	s := reflect.ValueOf(i)
	subTypes := map[string]tftypes.Type{}
	for _, value := range s.MapKeys() {
		if s.MapIndex(value).Interface() == nil {
			continue
		}
		t, v, err := marshal(s.MapIndex(value).Interface())
		if err != nil {
			return nil, tftypes.Value{}, err
		}
		vals[value.String()] = v
		subTypes[value.String()] = t
	}
	return tftypes.Object{AttributeTypes: subTypes}, tftypes.NewValue(tftypes.Object{
		AttributeTypes: subTypes,
	}, vals), nil
}

//handles maps(map[string]string/number/complex type etc where subtypes are consistent)
func marshalMap(i interface{}) (tftypes.Type, tftypes.Value, error) {
	vals := map[string]tftypes.Value{}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Map {
		return nil, tftypes.Value{}, fmt.Errorf("should only be called with map")
	}
	s := reflect.ValueOf(i)
	var typ tftypes.Type
	for _, value := range s.MapKeys() {
		if s.MapIndex(value).Interface() == nil {
			continue
		}
		t, v, err := marshal(s.MapIndex(value).Interface())
		if err != nil {
			return nil, tftypes.Value{}, err
		}
		typ = t
		vals[value.String()] = v
	}

	return tftypes.Map{AttributeType: typ}, tftypes.NewValue(tftypes.Map{
		AttributeType: typ,
	}, vals), nil
}

//handles lists([]string/number/complex type etc)
func marshalSlice(i interface{}) (tftypes.Type, tftypes.Value, error) {
	vals := []tftypes.Value{}
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Slice {
		return nil, tftypes.Value{}, fmt.Errorf("should only be called with slice")
	}
	s := reflect.ValueOf(i)
	var typ tftypes.Type
	for i := 0; i < s.Len(); i++ {
		t, v, err := marshal(s.Index(i).Interface())
		if err != nil {
			return nil, tftypes.Value{}, err
		}
		typ = t
		vals = append(vals, v)
	}

	return tftypes.List{ElementType: typ}, tftypes.NewValue(tftypes.List{
		ElementType: typ,
	}, vals), nil
}

func marshalTuple(i interface{}) (tftypes.Type, tftypes.Value, error) {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Slice {
		return nil, tftypes.Value{}, fmt.Errorf("should only be called with slice")
	}
	vals := []tftypes.Value{}
	s := reflect.ValueOf(i)
	var subTypes []tftypes.Type
	for i := 0; i < s.Len(); i++ {
		t, v, err := marshal(s.Index(i).Interface())
		if err != nil {
			return nil, tftypes.Value{}, err
		}
		vals = append(vals, v)
		subTypes = append(subTypes, t)
	}
	return tftypes.Tuple{ElementTypes: subTypes}, tftypes.NewValue(tftypes.Tuple{
		ElementTypes: subTypes,
	}, vals), nil
}

//Checks all the fields in the descriptor to ensure all required fields are set
//
func descriptorsHasClassName(className string, desc *models.DescriptorsView) *tfprotov5.Diagnostic {
	var classes []string
	for _, value := range desc.Items {
		classes = append(classes, *value.ClassName)
		if *value.ClassName == className {
			return nil
		}
	}
	return &tfprotov5.Diagnostic{
		Severity:  tfprotov5.DiagnosticSeverityError,
		Summary:   "Class Name Validation Failure",
		Detail:    fmt.Sprintf("unable to find className '%s' available classNames: %s", className, strings.Join(classes, ", ")),
		Attribute: &tftypes.AttributePath{},
	}
}

//Checks the class name specified exists in the DescriptorsView
//
func validateConfiguration(className string, configuration asgotypes.GoPrimitive, desc *models.DescriptorsView) []*tfprotov5.Diagnostic {
	var diags []*tfprotov5.Diagnostic
	diags = append(diags, validateNoNullConfigurationAttributes(configuration)...)
	var conf string
	if str, ok := configuration.Value.(string); ok {
		conf = str
	} else {
		b, _ := json.Marshal(configuration.Value)
		conf = string(b)
	}
	if conf == "" {
		log.Println("[INFO] configuration is in a potentially unknown state, gracefully skipping configuration validation")
		return nil
	}
	for _, value := range desc.Items {
		if *value.ClassName == className {
			for _, f := range value.ConfigurationFields {
				if *f.Required {
					v := gjson.Get(conf, *f.Name)
					if !v.Exists() {
						diags = append(diags, &tfprotov5.Diagnostic{
							Severity:  tfprotov5.DiagnosticSeverityError,
							Summary:   "Configuration Validation Failure",
							Detail:    fmt.Sprintf("the field '%s' is required for the class_name '%s'", *f.Name, className),
							Attribute: &tftypes.AttributePath{},
						})
					}
					if !v.IsObject() && v.Str == "" {
						steps := []tftypes.AttributePathStep{
							tftypes.AttributeName("configuration"),
						}
						if _, ok := configuration.Value.(string); !ok {
							steps = append(steps, tftypes.ElementKeyString(*f.Name))
						}

						diags = append(diags, &tfprotov5.Diagnostic{
							Severity:  tfprotov5.DiagnosticSeverityError,
							Summary:   "Configuration Validation Failure",
							Detail:    fmt.Sprintf("the field '%s' is required for the class_name '%s'", *f.Name, className),
							Attribute: tftypes.NewAttributePathWithSteps(steps),
						})
					}
				}
			}
		}
	}
	return diags
}

func validateNoNullConfigurationAttributes(configuration asgotypes.GoPrimitive) []*tfprotov5.Diagnostic {
	var diags []*tfprotov5.Diagnostic
	if v, ok := configuration.Value.(map[string]interface{}); ok {
		for s := range v {
			if v[s] == nil {
				diags = append(diags,
					&tfprotov5.Diagnostic{
						Severity:  tfprotov5.DiagnosticSeverityError,
						Summary:   "Configuration Validation Failure",
						Detail:    fmt.Sprintf("configuration fields cannot be null, remove '%s' or set a non-null value", s),
						Attribute: &tftypes.AttributePath{},
					})
			}
		}
	}
	return diags
}

// Searches a given set of descriptors for a matching className, when found it will check all fields types for
// a CONCEALED flag or COMPOSITE if CONCEALED, we massage the configuration to to remove the encryptedValue returned by
// current API and set the value back to the original defined. For COMPOSITE fields we then iterate recursively on its
// fields.
//
// TODO This has a drawback that we cannot detect drift in CONCEALED fields due to the way the PingAccess API works.
func maskConfigFromDescriptorsAsMap(desc *models.DescriptorsView, className string, input, config map[string]interface{}) string {
	in, _ := json.Marshal(input)
	orig, _ := json.Marshal(config)
	var newConf string
	for _, value := range desc.Items {
		if *value.ClassName == className {
			newConf = maskConfigFromDescriptor(value, String(""), string(orig), string(in))
		}
	}
	return newConf
}

// Searches a given set of descriptors for a matching className, when found it will check all fields types for
// a CONCEALED flag or COMPOSITE if CONCEALED, we massage the configuration to to remove the encryptedValue returned by
// current API and set the value back to the original defined. For COMPOSITE fields we then iterate recursively on its
// fields.
//
// TODO This has a drawback that we cannot detect drift in CONCEALED fields due to the way the PingAccess API works.
func maskConfigFromDescriptors(desc *models.DescriptorsView, className string, input, config string) string {
	//var conf string
	//if input.Is(tftypes.String) {
	//	input.As(&conf)
	//} else {
	//	var configuration asgotypes.GoPrimitive
	//	input.As(&configuration)
	//	b, _ := json.Marshal(configuration.Value)
	//	conf = string(b)
	//}
	//var originalConfig string
	//if config.Is(tftypes.String) {
	//	config.As(&originalConfig)
	//} else {
	//	var configuration asgotypes.GoPrimitive
	//	config.As(&configuration)
	//	b, _ := json.Marshal(configuration.Value)
	//	originalConfig = string(b)
	//}
	var newConf string
	for _, value := range desc.Items {
		if *value.ClassName == className {
			newConf = maskConfigFromDescriptor(value, String(""), config, input)
		}
	}
	return newConf
	//if input.Is(tftypes.String) {
	//	return tftypes.NewValue(tftypes.String, newConf)
	//}
	//return tftypes.Value{}
}

func maskConfigFromDescriptor(desc *models.DescriptorView, input *string, originalConfig string, config string) string {
	for _, c := range desc.ConfigurationFields {
		config = maskConfigFromConfigurationField(c, input, originalConfig, config)
	}
	return config
}

func maskConfigFromConfigurationField(field *models.ConfigurationField, input *string, originalConfig string, config string) string {
	if *field.Type == "CONCEALED" {
		path := fmt.Sprintf("%s.value", *field.Name)
		v := gjson.Get(originalConfig, path)
		if v.Exists() {
			config, _ = sjson.Set(config, path, v.String())
		} else if v = gjson.Get(originalConfig, *field.Name); v.Exists() {
			config, _ = sjson.Set(config, *field.Name, v.String())
		}
		config, _ = sjson.Delete(config, fmt.Sprintf("%s.encryptedValue", *field.Name))
	} else if *field.Type == "COMPOSITE" {
		for _, value := range field.Fields {
			newInput := String(fmt.Sprintf("%s.%s", *input, *value.Name))
			if *input == "" {
				newInput = value.Name
			}
			config = maskConfigFromConfigurationField(value, newInput, originalConfig, config)
		}
	}
	return config
}
