package main

import (
	"github.com/khaitq-vnist/auto_ci_be/public/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.All()).Run()
}
