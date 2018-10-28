package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// UserAgent The http User-Agent used for testing
const UserAgent = "Golang HeadersCheck/0.1 (gui@iroqwa.org)"

// Configuration is the mapping of configs containing Scenario
type Configuration struct {
	URLs map[string]Scenario
}

// Scenario contain various resource to fetch and verify
type Scenario struct {
	URL       string
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

func fetchURL(url string, useragent string, followRedirect bool) *http.Response {

	// return the error, so client won't attempt redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 5 * time.Second,
	}
	if followRedirect {
		client = &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching url: ", err)
	}

	return resp
}

func main() {

	debug := flag.Bool("debug", false, "Enable debugging output")
	configFile := flag.String("config-file", "config", "Config file")
	userAgent := flag.String("user-agent", UserAgent, "User-Agent used for queries")
	followRedirect := flag.Bool("follow-redirect", false, "Follow redirect (http status codes 30X)")
	flag.Parse()

	configuration := readConfig(*configFile)
	if *debug {
		log.Println("[*] Debug enabled")
		log.Println("[*] Config read from", *configFile, "file:", configuration.URLs)
	}

	for _, config := range configuration.URLs {

		if *debug {
			log.Println("[*] Fetching url", config.URL, "with UA", *userAgent)
		}
		resp := fetchURL(config.URL, *userAgent, *followRedirect)

		// code
		if config.Code != strconv.Itoa(resp.StatusCode) {
			log.Fatal("Error: expected status code ", config.Code, " found ", resp.StatusCode)
		} else {
			if *debug {
				log.Println("Code", resp.StatusCode, "match config", config.Code)
			}
		}

		// headers
		for configKey, configValue := range config.Headers {
			if *debug {
				log.Println("read config header:", configKey, configValue)
			}
			var found = false
			for respKey, respValue := range resp.Header {
				if strings.EqualFold(configKey, respKey) {
					if strings.EqualFold(respValue[0], configValue) {
						if *debug {
							log.Println("Header", respKey, "with value", respValue[0], "in response match config")
						}
						found = true
						break
					}
				}
			}
			if !found {
				log.Fatal("Error", config.URL, ": header '", configKey, "' with value '", configValue, "' was not found in the response", resp)
			}
		}

		// noheaders
		for configKey, configValue := range config.Noheaders {
			if *debug {
				log.Println("read config noheader:", configKey, configValue)
			}
			var found = false
			for respKey, respValue := range resp.Header {
				if strings.EqualFold(configKey, respKey) {
					if *debug {
						log.Println("Header", respKey, "with value", respValue[0], "in response match config")
					}
					found = true
					break
				}
			}
			if found {
				log.Fatal("Error", config.URL, ": header '", configKey, "' was found in the response", resp)
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

	log.Println("All scenario executed successfully")
}
