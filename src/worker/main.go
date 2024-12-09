package main

import (
	bootstrap "github.com/khaitq-vnist/auto_ci_be/worker/boostrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.All()).Run()
}
