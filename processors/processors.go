package processors

import (
	"github.com/karimra/gnmic/formatters"
	"github.com/mitchellh/mapstructure"
)

var EventProcessors = map[string]Initializer{}

type Initializer func() EventProcessor

func Register(name string, initFn Initializer) {
	EventProcessors[name] = initFn
}

type EventProcessor interface {
	Init(interface{}) error
	Apply(*formatters.EventMsg)
}

func DecodeConfig(src, dst interface{}) error {
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
			Result:     dst,
		},
	)
	if err != nil {
		return err
	}
	return decoder.Decode(src)
}
