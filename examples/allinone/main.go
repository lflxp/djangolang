package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/lflxp/djangolang"
	"github.com/lflxp/djangolang/middlewares"
	ctls "github.com/lflxp/djangolang/tls"
	"github.com/lflxp/djangolang/utils"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lflxp/djangolang/examples/allinone/demo"
	_ "github.com/lflxp/djangolang/examples/allinone/docs"
	"github.com/skratchdot/open-golang/open"
)

var GinEngine *gin.Engine

type Args struct {
	Host        string
	Port        string
	IsHttps     bool
	OpenBrowser bool
	Auth        struct {
		Url         string
		IdentityKey string
		Dev         bool
	}
}

func Run(args *Args) {
	// gin.SetMode(gin.ReleaseMode)
	GinEngine = gin.Default()

	// 注册路由

	GinEngine.Use(gin.Logger())
	GinEngine.Use(gin.Recovery())
	GinEngine.Use(middlewares.Cors())
	GinEngine.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
	// GinEngine.Use(middlewares.NoRouteHandler)
	GinEngine.GET("/health", middlewares.RegisterHealthMiddleware)
	GinEngine.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/admin/index")
	})
	// 设置swagger
	middlewares.RegisterSwaggerMiddleware(GinEngine)
	middlewares.RegisterNodeExporter(GinEngine)
	Registertest(GinEngine)
	demo.RegisterDemo(GinEngine)
	demo.RegisterVpn(GinEngine)
	djangolang.RegisterControllerAdmin(GinEngine, true)
	slog.Info("ip %s port %s", args.Host, args.Port)

	if args.Host == "" {
		// instance.Fatal("Check Host or Port config already!!!")
		args.Host = "0.0.0.0"
	}

	if args.Port == "" {
		args.Port = "8002"
	}

	var server *http.Server
	if args.IsHttps {
		err := ctls.Refresh()
		if err != nil {
			panic(err)
		}

		pool := x509.NewCertPool()
		caCeretPath := "ca.crt"

		caCrt, err := os.ReadFile(caCeretPath)
		if err != nil {
			panic(err)
		}

		pool.AppendCertsFromPEM(caCrt)

		server = &http.Server{
			Addr:    fmt.Sprintf("%s:%s", args.Host, args.Port),
			Handler: GinEngine,
			TLSConfig: &tls.Config{
				ClientCAs:  pool,
				ClientAuth: tls.RequestClientCert,
			},
		}
	} else {
		server = &http.Server{
			Addr:    fmt.Sprintf("%s:%s", args.Host, args.Port),
			Handler: GinEngine,
		}

	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		slog.Warn("receive interrupt signal")
		if err := server.Close(); err != nil {
			slog.Error("Server Close:", err)
		}
	}()

	var openUrl string
	for index, ip := range utils.GetIPs() {
		if args.IsHttps {
			slog.Info("Listening and serving HTTPS on https://%s:%s", ip, args.Port)
		} else {
			slog.Info("Listening and serving HTTPS on http://%s:%s", ip, args.Port)
		}

		if index == 0 {
			openUrl = fmt.Sprintf("%s:%s", ip, args.Port)
		}
	}
	if args.IsHttps {
		if args.OpenBrowser {
			open.Start(fmt.Sprintf("https://%s", openUrl))
		}
		if err := server.ListenAndServeTLS("ca.crt", "ca.key"); err != nil {
			if err == http.ErrServerClosed {
				slog.Warn("Server closed under request")
			} else {
				slog.Error("Server closed unexpect %s", err.Error())
			}
		}
	} else {
		if args.OpenBrowser {
			open.Start(fmt.Sprintf("http://%s", openUrl))
		}
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				slog.Warn("Server closed under request")
			} else {
				slog.Error("Server closed unexpect %s", err.Error())
			}
		}
	}

	slog.Warn("Server exiting")
}

func main() {
	args := Args{
		IsHttps:     true,
		OpenBrowser: true,
		Port:        "9000",
	}
	Run(&args)
}
