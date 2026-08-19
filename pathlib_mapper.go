package kongutil

import (
	"errors"
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/chigopher/pathlib"
)

func PathlibMapper() kong.Option {
	return kong.TypeMapper(reflect.TypeFor[*pathlib.Path](), kong.MapperFunc(func(ctx *kong.DecodeContext, target reflect.Value) error {
		var value string
		if err := ctx.Scan.PopValueInto("path", &value); err != nil {
			return err
		}
		p := pathlib.NewPath(value)
		if p == nil {
			return errors.New("cannot create pathlib value")
		}
		target.Set(reflect.ValueOf(p))
		return nil
	}))
}
