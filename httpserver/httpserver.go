package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// PingReply ...
type PingReply struct {
	URLPath  string `json:"urlpath,omitempty"`
	Method   string `json:"method,omitempty"`
	ClientIP string `json:"clientip,omitempty"`
}

func (c *PingReply) ToString() string {
	pjson, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(pjson)
}

// HttpServer ...
type HttpServer struct {
	s *http.Server
}

// NewServer ...
func NewServer(listen string) *HttpServer {

	s := &HttpServer{}

	go func() {
		s.s = &http.Server{Addr: listen, Handler: s}
		s.s.ListenAndServe()
	}()
	return s
}

// Shutdown ...
func (s *HttpServer) Shutdown() {
	zap.L().Debug("Shutting down Search Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.s.Shutdown(ctx)
}

// ServeHTTP ...
func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	clientIP := getRemoteIP(req)

	zap.L().Debug(fmt.Sprintf("URLPath: %s, Method: %s, ClientIP: %s", req.URL.Path, req.Method, clientIP))

	switch req.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<html>")
		fmt.Fprintf(w, "<h3>Welcome to the test server; to create a unique connection use the command curl http://"+req.Host+"</h3")
		fmt.Fprintf(w, "</html>")
		return
	}

	// By default respond with JSON PingReply
	w.Header().Set("Content-Type", "application/json")

	pingReply := &PingReply{
		URLPath:  req.URL.Path,
		Method:   req.Method,
		ClientIP: clientIP,
	}
	fmt.Fprintf(w, pingReply.ToString()+"\n")
}

// if reqHeadersBytes, err := json.Marshal(req.Header); err != nil {
// 	zap.L().Debug("Could not Marshal Req Headers")
// } else {
// 	zap.L().Debug(string(reqHeadersBytes))
// }

func getRemoteIP(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return req.RemoteAddr
}
