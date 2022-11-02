package httpServer

import (
	"net/http"
	"net/http/pprof"
	"time"

	sm "github.com/krassor/serverHttp/pkg/supportModule"
)

func ServerHttpStart(httpPort string) error {
	sm.PrintlnWithTimeShtamp("Server HTTP starting...")
	PORT := ":" + httpPort
	r := http.NewServeMux()

	srv := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	r.HandleFunc("api/v1/news", newsHandler)
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/api/coins", coinsHandler)
	r.HandleFunc("/change", changeElement)
	r.HandleFunc("/list", listAll)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	sm.PrintlnWithTimeShtamp("Server HTTP listening...")
	err := srv.ListenAndServe()
	defer srv.Close()
	if err != nil {
		return err
	}
	return nil
}
