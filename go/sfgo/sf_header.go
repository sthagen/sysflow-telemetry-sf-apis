// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCES:
 *     GraphletRecord.avsc
 *     SysFlow.avsc
 */
package sfgo

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type SFHeader struct {
	Version int64 `json:"version"`

	Exporter string `json:"exporter"`

	Ip string `json:"ip"`

	Filename string `json:"filename"`
}

const SFHeaderAvroCRC64Fingerprint = "j\x19\x83'\xcc\x15\xa9&"

func NewSFHeader() *SFHeader {
	return &SFHeader{}
}

func DeserializeSFHeader(r io.Reader) (*SFHeader, error) {
	t := NewSFHeader()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeSFHeaderFromSchema(r io.Reader, schema string) (*SFHeader, error) {
	t := NewSFHeader()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeSFHeader(r *SFHeader, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Version, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Exporter, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ip, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Filename, w)
	if err != nil {
		return err
	}
	return err
}

func (r *SFHeader) Serialize(w io.Writer) error {
	return writeSFHeader(r, w)
}

func (r *SFHeader) Schema() string {
	return "{\"fields\":[{\"default\":3,\"name\":\"version\",\"type\":\"long\"},{\"name\":\"exporter\",\"type\":\"string\"},{\"default\":\"NA\",\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"filename\",\"type\":\"string\"}],\"name\":\"sysflow.entity.SFHeader\",\"type\":\"record\"}"
}

func (r *SFHeader) SchemaName() string {
	return "sysflow.entity.SFHeader"
}

func (_ *SFHeader) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *SFHeader) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *SFHeader) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *SFHeader) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *SFHeader) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *SFHeader) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *SFHeader) SetString(v string)   { panic("Unsupported operation") }
func (_ *SFHeader) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SFHeader) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Long{Target: &r.Version}
	case 1:
		return &types.String{Target: &r.Exporter}
	case 2:
		return &types.String{Target: &r.Ip}
	case 3:
		return &types.String{Target: &r.Filename}
	}
	panic("Unknown field index")
}

func (r *SFHeader) SetDefault(i int) {
	switch i {
	case 0:
		r.Version = 3
		return
	case 2:
		r.Ip = "NA"
		return
	}
	panic("Unknown field index")
}

func (r *SFHeader) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *SFHeader) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *SFHeader) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *SFHeader) Finalize()                        {}

func (_ *SFHeader) AvroCRC64Fingerprint() []byte {
	return []byte(SFHeaderAvroCRC64Fingerprint)
}
