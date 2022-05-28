// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     SysFlow.avsc
 */
package sfgo

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

func writeArrayPort(r []*Port, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writePort(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayPortWrapper struct {
	Target *[]*Port
}

func (_ *ArrayPortWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayPortWrapper) Finalize()                        {}
func (_ *ArrayPortWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r *ArrayPortWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r *ArrayPortWrapper) AppendArray() types.Field {
	var v *Port
	v = NewPort()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}