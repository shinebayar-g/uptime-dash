package handlers

import (
	"main/gen/uptime_dash/v1/uptime_dash_v1connect"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UptimeServer struct{}

func StartServer() {
	server := &UptimeServer{}
	api := http.NewServeMux()
	path, handler := uptime_dash_v1connect.NewUptimeServiceHandler(server)
	api.Handle(path, handler)
	http.ListenAndServe(net.JoinHostPort("0.0.0.0", os.Getenv("PORT")), h2c.NewHandler(api, &http2.Server{}))
}
