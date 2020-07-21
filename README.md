# TService (a.k.a test-service)

It's a dummy service that helps to easily develop a new API that uses third-party API.

By default, it returns a status 200 for any request + dumps a request into stdout.

# Usage

The easiest way to use the TService is to run it with the docker-compose:
`$ docker-compose up`

It also mounts the `configs` and `assets` folders where you can configure and put your responses.

# Configuration

Example of the config.yml file:

```
responses:
   response1:
      path: /lorem/ipsum
      definition:
         status_code: 404
         response: '{"resource":"not-found"}'
   response2:
      path: /lorem
      definition:
         status_code: 200
         response_file: lorem.json
      slowness:
         chance: 30
         duration: 5s
   response3:
      path: /lorem/error
      definition:
         status_code: 200
         response_file: lorem.json
      error:
         chance: 10
         status_code: 500
```      

In that example, we defined three possible responses.

### response1

It returns `{"resource":"not-found"}` when tservice will be requested by the path `/lorem/ipsum`

### response2

It returns the content of the file `lorem.json` that is located in `/assets` folder when TService will be requested by the path `/lorem`.

With the change 30 of 100 it will send a response with a timeout of the 5 seconds.

### response3

It returns the content of the file `lorem.json` that is located in `/assets` folder when TService will be requested by the path `/lorem/error`.

With the change 10 of 100 it will send a response with 500 status code as the error.

# Example of the build

`$ make app_build && bin/tservice --config=configs/config.yml --assets=assets`