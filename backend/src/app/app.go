package app

import (
	"context"
)

type options struct {
	ConfigFile string
	ModelFile  string
	WWWDir     string
	SwaggerDir string
	Version    string
}

type Option func(*options)

func Setup(ctx context.Context, opts ...Option)  {
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

