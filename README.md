# MUXworker
MUXworker is a Go package that provides a simple HTTP server for monitoring the health of a worker service. The server exposes a single endpoint, /server-health, that returns a JSON object containing the current CPU load average and the number of active tasks.

## Installation
To install MUXworker, use the following command:

`go get github.com/AllesMUX/MUXworker`

## Usage
To use MUXworker in your worker service, first import the package:

`import "github.com/AllesMUX/MUXworker"`

Then, call the `MUXworker.NewTasksCountService` function to start the HTTP server. The function takes a single argument, `port`, which specifies the port number that the server should listen on. For example:

`MUXworker.NewTasksCountService(":8080")`

This will start the HTTP server on port 8080. You can then send a GET request to the `/server-health` endpoint to retrieve the current health status of the worker service.

To manage tasks in your worker service, you can use functions, `IncrementTasks` and `DecrementTasks`, that allow you to increment or decrement the number of active tasks, respectively. For example:
```go
var tm = MUXworker.TasksManager{}
tm.NewTasksCountService(":8888")
var tasks = make(chan func(), 10)
tasks <- func() {
	tm.IncrementTasks()
	/* task code */
	tm.DecrementTasks()
}
```

And when you need to get the current server status, send a GET request to `/server-health`, you'll get JSON in return, for example:
```json
{
    "cpu_load_avg":20.041666666899495,
    "active_tasks":1
}
```

## Contributing
Contributions to MUXworker are welcome! If you have an idea for a new feature or have found a bug, please open an issue on the [GitHub issue tracker](https://github.com/AllesMUX/MUXworker/issues).

If you would like to contribute code, please fork the repository and submit a pull request.

## License
MUXworker is licensed under the [Apache License 2.0](https://github.com/AllesMUX/MUXworker/blob/main/LICENSE).