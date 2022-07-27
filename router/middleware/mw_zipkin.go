package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const endpointURL = "http://10.13.16.212:9411/api/v2/spans"

// NewZipKin new zipkinhttp client
func NewZipKin(serviceName, hostPort string) (*zipkinhttp.Client, *zipkin.Tracer) {
	// set up a span reporter
	reporter := reporterhttp.NewReporter(endpointURL)
	// defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		log.Fatalf("unable to create sampler: %+v\n", err)
	}

	// initialize our tracer
	tracer, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(endpoint),
	)

	// create global zipkin traced http client
	client, err := zipkinhttp.NewClient(tracer, zipkinhttp.ClientTrace(true))
	if err != nil {
		log.Fatalf("unable to create client: %+v\n", err)
	}
	return client, tracer
}

// Zipkin zipkin
func Zipkin(client *zipkinhttp.Client, tracer *zipkin.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := tracer.StartSpan(ctx.Request.RequestURI)

		span.Annotate(time.Now(), "expensive_calc_done")

		span.Tag("ip", ctx.ClientIP())
		span.Tag("method", ctx.Request.Method)
		span.Tag("url", ctx.Request.URL.String())
		span.Tag("proto", ctx.Request.Proto)
		span.Tag("user_agent", ctx.GetHeader("User-Agent"))
		span.Tag("content_length", fmt.Sprintf("%d", ctx.Request.ContentLength))
		span.Tag("header", fmt.Sprintf("%s", func() []byte {
			b, _ := json.Marshal(ctx.Request.Header)
			return b
		}()))

		ctx.Set("zipkinSpan", span)
		defer span.Finish()

		ctx.Next()
	}
}
