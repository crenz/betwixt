package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/zubairhamed/betwixt/webadmin/app"
	"log"
)

func main() {
	if store, err := app.NewBoltStore("app.db"); err == nil {
		webApp := app.NewWebApp(store, parseConfig())

		webApp.Serve()
	} else {
		log.Fatal(err)
	}
}

func parseConfig() app.ServerConfig {
	cfg := map[string]string{}
	if file, err := yaml.ReadFile("./config.yaml"); err == nil {
		m := file.Root.(yaml.Map)

		cfg["name"] = m.Key("name").(yaml.Scalar).String()
		cfg["http-port"] = m.Key("http").(yaml.Map).Key("port").(yaml.Scalar).String()

		return cfg
	}
	return cfg
}

/*
name: Betwixt LWM2M WebAdmin Server
http:
  port: 8081
coap:
  port: 5683
*/
