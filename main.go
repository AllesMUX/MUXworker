package MUXworker

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/AllesMUX/MUXworker/tasks"
    "github.com/AllesMUX/MUXworker/cpu"
)

type ServerStatus struct {
    CPULoadAvg float64 `json:"cpu_load_avg"`
    ActiveTasks int `json:"active_tasks"`
}

func IncrementTasks() {
    tasks.IncrementTasks()
}

func DecrementTasks() {
    tasks.DecrementTasks()
}

func NewTasksCountService(port string) {
    go func() {
        for {
            cpu.UpdateCPUStats()
        }
    }()
    
    http.HandleFunc("/server-health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(ServerStatus{
            CPULoadAvg: cpu.GetCPUStats().LoadAvg,
            ActiveTasks: tasks.GetActiveTasks().Count,
        })
    })

    fmt.Printf("Service is listening on port %s...\n", port)
    go http.ListenAndServe(port, nil)
}