package main

import (
	"contrib.go.opencensus.io/exporter/jaeger"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	conf "github.com/gusarow4321/TinyChat/gateway/internal/config"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
	"os"
)

var (
	// Name is the name of the compiled software.
	Name string = "Gateway"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// prefixs is the config environment variable prefix
	prefixs = []string{"TINY_CHAT_GATEWAY_", "GATEWAY_"}

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func customHandler(mux *runtime.ServeMux, logger log.Logger) (http.Handler, error) {
	if err := mux.HandlePath(
		"GET",
		"/health",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if _, err := w.Write([]byte("Gateway v" + Version + " serving")); err != nil {
				log.NewHelper(logger).Errorf("Health check error: %v", err)
			}
		},
	); err != nil {
		return nil, err
	}

	return wsproxy.WebsocketProxy(mux), nil
}

func newTracing(conf *conf.Tracing, h http.Handler, logger log.Logger) (*ochttp.Handler, func(), error) {
	exp, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: conf.Url,
		Process:           jaeger.Process{ServiceName: Name},
	})
	if err != nil {
		return nil, nil, err
	}

	trace.RegisterExporter(exp)
	// In production can be set to a trace.ProbabilitySampler.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	cleanup := func() {
		log.NewHelper(logger).Info("flushing jaeger exporter")
		exp.Flush()
	}

	return &ochttp.Handler{
		Handler: h,
	}, cleanup, nil
}

func newGatewayServer(conf *conf.Rest, oc *ochttp.Handler) *http.Server {
	return &http.Server{
		Addr:    conf.Addr,
		Handler: oc,
	}
}

func main() {
	flag.Parse()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)

	// config
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
			env.NewSource(prefixs...),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Rest, bc.Auth, bc.Messenger, bc.Tracing, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
