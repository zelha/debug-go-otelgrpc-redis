package main

import (
	redisotel "github.com/go-redis/redis/extra/redisotel/v9"
	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
)


func main(){
	otlpgrpc.NewDriver(
		// Configure metrics driver here
	)

	redisotel.NewTracingHook()
}
