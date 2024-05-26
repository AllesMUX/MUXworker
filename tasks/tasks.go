package tasks

type activeTasksStruct struct {
    Count int `json:"count"`
}

var activeTasks = activeTasksStruct{
    Count: 0,
}

func IncrementTasks() {
    activeTasks.Count++
}

func DecrementTasks() {
    activeTasks.Count--
}

func GetActiveTasks() activeTasksStruct {
    return activeTasks
}
