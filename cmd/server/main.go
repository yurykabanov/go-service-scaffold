package main

//go:generate swag init -o ../../docs

import (
	"github.com/yurykabanov/service-scaffold/internal/configfx"
	"github.com/yurykabanov/service-scaffold/internal/loggerfx"
	"github.com/yurykabanov/service-scaffold/internal/tracingfx"
	"github.com/yurykabanov/service-scaffold/internal/transportfx"
	"github.com/yurykabanov/service-scaffold/internal/validatorfx"
	"go.uber.org/fx"
)

// @title Some service scaffold
// @description Some service description
// @version 0.1.0
// @license.name MIT

func main() {
	app := fx.New(
		configfx.Module,
		loggerfx.Module,
		tracingfx.Module,
		validatorfx.Module,
		transportfx.Module,
	)

	app.Run()
}
