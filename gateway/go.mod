module github.com/haserta98/go-rmq-learning/gateway

go 1.24.0

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/fxamacker/cbor/v2 v2.8.0 // indirect
	github.com/gofiber/fiber/v3 v3.0.0-beta.4 // indirect
	github.com/gofiber/schema v1.4.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/tinylib/msgp v1.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.62.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	gorm.io/gorm v1.30.0 // indirect
)

require (
	github.com/haserta98/go-rmq-learning/shared v0.0.0
	github.com/haserta98/go-rmq-learning/order v0.0.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

replace github.com/haserta98/go-rmq-learning/shared => ../shared
replace github.com/haserta98/go-rmq-learning/order => ../order
