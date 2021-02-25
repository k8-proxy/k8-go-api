package middleware

import (
	"k8-go-api/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// AuthMiddleware to check authorization
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errauth := "you don't have  valid authoriaztion token"
		erremptyauth := "you didn't provide authoriaztion token"

		//log about request
		//there will be logging middleware soon
		/*
			log.Printf("method: %v\n", r.Method)
			log.Printf("URL: %v\n", r.URL)
			log.Printf("RemoteAddr: %v\n", r.RemoteAddr)
			log.Printf("Host: %v\n", r.Host)
			log.Printf("Content-Type: %v\n", r.Header.Get("Content-Type"))
			log.Printf("RequestURI: %v\n", r.RequestURI)
		*/
		//Authorization: Bearer
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResponseWithError(w, http.StatusUnauthorized, erremptyauth)
			return
		}

		authHeaderParts := strings.Fields(authHeader)
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			utils.ResponseWithError(w, http.StatusUnauthorized, errauth)
			return
		}

		if authHeaderParts[1] != "mysecrettoken" {
			utils.ResponseWithError(w, http.StatusUnauthorized, errauth)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Logmiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := zerolog.New(os.Stdout).With().
			Timestamp().
			Str("role", "my-service").
			Str("host", "host").
			Logger()

		c := alice.New()

		// Install the logger handler with default output on the console
		c = c.Append(hlog.NewHandler(log))

		c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}))
		c = c.Append(hlog.RemoteAddrHandler("ip"))
		c = c.Append(hlog.UserAgentHandler("user_agent"))
		c = c.Append(hlog.RefererHandler("referer"))
		c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

		h := c.Then(next)
		h.ServeHTTP(w, r)

	})

}
