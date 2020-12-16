package plugins

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// SFChannel defines a SysFlow channel for data transfer.
type SFChannel struct {
	In   chan *sfgo.SysFlow
	once sync.Once
}

// SFProcessor defines the SysFlow processor interface.
type SFProcessor interface {
	Register(pc SFPluginCache)
	Init(conf map[string]string) error
	Process(record interface{}, wg *sync.WaitGroup)
	GetName() string
	SetOutChan(ch interface{})
	Cleanup()
}
