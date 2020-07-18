# TService (a.k.a test-service)

It's a dummy service that helps to develop a new API that uses third-party API.

For now it always return a status 200 for any request + dump a request into stdout.

To define a response you can just put a content of the response as the file, where name is the path to the resource.

# Example of the build

`$ make app_build && bin/tservice --config=configs/config.yml --assets=assets`