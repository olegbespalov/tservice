# TService (a.k.a test-service)

Introduction
------------

The TService is the test-service, a fake API that you can use to mock third-party API. It's developed lightweight, has no dependency on any programming language or framework.  It can define service slowness emulate service errors. Config or responses changes on the fly without any service restarting.

<p align="center"><img src="/assets/demo.gif?raw=true"/></p>

Compatibility
-------------

The TService can be used as go application that compiled on a host machine, docker-compose or directly from the docker hub image.

Installation and usage
-------------

To install it, run:

   docker-compose up

You can also include the service in your existing compose using [docker hub's image](https://hub.docker.com/repository/docker/letniy/tservice)

```yml
tservice:
   image: docker.io/letniy/tservice:latest
   ports:
      - "8080:8080"
   volumes:
      - ./configs:/configs
      - ./responses:/assets
   networks:
      app_net:
   entrypoint: ["/bin/app", "-config", "/configs/config.yml", "-assets", "/assets"]
```

A full example located in the repository [olegbespalov/tservice-example](https://github.com/olegbespalov/tservice-example).

or you can run locally if you have go installed

   make app_build && bin/tservice --config=configs/config.example.yml --assets=assets

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