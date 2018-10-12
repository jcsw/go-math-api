package pkg

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/jcsw/go-math-api/pkg/controller"
	"github.com/jcsw/go-math-api/pkg/driver/mongodb"
	"github.com/jcsw/go-math-api/pkg/driver/properties"
	"github.com/jcsw/go-math-api/pkg/driver/syslog"
	"github.com/jcsw/go-math-api/pkg/math"
)

type key int

const (
	requestIDKey key = 0
)

var healthy int32

// Server define the Server
type Server struct {
	http      *http.Server
	startDate time.Time
}

// Initialize initialize the all components to server
func (server *Server) Initialize(env string) {
	server.startDate = time.Now()

	syslog.Info("Initialize server by env [%s]", env)

	properties.LoadProperties(env)
	mongodb.InitializeMongoClient()

	router := http.NewServeMux()
	router.HandleFunc("/health", health)
	router.HandleFunc("/monitor", controller.MonitorHandler)

	mathRepository := math.RepositoryMongodb{MongoClient: mongodb.RetrieveMongoClient()}
	mathService := math.NewService(&mathRepository)
	mathHandler := controller.MathHandler{MathService: mathService}

	router.HandleFunc("/math/operation", mathHandler.Register)

	server.http = &http.Server{
		Addr:         fmt.Sprintf(":%d", properties.AppProperties.ServerPort),
		ErrorLog:     syslog.Logger(),
		Handler:      tracing()(logging()(router)),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
}

// Start initializes the server
func (server *Server) Start() {
	syslog.Info("Server is ready to handle requests at port %d, elapsed time to start was %v", properties.AppProperties.ServerPort, time.Since(server.startDate))

	atomic.StoreInt32(&healthy, 1)
	if err := server.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		syslog.Fatal("Could not listen on port [%s]\n%v", properties.AppProperties.ServerPort, err)
	}
}

// Stop stop the server
func (server *Server) Stop() {
	syslog.Info("Server is shutting down...")

	atomic.StoreInt32(&healthy, 0)

	mongodb.CloseMongoClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.http.SetKeepAlivesEnabled(false)
	if err := server.http.Shutdown(ctx); err != nil {
		syslog.Fatal("Could not gracefully shutdown the server\n%v", err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			requestID, ok := r.Context().Value(requestIDKey).(string)
			if !ok {
				requestID = "unknown"
			}
			syslog.Info("requestID=%s, method=%s path=%s remoteAddr=%s elapsedTime=%v",
				requestID, r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		})
	}
}
func tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = newRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func newRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
