module github.com/khaitq-vnist/auto_ci_be/adapter

go 1.21.0

replace github.com/khaitq-vnist/auto_ci_be/core => ../core

require (
	github.com/khaitq-vnist/auto_ci_be/core v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.14.0 // indirect
)
