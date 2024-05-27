package structs

type ServerStatusJSON struct {
    CPULoadAvg float64 `json:"cpu_load_avg"`
    ActiveTasks int `json:"active_tasks"`
}