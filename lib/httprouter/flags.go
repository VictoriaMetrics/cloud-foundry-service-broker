package httprouter

import "flag"

var (
	bindPort = flag.Int("bind-port", 8080, "port to bind")
	bindHost = flag.String("bind-host", "0.0.0.0", "host to bind")
)
