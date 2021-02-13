package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-go-contrib/asgotypes"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingaccess-sdk-go/pingaccess/models"
	"github.com/tidwall/gjson"
)

func dynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	if !val.Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The data source got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	return values, nil
}

func resourceDynamicValueToTftypesValues(conf *tfprotov5.DynamicValue, types tftypes.Type) (map[string]tftypes.Value, []*tfprotov5.Diagnostic) {
	val, err := conf.Unmarshal(types)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	if !val.Is(types) {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.",
			},
		}
	}
	values := map[string]tftypes.Value{}
	err = val.As(&values)
	if err != nil {
		return nil, []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		}
	}
	return values, nil
}

func planResourceChangeValidationError(err error) *tfprotov5.PlanResourceChangeResponse {
	return &tfprotov5.PlanResourceChangeResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  err.Error(),
			},
		},
	}
}

func readResourceChangeError(err error) *tfprotov5.ReadResourceResponse {
	return &tfprotov5.ReadResourceResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   err.Error(),
			},
		},
	}
}

func planResourceChangeError(err error) *tfprotov5.PlanResourceChangeResponse {
	return &tfprotov5.PlanResourceChangeResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		},
	}
}

func applyResourceChangeError(err error) *tfprotov5.ApplyResourceChangeResponse {
	return &tfprotov5.ApplyResourceChangeResponse{
		Diagnostics: []*tfprotov5.Diagnostic{
			{
				Severity: tfprotov5.DiagnosticSeverityError,
				Summary:  "Unexpected configuration format",
				Detail:   "The resource got a configuration that did not match its schema, This may indication an error in the provider.\n\nError: " + err.Error(),
			},
		},
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
func descriptorsHasClassName(className string, desc *models.DescriptorsView) error {
	var classes []string
	for _, value := range desc.Items {
		classes = append(classes, *value.ClassName)
		if *value.ClassName == className {
			return nil
		}
	}
	return fmt.Errorf("unable to find className '%s' available classNames: %s", className, strings.Join(classes, ", "))
}

//Checks the class name specified exists in the DescriptorsView
//
func validateConfiguration(className string, configuration asgotypes.GoPrimitive, desc *models.DescriptorsView) error {
	var diags diag.Diagnostics
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
					if v := gjson.Get(conf, *f.Name); !v.Exists() {
						diags = append(diags, diag.Errorf("the field '%s' is required for the class_name '%s'", *f.Name, className)...)
					}
				}
			}
		}
	}
	if diags.HasError() {
		msgs := []string{
			"configuration validation failed against the class descriptor definition",
		}
		for _, diagnostic := range diags {
			msgs = append(msgs, diagnostic.Summary)
		}
		return fmt.Errorf(strings.Join(msgs, "\n"))
	}
	return nil
}
