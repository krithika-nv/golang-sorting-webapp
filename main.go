package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"strings"
	"fmt"
	"sort"
)

var (
	number_of_requests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_of_requests",
		Help: "Number of http requests.",
	},
))

var port1,port2 string

func init() {
	prometheus.MustRegister(number_of_requests)
	port1 = os.Getenv("PORT")
	if port1 == "" {
		port1 = "8080"
	}

    port2 = os.Getenv("PROM_PORT")
	if port2 == "" {
		port2 = "9110"
	}
}


func increment(res http.ResponseWriter, req *http.Request) {
	if !strings.Contains(req.URL.String(),"favicon.ico") { //Check to avoid counting /favicon.ico request
	fmt.Fprintf(res, "%s", "<table border=\"1\">")
	fmt.Fprintf(res, "%s", "<th>Header Field</th>")
	fmt.Fprintf(res, "%s", "<th>Header Value</th>")
	
    hmap := map[string]string{}
    var h_key []string
    for k, v := range req.Header {
    	hmap[k] = strings.Join(v,",")	
    	h_key = append(h_key, fmt.Sprintf("%s", k))
    }
    
    sort.Strings(h_key)

    for i:=0;i<len(h_key);i++{
    	fmt.Fprintf(res, "%s", "<tr><td>")
        fmt.Fprintf(res, "%v", h_key[i])
        fmt.Fprintf(res, "%s", "</td><td>")
        fmt.Fprintf(res, "%v", hmap[h_key[i]])
        fmt.Fprintf(res, "%s", "</td></tr>")
    }

    fmt.Fprintf(res, "%s", "</table>")
    number_of_requests.Inc()
	}
}

func main() {
    serverMuxA := http.NewServeMux()
    serverMuxA.HandleFunc("/", increment)

    serverMuxB := http.NewServeMux()
    serverMuxB.Handle("/metrics", promhttp.Handler())

    go func() {
        log.Fatal(http.ListenAndServe("0.0.0.0:"+port1, serverMuxA))
    }()

    log.Fatal(http.ListenAndServe("0.0.0.0:"+port2, serverMuxB))
}
