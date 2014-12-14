package main

import (
	"testing"
)

const yaml_params = `
---

port: 9999

host: 0.0.0.0  #localhost

# all the params relative to mongodb
influxdb:
    host: "bigtits.com"
    username: "alice"
    password: "qwerty"
    database: "tits"
    issecure: yes
`

func TestLoadFromYamlData(t *testing.T) {
	p := GetDefaultParams()
	if err := p.LoadFromYamlData([]byte(yaml_params)); err != nil {
		t.Error(err)
		return
	}
	if p.Port != 9999 {
		t.Error("port")
	}
	if p.Host != "0.0.0.0" {
		t.Error("host")
	}
	if p.InfluxDBClientConfig.Host != "bigtits.com" {
		t.Error("InfluxDBClientConfig.Host", p.InfluxDBClientConfig.Host, " fuck")
	}
	if p.InfluxDBClientConfig.Database != "tits" {
		t.Error("InfluxDBClientConfig.Database")
	}
	if !p.InfluxDBClientConfig.IsSecure {
		t.Error("InfluxDBClientConfig.IsSecure")
	}

}
