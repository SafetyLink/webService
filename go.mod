module github.com/SafetyLink/webService

go 1.20

replace github.com/SafetyLink/commons => ../../commons

require (
	buf.build/gen/go/asavor/safetylink/grpc/go v1.3.0-20230503003809-2accec1d0e5f.1
	buf.build/gen/go/asavor/safetylink/protocolbuffers/go v1.30.0-20230503003809-2accec1d0e5f.1
	github.com/SafetyLink/commons v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.44.0
	github.com/joho/godotenv v1.5.1
	github.com/rabbitmq/amqp091-go v1.8.0
	go.uber.org/fx v1.19.2
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.53.0
)

require (
	buf.build/gen/go/asavor/safetylink/bufbuild/connect-go v1.7.0-20230503003809-2accec1d0e5f.1 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.15.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.47.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)