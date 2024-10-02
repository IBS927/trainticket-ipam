package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServiceInfo struct {
	IP   string
	Node string
	Port string
}
/*
var services = map[string]ServiceInfo{
	"ts-ui-dashboard": {IP: "192.168.11.11", Node: "192.168.11.101", Port: "8080"},
	"ts-gateway-service": {IP: "192.168.11.12", Node: "192.168.11.101", Port: "18888"},
	"ts-travel-mysql":{IP: "192.168.11.13", Node: "192.168.11.101", Port: "3306"},
	"ts-travel-service":{IP: "192.168.11.14", Node: "192.168.11.101", Port: "12346"},
	"ts-basic-service":{IP: "192.168.11.15", Node: "192.168.11.101", Port: "15680"},
	"ts-seat-service":{IP: "192.168.11.16", Node: "192.168.11.101", Port: "18898"},
	"ts-station-mysql":{IP: "192.168.11.17", Node: "192.168.11.101", Port: "3306"},
	"ts-station-service":{IP: "192.168.11.18", Node: "192.168.11.101", Port: "12345"},
	"ts-train-mysql":{IP: "192.168.11.19", Node: "192.168.11.102", Port: "3306"},
	"ts-train-service":{IP: "192.168.11.20", Node: "192.168.11.102", Port: "14567"},
	"ts-price-mysql":{IP: "192.168.11.21", Node: "192.168.11.102", Port: "3306"},
	"ts-price-service":{IP: "192.168.11.22", Node: "192.168.11.102", Port: "16579"},
	"ts-order-mysql":{IP: "192.168.11.23", Node: "192.168.11.102", Port: "3306"},
	"ts-order-service":{IP: "192.168.11.24", Node: "192.168.11.102", Port: "12031"},
	"ts-config-mysql":{IP: "192.168.11.25", Node: "192.168.11.102", Port: "3306"},
	"ts-config-service":{IP: "192.168.11.26", Node: "192.168.11.102", Port: "15679"},
	"ts-route-mysql":{IP: "192.168.11.27", Node: "192.168.11.102", Port: "3306"},
	"ts-route-service":{IP: "192.168.11.28", Node: "192.168.11.102", Port: "11178"},
	"zipkin":{IP: "192.168.11.41", Node: "192.168.11.101", Port: "9411"},
}
var services_no_sql = map[string]ServiceInfo{
	"ts-ui-dashboard": {IP: "192.168.11.11", Node: "192.168.11.101", Port: "8080"},
        "ts-gateway-service": {IP: "192.168.11.12", Node: "192.168.11.101", Port: "18888"},
        "ts-travel-service":{IP: "192.168.11.14", Node: "192.168.11.101", Port: "12346"},
        "ts-basic-service":{IP: "192.168.11.15", Node: "192.168.11.101", Port: "15680"},
        "ts-seat-service":{IP: "192.168.11.16", Node: "192.168.11.101", Port: "18898"},
        "ts-station-service":{IP: "192.168.11.18", Node: "192.168.11.101", Port: "12345"},
        "ts-train-service":{IP: "192.168.11.20", Node: "192.168.11.102", Port: "14567"},
        "ts-price-service":{IP: "192.168.11.22", Node: "192.168.11.102", Port: "16579"},
        "ts-order-service":{IP: "192.168.11.24", Node: "192.168.11.102", Port: "12031"},
        "ts-config-service":{IP: "192.168.11.26", Node: "192.168.11.102", Port: "15679"},
        "ts-route-service":{IP: "192.168.11.28", Node: "192.168.11.102", Port: "11178"},

}
var services_mysql= map[string]ServiceInfo{
	"ts-travel-mysql":{IP: "192.168.11.13", Node: "192.168.11.101", Port: "3306"},
	"ts-station-mysql":{IP: "192.168.11.17", Node: "192.168.11.101", Port: "3306"},
	"ts-train-mysql":{IP: "192.168.11.19", Node: "192.168.11.102", Port: "3306"},
	"ts-price-mysql":{IP: "192.168.11.21", Node: "192.168.11.102", Port: "3306"},
	"ts-order-mysql":{IP: "192.168.11.23", Node: "192.168.11.102", Port: "3306"},
	"ts-config-mysql":{IP: "192.168.11.25", Node: "192.168.11.102", Port: "3306"},
	"ts-route-mysql":{IP: "192.168.11.27", Node: "192.168.11.102", Port: "3306"},
	"zipkin":{IP: "192.168.11.41", Node: "192.168.11.101", Port: "9411"},
}*/
var services = map[string]ServiceInfo{
	"nginx_r":{IP: "192.168.11.11", Node: "192.168.11.101", Port: "80"},
	"nginx_b":{IP: "192.168.11.12", Node: "192.168.11.102", Port: "80"},
}
var services_no_sql = map[string]ServiceInfo{
	"nginx_r":{IP: "192.168.11.11", Node: "192.168.11.101", Port: "80"},
        "nginx_b":{IP: "192.168.11.12", Node: "192.168.11.102", Port: "80"},
}
var services_mysql= map[string]ServiceInfo{}

var ipPairs = map[string]string{
	"192.168.11.102": "192.168.11.202",
	"192.168.11.101": "192.168.11.201",
}
func ipHandler(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Path[len("/ip/"):]

	if service, ok := services[serviceName]; ok {
		w.Header().Set("content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(service.IP); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	} else {
		http.Error(w, "err: Service not found", http.StatusNotFound)
	}
}

// nodeHandler は指定されたサービスのノード名を返します
func nodeHandler(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Path[len("/node/"):]

	if service, ok := services[serviceName]; ok {
		w.Header().Set("content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(service.Node); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
		}
	} else {
		http.Error(w, "err: Service not found", http.StatusNotFound)
	}
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(services); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-Type", "application/json")

        encoder := json.NewEncoder(w)
        if err := encoder.Encode(services_no_sql); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte("Internal server error"))
        }
}

func mysqlHandler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("content-Type", "application/json")

        encoder := json.NewEncoder(w)
        if err := encoder.Encode(services_mysql); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte("Internal server error"))
        }
}

func snic_pair(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "application/json")

        encoder := json.NewEncoder(w)
        if err := encoder.Encode(ipPairs); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte("Internal server error"))
        }
}
func main() {
	http.HandleFunc("/ip/", ipHandler)
	http.HandleFunc("/node/", nodeHandler)
	http.HandleFunc("/all", allHandler)
	http.HandleFunc("/all_service",serviceHandler)
	http.HandleFunc("/all_mysql",mysqlHandler)
	http.HandleFunc("/snic_pair",snic_pair)
	fmt.Println("Server starting on port 8000...")
	http.ListenAndServe(":8000", nil)
}
