package handlers

import (
	"main/gen/uptime_dash/v1/uptime_dash_v1connect"
	"net"
	"net/http"
	"os"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UptimeServer struct{}

func StartServer() {
	server := &UptimeServer{}
	mux := http.NewServeMux()
	path, handler := uptime_dash_v1connect.NewUptimeServiceHandler(server)
	mux.Handle(path, handler)
	base := cors.Default().Handler(mux)
	http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), h2c.NewHandler(base, &http2.Server{}))
}
