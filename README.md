Schedule everything

## Usage
```
NAME:
   okgo job - work with okgo jobs

USAGE:
   okgo job command [command options] [arguments...]

COMMANDS:
   create, c  create a new job
   list, ls list all defined jobs
   remove, rm remove an existing job
   run, r run a job
   purge  remove all jobs
   help, h  Shows a list of commands or help for one command
```

## HTTP API
Start a server
```
$ okgo server
```
Use API
```
GET    /api/v1/jobs         List all jobs
GET    /api/v1/jobs/:name   Get job by name
POST   /api/v1/jobs         Create a new job {"name": "...", "command": "..."}
DELETE /jobs/:name          Delete a job by name
```
