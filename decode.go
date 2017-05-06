package runscope

import (
	"reflect"
	"time"
	"github.com/mitchellh/mapstructure"
)

func floatToTimeDurationHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Float64 {
			return data, nil
		}

		if t != reflect.TypeOf(time.Now()) {
			return data, nil
		}

		// Convert it by parsing
		value := int64(data.(float64))
		return time.Unix(value, 0), nil
	}
}

func Decode(result interface{}, response interface{}) error {

	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   result,
		TagName:  "json",
		DecodeHook: floatToTimeDurationHookFunc(),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(response)
	return err
}
