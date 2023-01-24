## Fiber Starter Pack

build with [Fiber](https://gofiber.io/)

setup to your local env first

```shell
cp env-example.sh local_env.sh
```

get go package installed

```shell
go get 
```

to run 

```shell
go run main.go
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

#### why every domain should have different routes ?

the only i think that can do is whenever you need to separate this monolithic service into multiple services
you can just plug the `domain` folder and put it in new service and you ready to run

### Todo
- [ ] add new event driven connection
- [ ] add another domain layer and try communicate with event
