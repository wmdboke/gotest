package main

import (
	//	"context"
	"fmt"
	"github.com/uber/jaeger-client-go"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	//	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
)

const (
	URL        = "http://localhost:8080"
	LIST_API   = "/getList"
	RESULT_API = "/getResult"
)

var (
	flag = make(chan bool)
)

func TraceInit(serviceName string, samplerType string, samplerParam float64) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  samplerType,
			Param: samplerParam,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "127.0.0.1:6831",
			LogSpans:           true,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Init failed: %v\n", err))
	}

	return tracer, closer
}

func sendRequest(req *http.Request) {
	go func(req *http.Request) {
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			fmt.Printf("Do send requst failed(%s)\n", err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ReadAll error(%s)\n", err)
			return
		}

		if resp.StatusCode != 200 {
			return
		}
		fmt.Printf("Response:%s\n", string(body))
		flag <- true
	}(req)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Argument error(getlist or getresult number) ")
		os.Exit(1)
	}

	tracer, closer := TraceInit("CS-tracing", "const", 1)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan(fmt.Sprintf("%s trace", os.Args[1]))
	span.SetTag("trace to", os.Args[1])
	defer span.Finish()
	api := ""
	var err error

	if os.Args[1] == "getlist" {
		api = LIST_API
	}

	reqURL := URL + api
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = span.Tracer().Inject(span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	if os.Args[1] == "getresult" {
		q := req.URL.Query()
		q.Add("num", os.Args[2])
		req.URL.RawQuery = q.Encode()
	}
	sendRequest(req)

	<-flag
}
