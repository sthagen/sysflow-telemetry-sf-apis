package plugins

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// SFChannel defines a SysFlow channel for data transfer.
type SFChannel struct {
	In chan *sfgo.SysFlow
}

// SFProcessor defines the SysFlow processor interface.
type SFProcessor interface {
	Process(record interface{}, wg *sync.WaitGroup)
	Init(conf map[string]string) error
	SetOutChan(ch interface{})
	Cleanup()
}
