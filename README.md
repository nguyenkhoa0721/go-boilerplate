## Project structures

1. `/api` Document for api. OpenAPI or Swagger specifications, JSON Schema files, protocol definition files.
2. `/build` Script for build. Docker file for local, dev, production
3. `/cmd` The entry point for our application
4. `/config` Initialization of the general app configurations
5. `/internal` Internal logic of application. Internal contain module which has:
    1. `/presenter` Presenter layer. It is named with user domain, such as: public/http_handler.go for user, admin/http_handler.go for admin
    2. `/domain` Use case layer
        * `service` Business's logic
        * `validation` Define use case input/DTO
        * `interface` Define use case output
    3. `/data` Data layer
        * `*_repo` Repository 

6. `/pkg`
7. `/script` Scripts for migration, seeding,...

## Generator
```
go run ./cmd/cli/main.go
```