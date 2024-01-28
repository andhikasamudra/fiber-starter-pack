## Fiber Starter Pack

build with [Fiber](https://gofiber.io/) and [Bun](https://bun.uptrace.dev/) and [GRPC](https://grpc.io/)

setup to your local env first

```shell
cp env-example.sh local_env.sh
source local_env.sh
```

get go package installed

```shell
go get 
```

to run 

```shell
go run main.go
```

OR

using docker compose 
```shell
docker-compose up
```

## Architecture

```shell
.
├── ...
└── pkg/
    └── domainName/
        ├── dto
        ├── models
        ├── services
        ├── handler
        └── route.go
```

in this architecture refers from domain driven design principles
the domain grouped in folder `pkg` 
in `pkg` there is `domain` folder to grouped each related domain modules 
such as `services`, `handler`, `models`, `dto`
the `route.go` file will be represent of each domain api routes. 

### Why every domain should have different routes ?

well the perfect case is when you need to separate this domain into multiple services 
you can just "put this out" into new services