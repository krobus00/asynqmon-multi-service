# Asynq Monitoring

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
