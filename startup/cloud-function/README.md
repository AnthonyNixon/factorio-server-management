# Server Startup cloud function
Calling this cloud function will start up the server if it's currently shut down.

## Settings
Memory: 128 MiB
Timeout: 10
Runtime: Go 1.13

### Env Vars
`PROJECT_ID`: Project ID of the Google cloud Project
`REGION`: Region the instance is deployed in
`ZONE`: Zone the instance is deployed in
`INSTANCE_NAME`: Name of the server instance