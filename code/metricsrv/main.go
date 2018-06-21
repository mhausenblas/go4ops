package main

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		v, _ := mem.VirtualMemory()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(struct {
			MemoryTotal uint64 `json:"memory_total"`
		}{
			v.Total,
		})
	})
	_ = http.ListenAndServe(":9876", nil)
}
