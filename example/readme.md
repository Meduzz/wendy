# An example service

This is a super simple example of a wendy service. It is a very naive service discovery api.

## Add a service

    POST localhost:8080/api

    {
        "module":"service",
        "method":"add",
        "body": {
            "name":"service-discovery",
            "tags":["awesome"],
            "host":"127.0.0.2",
            "port":8080
        }
    }

## List services

    POST localhost:8080/api

    {
        "module":"service",
        "method":"list",
        "body": "service-discovery"
    }

## Find a service (random lb in a pool)

    POST localhost:8080/api

    {
        "module":"service",
        "method":"find",
        "body": "service-discovery"
    }

## Remove a service from the pool

    POST localhost:8080/api

    {
        "module":"service",
        "method":"remove",
        "body": {
            "name":"service-discovery",
            "host":"127.0.0.2",
            "port":8080
        }
    }
