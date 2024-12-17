module github.com/samber/slog-otel

go 1.21
toolchain go1.22.5

require (
	go.opentelemetry.io/otel/trace v1.33.0
	go.uber.org/goleak v1.3.0
)

require go.opentelemetry.io/otel v1.33.0 // indirect
