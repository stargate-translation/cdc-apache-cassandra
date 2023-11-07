package main

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
)

func main() {
    http.HandleFunc("/alter", func(w http.ResponseWriter, r *http.Request) {
        // Parse the query parameter "table" from the URL
        table := r.URL.Query().Get("table")
        if table == "" {
            http.Error(w, "Missing 'table' parameter", http.StatusBadRequest)
            return
        }

        // Execute the CQL command using cqlsh
        cmd := exec.Command(
            "docker", "exec", "cassandra-1", "cqlsh", "-e", fmt.Sprintf("ALTER TABLE %s WITH cdc=true;", table))
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err := cmd.Run()
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to run cqlsh: %v", err), http.StatusInternalServerError)
            return
        }

        // Respond with a success message
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Table altered with CDC enabled"))
    })

    // Start the HTTP server
    if err := http.ListenAndServe(":7111", nil); err != nil {
        fmt.Println("Error:", err)
    }
}
