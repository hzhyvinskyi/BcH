package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const (
	port         = ":8082"
	serviceBHost = "serviceb:8084"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST_RECEIVED")
	ctx := r.Context()

	span, _ := opentracing.StartSpanFromContext(ctx, "handler")
	defer span.Finish()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://"+serviceBHost+"/trace", nil)
	if err != nil {
		log.Fatalf("FailedToCreateNewRequestWithContext: %s\n", err)
	}
	fmt.Println("REQUEST_SENT")

	client := http.Client{
		Timeout: time.Duration(time.Second * 30),
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("FailedToPerformHTTPRequest: %s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("FailedToReadResponseData: %s\n", err)
	}

	fmt.Println(string(body))
	_, err = w.Write([]byte(string(body)))
	if err != nil {
		log.Fatalf("FailedToWriteTheDataToTheResponse: %s\n", err)
	}
}

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	return tracer, closer, err
}

func main() {
	tracer, closer, err := initJaeger("servicea")
	if err != nil {
		return
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/trace", handler)

	log.Println("HTTP Server is running on port "+port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("FATAL")
	}
}
