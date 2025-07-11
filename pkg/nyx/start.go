package nyx

import (
	"net/http"
	"time"

	NyxShared "github.com/MatteoMori/nyx/pkg/shared"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
Nyx Entrypoint
*/
func Start(nyxConfig NyxShared.Config) {
	go startMetricsServer(nyxConfig) // Start metrics on port 9090
	startAppServer()                 // Start app on port 8080 (blocking)
}

/*
Nyx web app
*/
func startAppServer() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(MetricsMiddleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Nyx!"})
	})

	r.Run(":8080")
}

/*
Prometheus metrics server
*/
func startMetricsServer(nyxConfig NyxShared.Config) {
	metricsRouter := gin.Default()
	metricsRouter.SetTrustedProxies(nil)

	// Expose Prometheus metrics
	metricsRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

	metricsRouter.Run(":" + nyxConfig.PrometheusPort) // Start metrics server on configured port
}

// Gin middleware to collect metrics
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		httpRequestsTotal.WithLabelValues(path, c.Request.Method, http.StatusText(status)).Inc()
		httpRequestDuration.WithLabelValues(path, c.Request.Method).Observe(duration)
	}
}
