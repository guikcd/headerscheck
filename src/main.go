package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
	"strconv"
)

type Configuration struct {
	Urls map[string]Scenario
}

type Scenario struct {
	Url       string
	Code      string
	Headers   map[string]string
	Noheaders map[string]string
	Body      string
	Nobody    string
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
		log.Println("[*] Debug enabled")
		log.Println("[*] Config read from", *configFile, "file:", configuration.Urls)
	}

	for _, config := range configuration.Urls {

		if *debug {
			log.Println("[*] Fetching url", config.Url)
		}
		resp := fetchUrl(config.Url)

		// code
		if config.Code != strconv.Itoa(resp.StatusCode) {
			log.Fatal("Expected status code ", config.Code, " found ", resp.StatusCode)
		} else {
			if *debug {
				log.Println("Code", resp.StatusCode, "match config", config.Code)
			}
		}

		// headers
		for config_key, config_value := range config.Headers {
			if *debug {
				log.Println("read config header:", config_key, config_value)
			}
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
				log.Fatal(config.Url, ": header '", config_key, "' with value '", config_value, "' was not found in the response", resp)
			}
		}

		// noheaders
		for config_key, config_value := range config.Noheaders {
			if *debug {
				log.Println("read config noheader:", config_key, config_value)
			}
			var found bool = false
			for resp_key, resp_value := range resp.Header {
				if strings.EqualFold(config_key, resp_key) {
					if *debug {
						log.Println("Header", resp_key, "with value", resp_value[0], "in response match config")
					}
					found = true
					break
				}
			}
			if found {
				log.Fatal(config.Url, ": header '", config_key, "' was found in the response", resp)
			}
		}

		// body
		if config.Body != "" {
			log.Println("WARNING: 'body' is not yet supported")
		}
		if config.Nobody != "" {
			log.Println("WARNING: 'nobody' is not yet supported")
		}

	}

	log.Println("All headers with values found")
}
