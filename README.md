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


## Quickstart

This project requires only one file to work and that is `dirs.json`.
Inside `dirs.json` you must have all OpenAPI keys that are necessary to build a great documentation

An example of valid `dirs.json` would be:
```json
{
  "base": "template.yml",
  "info": "info.yml",
  "security": "security.yml",
  "tags": "./modules/tags",
  "paths": "./modules/paths",
  "components": "./modules/components",
  "schemas": "./modules/schemas",
  "requests": "./modules/requests",
  "definitions": "./modules/definitions",
  "servers": "servers.yml"
}
```

This also means all these files **must** exist.

Note that `tags`, `paths`, `components`, `schemas`, `requests` and `definitions` are all directories and therefore, all of the yaml files inside them are considered as one of their respectives. This means that you can have X files with any name inside `paths/` and as long as they are valid OpenAPI path definitions they are be good to go!

