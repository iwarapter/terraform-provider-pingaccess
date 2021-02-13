package asgotypes

import (
	"errors"
	"math/big"
	"reflect"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

// GoPrimitive is a way to get at the contents of a tftypes.Value without
// asserting anything about the tftypes.Value except that it is fully known. It
// is the equivalent of unmarshalling JSON to an interface{}.
//
// GoPrimitive is not meant to be used during normal Terraform development. As
// tempting as it is, you should not use it during normal provider development
// to obtain easy access to the contents of a tftypes.Value as a standard Go
// type. Terraform relies heavily on the concept of unknown values; values that
// are typed, that will have a value at some point, but that value is not
// currently known. Go is incapable of expressing this concept using the
// builtin type system, so trying to convert an aggregate type to a Go type
// always runs the risk that one of the elements or attributes of the aggregate
// type is unknown, and the Go type will not be able to preserve that
// information.
//
// GoPrimitive is largely a helper for debugging and the very, very rare cases
// when a value is guaranteed to be fully known by the Terraform protocol (for
// example, when working with the PlannedState of ApplyResourceChange when none
// of the attributes are Computed) and the provider wants to pass an opaque
// blob of information to an API and doesn't know or care about the types
// involved. When the types are known ahead of time, other helpers are more
// appropriate.
type GoPrimitive struct {
	Value interface{}
}

// FromTerraform5Value controls how the GoPrimitive will be populated by a
// tftypes.Value.
func (dt *GoPrimitive) FromTerraform5Value(value tftypes.Value) error {
	if !value.IsKnown() {
		return errors.New("cannot decode unknown values to Go types")
	}
	if value.IsNull() {
		dt.Value = nil
		return nil
	}
	switch {
	case value.Is(tftypes.String):
		var str string
		err := value.As(&str)
		if err != nil {
			return err
		}
		dt.Value = str
		return nil
	case value.Is(tftypes.Number):
		num := big.NewFloat(-42)
		err := value.As(&num)
		if err != nil {
			return err
		}
		dt.Value = num
		return nil
	case value.Is(tftypes.Bool):
		var b bool
		err := value.As(&b)
		if err != nil {
			return err
		}
		dt.Value = b
		return nil
	case value.Is(tftypes.Object{}):
		msv := map[string]tftypes.Value{}
		err := value.As(&msv)
		if err != nil {
			return err
		}
		res := map[string]interface{}{}
		for k, v := range msv {
			var vgp GoPrimitive
			err = v.As(&vgp)
			if err != nil {
				return err
			}
			res[k] = vgp.Value
		}
		dt.Value = res
		return nil
	case value.Is(tftypes.Tuple{}):
		vals := []tftypes.Value{}
		err := value.As(&vals)
		if err != nil {
			return err
		}
		res := []interface{}{}
		for _, v := range vals {
			var vgp GoPrimitive
			err = v.As(&vgp)
			if err != nil {
				return err
			}
			res = append(res, vgp.Value)
		}
		dt.Value = res
		return nil
	case value.Is(tftypes.List{}) || value.Is(tftypes.Set{}):
		vals := []tftypes.Value{}
		err := value.As(&vals)
		if err != nil {
			return err
		}
		var tmp []interface{}
		if len(vals) < 1 {
			dt.Value = tmp
			return nil
		}
		for _, v := range vals {
			var vgp GoPrimitive
			err = v.As(&vgp)
			if err != nil {
				return err
			}
			tmp = append(tmp, vgp.Value)
		}
		typ := reflect.TypeOf(tmp[0])
		sliceTyp := reflect.SliceOf(typ)
		res := reflect.MakeSlice(sliceTyp, 0, len(tmp))
		for _, v := range tmp {
			res = reflect.Append(res, reflect.ValueOf(v))
		}
		dt.Value = res.Interface()
		return nil
	case value.Is(tftypes.Map{}):
		msv := map[string]tftypes.Value{}
		err := value.As(&msv)
		if err != nil {
			return err
		}
		tmp := map[string]interface{}{}
		if len(msv) < 1 {
			dt.Value = tmp
			return nil
		}
		var typ reflect.Type
		for k, v := range msv {
			var vgp GoPrimitive
			err = v.As(&vgp)
			if err != nil {
				return err
			}
			if typ == nil {
				typ = reflect.TypeOf(vgp.Value)
			}
			tmp[k] = vgp.Value
		}
		mapTyp := reflect.MapOf(reflect.TypeOf(""), typ)
		res := reflect.MakeMapWithSize(mapTyp, len(tmp))
		for k, v := range tmp {
			res.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		dt.Value = res.Interface()
		return nil
	}
	return errors.New("unknown type")
}
