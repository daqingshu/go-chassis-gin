package gin

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	archaius "github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2/core/server"
	"github.com/go-chassis/go-chassis/v2/pkg/metrics"
	"github.com/go-chassis/openlog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	//Name is a variable of type string which indicates the protocol being used
	Name = "gin"
	//DefaultMetricPath DefaultMetricPath
	DefaultMetricPath = "metrics"
	//DefaultProfilePath DefaultProfilePath
	DefaultProfilePath = "profile"
	//ProfileRouteRuleSubPath ProfileRouteRuleSubPath
	ProfileRouteRuleSubPath = "route-rule"
	//ProfileDiscoverySubPath ProfileDiscoverySubPath
	ProfileDiscoverySubPath = "discovery"
	//MimeFile MimeFile
	MimeFile = "application/octet-stream"
	//MimeMult MimeMult
	MimeMult = "multipart/form-data"
)

type ginServer struct {
	engine *gin.Engine
	opts   server.Options
	mux    sync.RWMutex
	server *http.Server
}

func newGinServer(opts server.Options) server.ProtocolServer {
	engine := gin.New()
	if archaius.GetBool("servicecomb.metrics.enable", false) {
		metricGroup := engine.Group("")
		metricPath := archaius.GetString("servicecomb.metrics.apiPath", DefaultMetricPath)
		if !strings.HasPrefix(metricPath, "/") {
			metricPath = "/" + metricPath
		}
		openlog.Info("Enabled metrics API on " + metricPath)
		metricGroup.GET(metricPath, prometheusHandleFunc)
	}

	return &ginServer{
		opts:   opts,
		engine: engine,
	}
}

func prometheusHandleFunc(c *gin.Context) {
	promhttp.HandlerFor(metrics.GetSystemPrometheusRegistry(), promhttp.HandlerOpts{}).ServeHTTP(c.Writer, c.Request)
}

func (r *ginServer) Register(schema interface{}, options ...server.RegisterOption) (string, error) {
	return "", nil
}

func (r *ginServer) Start() error {
	return nil
}

func (r *ginServer) Stop() error {
	return nil
}

func (r *ginServer) String() string {
	return Name
}
