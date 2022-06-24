module github.com/myk4040okothogodo/EvolveAPI

go 1.18

replace github.com/myk4040okothogodo/EvolveAPI/pkg/db => ./pkg/db

replace github.com/myk4040okothogodo/EvolveAPI/pkg/api => ./pkg/api

replace github.com/myk4040okothogodo/EvolveAPI/docs => ./docs

require (
	github.com/golang/mock v1.6.0
	github.com/joho/godotenv v1.4.0
	github.com/myk4040okothogodo/EvolveAPI/pkg/api v0.0.0-20220624195938-26d8e7e0cfb1
	github.com/myk4040okothogodo/EvolveAPI/pkg/db v0.0.0-20220624195938-26d8e7e0cfb1
	github.com/spf13/pflag v1.0.5
	github.com/swaggo/swag v1.8.3
	go.uber.org/zap v1.21.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/myk4040okothogodo/EvolveAPI/pkg/middleware v0.0.0-20220624195938-26d8e7e0cfb1 // indirect
	github.com/myk4040okothogodo/EvolveAPI/pkg/types v0.0.0-20220624195938-26d8e7e0cfb1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/swaggo/http-swagger v1.3.0 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
