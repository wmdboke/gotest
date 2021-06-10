package main

import (
	"fmt"
	"github.com/uber/jaeger-client-go"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
)

var (
	tracer opentracing.Tracer
)

func GetListPoc(w http.ResponseWriter, req *http.Request) {
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	span := tracer.StartSpan("GetListProc", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	fmt.Println("Get request getList")
	respList := []string{"l1", "l2", "l3", "l4", "l5"}
	respString := ""

	for _, v := range respList {
		respString += v + ","
	}

	fmt.Println(respString)
	io.WriteString(w, respString)
}

func TraceInit(serviceName, samplerType string, samplerParam float64) (opentracing.Tracer, io.Closer) {
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

func main() {
	var closer io.Closer
	tracer, closer = TraceInit("Trace-Server", "const", 1)
	defer closer.Close()

	http.HandleFunc("/getList", GetListPoc)
	http.ListenAndServe(":8080", nil)
}
