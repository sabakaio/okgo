# okgo

![logo](https://cloud.githubusercontent.com/assets/822951/12686689/13537922-c6c3-11e5-9420-93824fd63928.png)

Schedule everything

## Usage
```
NAME:
   okgo task - work with okgo tasks

USAGE:
   okgo task command [command options] [arguments...]

COMMANDS:
   create, c  create a new task
   list, ls list all defined tasks
   remove, rm remove an existing task
   run, r run a task
   purge  remove all tasks
   help, h  Shows a list of commands or help for one command
```

## HTTP API
Start a server
```
$ okgo server
```
Use API
```
GET    /api/v1/tasks         List all tasks
GET    /api/v1/tasks/:name   Get task by name
POST   /api/v1/tasks         Create a new task {"name": "...", "command": "..."}
DELETE /tasks/:name          Delete a task by name
```
