package MUXworker

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/AllesMUX/MUXworker/cpu"
    "github.com/AllesMUX/MUXworker/structs"
)


type TasksManager struct {
    count int
}

func (tm *TasksManager) IncrementTasks() {
    fmt.Printf("incrim %d\n",tm.count)
    tm.count++
}

func (tm *TasksManager) DecrementTasks() {
    fmt.Printf("decrim %d\n",tm.count)
    tm.count--
}

func (tm *TasksManager) GetActiveTasks() int {
    return tm.count
}

func (tm *TasksManager) NewTasksCountService(port string) {
    tm.count = 0
    go func() {
        for {
            cpu.UpdateCPUStats()
        }
    }()
    
    http.HandleFunc("/server-health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        json.NewEncoder(w).Encode(structs.ServerStatusJSON{
            CPULoadAvg: cpu.GetCPUStats().LoadAvg,
            ActiveTasks: tm.GetActiveTasks(),
        })
    })

    fmt.Printf("MUXworker service is listening on port %s...\n", port)
    go http.ListenAndServe(port, nil)
}