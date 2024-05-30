# Wolfios Docker Buildpack

::: warning
This project is in early development stage, so it may not work as expected.

The idea of this project is to generate for projects a customized `Dockerfile` using Wolfi Docker images.

Right now supported languages are:

- Nodejs
- PHP

## Usage

```bash
go run . <project-name>
```

and then you will have a `Dockerfile` in the root of your project.
