# Asynq Monitoring
![alt text](https://github.com/krobus00/asynqmon-multi-service/blob/master/assets/list_services.png?raw=true)


## Setup
Copy config.yml.example to config.yml
```bash
cp config.yml.example config.yml
```
List your Asynq service name and redis dsn
```bash
services:
  - name: 'product-service-worker'
    redis_host: 'redis://localhost/1'
  - name: 'storage-service-server'
    redis_host: 'redis://localhost/2'
```

## Build your binary file
```bash
make run build
```

## Start server
for dev env you can use
```bash
make run dev server
```

for production env you can use
```bash
make run server
```

## Run in docker
Start docker container with this command
```bash
docker run -d -v "$(pwd)"/config.yml:/config.yml:ro -p 5000:5000 --name asynqmon-multi-service krobus00/asynqmon-multi-service:v0.1.0
```

Stop docker container with this command
```bash
docker stop asynqmon-multi-service 
docker rm asynqmon-multi-service
```