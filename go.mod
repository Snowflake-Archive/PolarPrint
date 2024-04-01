module github.com/snowflake-software/polarprint

go 1.22.1

require (
	github.com/gofiber/fiber/v2 v2.52.4
	github.com/gofiber/template/handlebars/v2 v2.1.9
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/snowflake-software/polarprint/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mailgun/raymond/v2 v2.0.48 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
)

replace github.com/snowflake-software/polarprint/utils => ./pkg/utils
