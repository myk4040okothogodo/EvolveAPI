module github.com/myk4040okothogodo/EvolveAPI/pkg/middleware

go 1.18

replace github.com/myk4040okothogodo/EvolveAPI/pkg/types => ../types

replace github.com/myk4040okothogodo/EvolveAPI/pkg/db => ../db

require (
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/render v1.0.1
	github.com/myk4040okothogodo/EvolveAPI/pkg/db v0.0.0-00010101000000-000000000000
	github.com/myk4040okothogodo/EvolveAPI/pkg/types v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.21.0
)

require (
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
