module github.com/myk4040okothogodo/EvolveAPI

go 1.18

replace github.com/myk4040okothogodo/EvolveAPI/pkg/api => ./pkg/api

replace github.com/myk4040okothogodo/EvolveAPI/pkg/db => ./pkg/db

replace github.com/myk4040okothogodo/EvolveAPI/docs => ./docs

require (
	github.com/joho/godotenv v1.4.0
	github.com/myk4040okothogodo/EvolveAPI/pkg/api v0.0.0-20220625095139-884122353443
	github.com/myk4040okothogodo/EvolveAPI/pkg/db v0.0.0-20220625101743-b9f86f911f24
	github.com/spf13/pflag v1.0.5
	github.com/swaggo/swag v1.8.3
	go.uber.org/zap v1.21.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
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
	github.com/myk4040okothogodo/EvolveAPI/pkg/middleware v0.0.0-20220625101743-b9f86f911f24 // indirect
	github.com/myk4040okothogodo/EvolveAPI/pkg/types v0.0.0-20220625101743-b9f86f911f24 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/swaggo/http-swagger v1.3.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
