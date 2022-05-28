// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     SysFlow.avsc
 */
package sfgo

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewProcessEventWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewProcessEvent()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type ProcessEventReader struct {
	r io.Reader
	p *vm.Program
}

func NewProcessEventReader(r io.Reader) (*ProcessEventReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewProcessEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &ProcessEventReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r ProcessEventReader) Read() (*ProcessEvent, error) {
	t := NewProcessEvent()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
