# Go Start


1. create .env from .env.dist
2. make docker-build
3. make docker-up
4. You are ready!

Make sure that the ports in the .env are not being used. Otherwise, docker will throw an error.

## Github Actions
Fork this project and add a self runner to the server. Deploy ready

## Makefile

make {command}

### docker

- docker-build
- docker-up
- docker-restart
- docker-exec

### swagger

- swagger (update docs)

### test

- test
  
## Swagger

```
http://{{domain}}/api/swagger
```

## Metrics

Basic Auth : system username and password from .env  

```
http://{{domain}}/api/system/metrics
```