package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

// Metriken für HTTP-Methoden und Statuscodes simulieren
var (
    methodCounts = map[string]int{
        "GET":     0,
        "POST":    0,
        "UPDATE":  0,
        "PUT":     0,
        "DELETE":  0,
        "PATCH":   0,
    }
    statusCodesCounts = map[int]int{
        200: 0,
        403: 0,
        404: 0,
        500: 0,
    }
    mutex sync.Mutex
)

func recordMetrics(method string, statusCode int) {
    mutex.Lock()
    defer mutex.Unlock()
    methodCounts[method]++
    statusCodesCounts[statusCode]++
}

func simulateTraffic() {
    for {
        // Simuliere eine zufällige Anfrage
        methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
        statuses := []int{200, 404, 500}
        selectedMethod := methods[rand.Intn(len(methods))]
        selectedStatus := statuses[rand.Intn(len(statuses))]
        recordMetrics(selectedMethod, selectedStatus)
        time.Sleep(1 * time.Second)
    }
}
func handler(w http.ResponseWriter, r *http.Request) {
    // Hier könnte Ihre Logik zur Verarbeitung der Anfrage stehen.
    // Zum Beispiel könnte basierend auf der Anfrage eine Aktion ausgeführt werden,
    // und es könnte eine Antwort mit einem spezifischen Statuscode zurückgegeben werden.
    
    // Beispiel: Antwort basierend auf der Methode der Anfrage
    switch r.Method {
    case "GET":
        // Simuliere eine erfolgreiche GET-Anfrage
        w.WriteHeader(http.StatusOK) // 200 Statuscode
        fmt.Fprintln(w, "GET Anfrage erfolgreich")
        recordMetrics(r.Method, http.StatusOK) // Aufzeichnung der Metrik
    case "POST":
        // Simuliere eine fehlerhafte POST-Anfrage
        w.WriteHeader(http.StatusInternalServerError) // 500 Statuscode
        fmt.Fprintln(w, "POST Anfrage fehlgeschlagen")
        recordMetrics(r.Method, http.StatusInternalServerError) // Aufzeichnung der Metrik
    default:
        // Für andere Methoden geben wir einfach einen 405 Method Not Allowed zurück.
        w.WriteHeader(http.StatusMethodNotAllowed) // 405 Statuscode
        fmt.Fprintf(w, "Methode %s nicht erlaubt", r.Method)
        recordMetrics(r.Method, http.StatusMethodNotAllowed) // Aufzeichnung der Metrik
    }
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()
    fmt.Fprintln(w, "# HELP http_requests_total The total number of HTTP requests.")
    fmt.Fprintln(w, "# TYPE http_requests_total counter")
    for method, count := range methodCounts {
        fmt.Fprintf(w, "b1_http_requests_total{method=\"%s\"} %d\n", method, count)
    }

    fmt.Fprintln(w, "# HELP http_responses_total The total number of HTTP responses by status code.")
    fmt.Fprintln(w, "# TYPE http_responses_total counter")
    for status, count := range statusCodesCounts {
        fmt.Fprintf(w, "b1_http_responses_total{status=\"%d\"} %d\n", status, count)
    }
    fmt.Fprintln(w, "# HELP http_request_responses_total The total number of HTTP responses by status code.")
    fmt.Fprintln(w, "# TYPE http_request_responses_total counter")
    for method := range methodCounts {
      for status := range statusCodesCounts {
          count := rand.Intn(1000)
          fmt.Fprintf(w, "b1_http_request_responses_total{status=\"%d\", method=\"%s\"} %d\n", status, method, count)
      }
    }
}


func main() {
    rand.Seed(time.Now().UnixNano())
    go simulateTraffic() // Startet die Simulation von Traffic

    http.HandleFunc("/metrics", metricsHandler)
    http.HandleFunc("/", handler) // Setzt den Handler für die Wurzel-URL

    http.ListenAndServe(":8080", nil)
}
