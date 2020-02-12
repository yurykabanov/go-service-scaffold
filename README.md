# Some Service

## Overview

TODO: Some service description

## Quickstart

1. Prepare config file using [example](./config/some_service.example.yml) 

2. (optional) Prepare Jaeger-related [environment variables](https://github.com/jaegertracing/jaeger-client-go)
if you want to collect traces

2. Run server with:
```bash
# Assuming config in /path/to/config.yml
./server -c /path/to/config.yml
```
or
```bash
# Assuming config in one of the following places:
# 1. ./some_service.yml 
# 2. ./config/some_service.yml
# 3. /etc/some_service
./server
```

See the next section for more information about configuration.

## Configuration

This project utilizes [Viper](https://github.com/spf13/viper) which provides wide
variety of ways to configure application:
- it supports JSON, TOML, YAML, HCL or Java properties formats
- it provides ability to use environment variables, command line flags and
configuration files

Viper will use the following [precedence](https://github.com/spf13/viper#why-viper):
- flags
- environment variables
- configuration file
- defaults

### Command line flags

The application provides single flag:
- `-c /path/to/config.yml` - override config path.

### Environment variables

Environment variables use the following naming convention: any flag with name
`some.flag.value` corresponds to `SOME_SERVICE_SOME_FLAG_VALUE` variable (note the
prefix `SOME_SERVICE_`).

The application also uses [Jaeger](https://www.jaegertracing.io/) to collect traces,
see Go client's [documentation](https://github.com/jaegertracing/jaeger-client-go#environment-variables)
for more details.

### Configuration files

This service will try to search config file in `./`, `./config/` and
`/etc/some_service/` directories with names `some_service.*` (any format
described above will work).

Example of configuration file in YAML is provided in `./config/` directory.

## Build from scratch

Run build script:

```bash
./scripts/build.sh
```
