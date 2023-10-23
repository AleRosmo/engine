module github.com/AleRosmo/engine

go 1.21.3

require (
	github.com/AleRosmo/shared_errors v0.0.0-20231011224936-2d211750e59f
	github.com/YalkChat/database v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.14.0
	google.golang.org/protobuf v1.31.0
	gorm.io/gorm v1.25.5
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gorm.io/driver/postgres v1.5.3 // indirect
)

replace github.com/YalkChat/database => ../database
