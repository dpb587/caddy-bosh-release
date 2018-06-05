package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	_ "github.com/caddyserver/dnsproviders/azure"
	_ "github.com/caddyserver/dnsproviders/googlecloud"
	_ "github.com/caddyserver/dnsproviders/route53"
)

var run = caddymain.Run

func main() {
	run()
}
