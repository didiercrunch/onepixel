package main

import (
	"io/ioutil"
	"github.com/influxdb/influxdb/client"

	goyaml "gopkg.in/yaml.v2"
)

const PARAM_FILE string = "params.yml"

type Params struct {
	Port                 int                  `yaml:"port,omitempty"`
	Host                 string               `yaml:"host,omitempty"`
	InfluxDBClientConfig *client.ClientConfig `yaml:"influxdb,omitempty"`
	InfluxDBClient       *client.Client       `yaml:"-"`
}

func (this *Params) LoadFromYamlFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return this.LoadFromYamlData(data)
}

func (this *Params) LoadFromYamlData(data []byte) error {
	var err error
	p := new(Params)
	if err = goyaml.Unmarshal(data, p); err != nil {
		return err
	}

	if p.Host != "" {
		this.Host = p.Host
	}
	if p.Port > 0 {
		this.Port = p.Port
	}
	if p.InfluxDBClientConfig == nil {
		return nil
	}
	if p.InfluxDBClientConfig.Host != "" {
		this.InfluxDBClientConfig.Host = p.InfluxDBClientConfig.Host
	}
	if p.InfluxDBClientConfig.Database != "" {
		this.InfluxDBClientConfig.Database = p.InfluxDBClientConfig.Database
	}
	if p.InfluxDBClientConfig.Password != "" {
		this.InfluxDBClientConfig.Password = p.InfluxDBClientConfig.Password
	}
	if p.InfluxDBClientConfig.Username != "" {
		this.InfluxDBClientConfig.Username = p.InfluxDBClientConfig.Username
	}
	this.InfluxDBClientConfig.IsSecure = p.InfluxDBClientConfig.IsSecure

	return nil
}

func (this *Params) initInfluxDB() error {
	var err error
	this.InfluxDBClient, err = client.NewClient(this.InfluxDBClientConfig)

	return err
}

func GetDefaultParams() *Params {
	influxdbClientConfig := &client.ClientConfig{}
	return &Params{7777, "localhost", influxdbClientConfig, nil}
}

var params = GetDefaultParams()

func init() {
	if err := params.LoadFromYamlFile(PARAM_FILE); err != nil {
		panic(err)
	}
	if err := params.initInfluxDB(); err != nil {
		panic(err)
	}
}
