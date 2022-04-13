# Golang : Use gorm with mysql in gin

This repository guides to how ORM can be implemented in Golang. After cloning the code, follow below steps to let the project run on your system.

1. Run `go build main.go`

2. Run `go run main.go`

After running the commands, you can find server started on `8080` .


go的项目结构参考：
```
.
├── .air.conf                # Config file for air.
├── Dockerfile
├── Makefile
├── README.md
├── apiary.apib              # API Documentation.
├── cmd                      # Main applications for this project.
│   ├── migration            # Database migration.
│   └── server
│       ├── main.go          # Entry point of the API Server.
│       └── routes           # Registering route handlers and middleware.
├── configs                  # Contains all configuration files.
├── docs                     # Design and user documents.
├── go.mod
├── go.sum
├── internal                 # Private application and library code.
│   ├── consts               # Provides global constants.
│   ├── entity               # Core Layer & Data Serialization and Deserialization.
│   ├── handler              # Presentation Layer, defines handler methods of endpoints.
│   ├── repository           # Persistence Layer.
│   ├── service              # Interactors Layer.
│   └── structure            # Data Transfer Object.
│       ├── request
│       └── response
├── pkg                      # Library code that's ok to use by external applications.
│   └── initializer          # Application configuration loader.
├── test                     # Additional external test apps and test data.
└── web                      # Web application specific components: static web assets, server side templates and SPAs.
```
