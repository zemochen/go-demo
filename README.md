# go-demo
This is a online mall demo build by golang. I build this for learn golang 


# Env 

* docker

## Common bash most used



### Makefile

Use makefile generate the project file,
eg:
generate `app/frontend`

```bash
make gen-frontend
```

## Service Discovery

### Consul

**Docker compose**
see: [docker-compose.yaml](docker-compose.yaml)

```bash
# start consul backend
docker-compose up -d
```

**Consul UI**
> The admin web side for manage service in consul
- http://localhost:8500/ui  
