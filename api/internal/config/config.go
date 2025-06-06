package config

import (
	"reflect"

	"github.com/caarlos0/env/v11"
)

type SecretType []byte

type Config struct {
	JwtSecret SecretType `env:"JWT_SECRET,required"`
}

func LoadConfig() (*Config, error) {
	c, err := env.ParseAsWithOptions[Config](
		env.Options{
			FuncMap: map[reflect.Type]env.ParserFunc{
				reflect.TypeOf(SecretType{}): func(v string) (interface{}, error) {
					return []byte(v), nil
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
