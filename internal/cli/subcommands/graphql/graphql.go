package graphql

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/hortbot/hortbot/internal/cli"
	"github.com/hortbot/hortbot/internal/cli/flags/jaegerflags"
	"github.com/hortbot/hortbot/internal/cli/flags/promflags"
	"github.com/hortbot/hortbot/internal/cli/flags/sqlflags"
	"github.com/hortbot/hortbot/internal/db/graph"
	"github.com/hortbot/hortbot/internal/db/graph/generated"
	"github.com/hortbot/hortbot/internal/pkg/ctxlog"
	"github.com/hortbot/hortbot/internal/web/mid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type cmd struct {
	cli.Common
	Jaeger     jaegerflags.Jaeger
	Prometheus promflags.Prometheus
	SQL        sqlflags.SQL

	Addr string `long:"web-addr" env:"HB_GRAPHQL_ADDR" description:"Server address for the web server"`
}

// Command returns a fresh web command.
func Command() cli.Command {
	return &cmd{
		Common:     cli.Default,
		SQL:        sqlflags.Default,
		Jaeger:     jaegerflags.Default,
		Prometheus: promflags.Default,
		Addr:       ":5001",
	}
}

func (*cmd) Name() string {
	return "graphql"
}

func (c *cmd) Main(ctx context.Context, _ []string) {
	defer c.Jaeger.Trace(ctx, c.Name(), c.Debug)()
	c.Prometheus.Run(ctx)

	driverName := c.SQL.DriverName()
	driverName = c.Jaeger.DriverName(ctx, driverName, c.Debug)
	db := c.SQL.Open(ctx, driverName)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	r := chi.NewRouter()

	logger := ctxlog.FromContext(ctx)
	r.Use(mid.Logger(logger))
	r.Use(mid.RequestID)

	r.Use(func(next http.Handler) http.Handler {
		return promhttp.InstrumentHandlerCounter(metricRequest, next)
	})

	r.Use(mid.RequestLogger)
	r.Use(mid.Tracer)
	r.Use(mid.Recoverer)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(c.Addr, r))
}

var metricRequest = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "hortbot",
	Subsystem: "graphql",
	Name:      "request_total",
	Help:      "Total number of HTTP requests.",
}, []string{"code", "method"})
