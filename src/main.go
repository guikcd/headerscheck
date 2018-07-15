package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

type SiteConfiguration struct {
	Url     string
	Header  string
	Headers map[string]string
}

type Configuration struct {
	Site SiteConfiguration
}

func readConfig(configFile string) Configuration {

	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to unmarshal config %s", err)
	}

	return configuration
}

func fetchUrl(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching url: ", err)
	}
	return resp
}

func main() {

	debug := flag.Bool("debug", false, "Enable debugging output")
	configFile := flag.String("config-file", "config", "Config file")
	flag.Parse()

	configuration := readConfig(*configFile)
	if *debug {
		log.Println("Debug enabled")
		log.Println("Config read from", *configFile, "file:", configuration.Site.Headers)
	}

	if *debug {
		log.Println("Fetching url", configuration.Site.Url)
	}
	resp := fetchUrl(configuration.Site.Url)

	for config_key, config_value := range configuration.Site.Headers {
		var found bool = false
		for resp_key, resp_value := range resp.Header {
			if strings.EqualFold(config_key, resp_key) {
				if strings.EqualFold(resp_value[0], config_value) {
					if *debug {
						log.Println("Header", resp_key, "with value", resp_value[0], "in response match config")
					}
					found = true
					break
				}
			}
		}
		if !found {
			log.Fatal(configuration.Site.Url, ": header '", config_key, "' with value '", config_value, "' was not found in the response")
		}
	}

	log.Println("All headers with values found")
}
