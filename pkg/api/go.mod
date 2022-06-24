module github.com/myk4040okothogodo/EvolveAPI/pkg/api

go 1.18

replace github.com/myk4040okothogodo/EvolveAPI/pkg/db => ../db

replace github.com/myk4040okothogodo/EvolveAPI/pkg/middleware => ../middleware

replace github.com/myk4040okothogodo/EvolveAPI/pkg/types => ../types

require (
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/render v1.0.1
	github.com/myk4040okothogodo/EvolveAPI/pkg/db v0.0.0-20220624195938-26d8e7e0cfb1
	github.com/myk4040okothogodo/EvolveAPI/pkg/middleware v0.0.0-20220624195938-26d8e7e0cfb1
	github.com/myk4040okothogodo/EvolveAPI/pkg/types v0.0.0-20220624195938-26d8e7e0cfb1
	github.com/swaggo/http-swagger v1.3.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
