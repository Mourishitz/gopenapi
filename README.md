# GOpenAPI

Merge multiple YAML files into one single API spec with just the run of a command!

With this project it is possible to separate Open API specs using different YAML files. It uses a `dirs.json` so you can declarate paths your own way.

It also contains examples on how to structure YAML files inside `src`. This repository works as a template to build your first GOpenAPI project. It comes with a simple but fully documented API that can be edited or used as examples on how to configure your API spec.

The project can compile back to a single YAML file on the root of the project called `swagger.yaml`.

An additional bundle was added to serve this YAML to an index.html. This bundle can be found inside `dist` and it contains the official [Swagger UI static distribution](https://github.com/swagger-api/swagger-ui/blob/master/dist/index.html).


## Install

### Using Golang

Install using the release tag

```bash
    go install github.com/mourishitz/gopenapi/gopenapi@release
```

If you wish to, you can also clone the project and use it as a template to your first Open API spec.


### From source

Clone the project

```bash
  git clone https://github.com/Mourishitz/GOpenAPI
```

Go to the project directory

```bash
  cd GOpenAPI
```

Compile your YAML files

```bash
  go build -o ./build ./cmd
```

Run the binary (or add it to your $PATH)

```bash
  ./build
```

