// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCES:
 *     GraphletRecord.avsc
 *     SysFlow.avsc
 */
package sfgo

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type PoidUnionTypeEnum int

const (
	PoidUnionTypeEnumOID PoidUnionTypeEnum = 1
)

type PoidUnion struct {
	Null      *types.NullVal
	OID       *OID
	UnionType PoidUnionTypeEnum
}

func writePoidUnion(r *PoidUnion, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case PoidUnionTypeEnumOID:
		return writeOID(r.OID, w)
	}
	return fmt.Errorf("invalid value for *PoidUnion")
}

func NewPoidUnion() *PoidUnion {
	return &PoidUnion{}
}

func (_ *PoidUnion) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *PoidUnion) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *PoidUnion) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *PoidUnion) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *PoidUnion) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *PoidUnion) SetString(v string)  { panic("Unsupported operation") }
func (r *PoidUnion) SetLong(v int64) {
	r.UnionType = (PoidUnionTypeEnum)(v)
}
func (r *PoidUnion) Get(i int) types.Field {
	switch i {
	case 0:
		return r.Null
	case 1:
		r.OID = NewOID()
		return r.OID
	}
	panic("Unknown field index")
}
func (_ *PoidUnion) NullField(i int)                  { panic("Unsupported operation") }
func (_ *PoidUnion) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *PoidUnion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PoidUnion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *PoidUnion) Finalize()                        {}

func (r *PoidUnion) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case PoidUnionTypeEnumOID:
		return json.Marshal(map[string]interface{}{"OID": r.OID})
	}
	return nil, fmt.Errorf("invalid value for *PoidUnion")
}

func (r *PoidUnion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["OID"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.OID)
	}
	return fmt.Errorf("invalid value for *PoidUnion")
}
