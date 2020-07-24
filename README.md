# TService (a.k.a test-service)

Introduction
------------

The TService is the test-service, a fake API that you can use to mock third-party API. It's developed lightweight, has no dependency on any programming language or framework. It can define service slowness emulate service errors. A config or a response changes on the fly without any service restarting.

<p align="center"><img src="/assets/usage.gif?raw=true"/></p>

Compatibility
-------------

The TService can be used as go application that compiled on a host machine, docker-compose or directly from the docker hub image.

Installation and usage
-------------

To launch it, run:

    docker-compose up

You can also include the service in your existing compose using [docker hub's image](https://hub.docker.com/repository/docker/letniy/tservice)

```yml
tservice:
   image: docker.io/letniy/tservice:latest
   ports:
      - "8085:8085"
   volumes:
      - ./configs:/configs
   networks:
      app_net:
   entrypoint: ["/bin/app", "-config", "/configs/config.yml", "-responsePath", "/configs/responses", "-port", "8085"]
```

A full example located in the repository [olegbespalov/tservice-example](https://github.com/olegbespalov/tservice-example).

or you can run locally if you have go installed

    make app_build && bin/tservice --config=configs/config.example.yml --responsePath=configs/responses --port=8085

Configuration
-------------

An example of the configuration you can find in the configs directory.

Example of the config.yml file:

```yaml
responses:
   response1:
      path: /lorem/ipsum
      definition:
         status_code: 200
         response: '{"hello":"TService"}'
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

License
-------

The TService package is licensed under the MIT. Please see the LICENSE file for details.