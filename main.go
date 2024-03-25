package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "runtime"
    "time"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    // Aktuelle Zeit für Timestamps
    now := time.Now().Unix()

    // Simulierte Metriken
    cpuUsage := 50.0 + rand.Float64()*50.0                  // CPU-Auslastung
    ramUsage := 40.0 + rand.Float64()*60.0                  // RAM-Auslastung
    diskUsage := 20.0 + rand.Float64()*80.0                 // Disk-Auslastung
    networkIn := 100.0 + rand.Float64()*900.0               // Netzwerkeingang
    networkOut := 100.0 + rand.Float64()*900.0              // Netzwerkausgang
    errorRate := rand.Float64() * 5.0                       // Fehlerquote
    responseTime := 100.0 + rand.Float64()*150.0            // Antwortzeit
    dbConnectionCount := rand.Intn(100)                     // DB-Verbindungen
    cacheHitRate := 70.0 + rand.Float64()*30.0              // Cache-Trefferquote
    goroutineCount := runtime.NumGoroutine()                // Goroutines
    threadCount := runtime.GOMAXPROCS(0)                    // Threads
    requestCount := rand.Intn(1000)                         // Anfragen
    successfulRequests := int(float64(requestCount) * (1 - errorRate/100)) // Erfolgreiche Anfragen
    failedRequests := requestCount - successfulRequests     // Fehlgeschlagene Anfragen
    heapAlloc := rand.Uint64() % 1000000000                 // Heap-Allokation
    mallocs := rand.Uint64() % 1000000                      // Mallocs
    frees := rand.Uint64() % 1000000                        // Frees
    liveObjects := mallocs - frees                          // Lebende Objekte
    gcPause := rand.Float64() * 100.0                       // GC-Pause
    lastGC := now - rand.Int63n(3600)                       // Letzte GC-Zeit

    metrics := fmt.Sprintf(`
# HELP cpu_usage CPU usage percentage
# TYPE cpu_usage gauge
b1_cpu_usage{service="myService"} %f
# HELP ram_usage RAM usage percentage
# TYPE ram_usage gauge
b1_ram_usage{service="myService"} %f
# HELP disk_usage Disk usage percentage
# TYPE disk_usage gauge
b1_disk_usage{service="myService"} %f
# HELP network_in Incoming network traffic in Mbps
# TYPE network_in gauge
b1_network_in{service="myService"} %f
# HELP network_out Outgoing network traffic in Mbps
# TYPE network_out gauge
b1_network_out{service="myService"} %f
# HELP error_rate Error rate percentage
# TYPE error_rate gauge
b1_error_rate{service="myService"} %f
# HELP response_time Average response time in ms
# TYPE response_time gauge
b1_response_time{service="myService"} %f
# HELP db_connection_count Number of database connections
# TYPE db_connection_count gauge
b1_db_connection_count{service="myService"} %d
# HELP cache_hit_rate Cache hit rate percentage
# TYPE cache_hit_rate gauge
b1_cache_hit_rate{service="myService"} %f
# HELP goroutine_count Number of goroutines
# TYPE goroutine_count gauge
b1_goroutine_count{service="myService"} %d
# HELP thread_count Number of threads
# TYPE thread_count gauge
b1_thread_count{service="myService"} %d
# HELP request_count Number of requests
# TYPE request_count counter
b1_request_count{service="myService"} %d
# HELP successful_requests Number of successful requests
# TYPE successful_requests counter
b1_successful_requests{service="myService"} %d
# HELP failed_requests Number of failed requests
# TYPE failed_requests counter
b1_failed_requests{service="myService"} %d
# HELP heap_alloc Bytes allocated and still in use
# TYPE heap_alloc gauge
b1_heap_alloc{service="myService"} %d
# HELP mallocs Total number of mallocs
# TYPE mallocs counter
b1_mallocs{service="myService"} %d
# HELP frees Total number of frees
# TYPE frees counter
b1_frees{service="myService"} %d
# HELP live_objects Number of live objects (mallocs - frees)
# TYPE live_objects gauge
b1_live_objects{service="myService"} %d
# HELP gc_pause Total GC pause time in ms
# TYPE gc_pause gauge
b1_gc_pause{service="myService"} %f
# HELP last_gc Time of last garbage collection in Unix timestamp
# TYPE last_gc gauge
b1_last_gc{service="myService"} %d
`, cpuUsage, ramUsage, diskUsage, networkIn, networkOut, errorRate, responseTime, dbConnectionCount,
        cacheHitRate, goroutineCount, threadCount, requestCount, successfulRequests, failedRequests,
        heapAlloc, mallocs, frees, liveObjects, gcPause, lastGC)

    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprint(w, metrics)
}

func main() {
    http.HandleFunc("/metrics", metricsHandler)
    http.ListenAndServe(":8080", nil)
}
