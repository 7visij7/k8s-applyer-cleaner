# K8s-applyer-cleaner
> The result of application work is reduce to 0 replicas for deployment with "applyer" in name, which run more that one hour.
---
## Build and run application
```Bash
go build -o k8s-applyer-cleaner
./k8s-applyer-cleaner
```
---

## Required variables and config
> Required enviroment variables: 
+ OPENSHIFT_SERVER - url to API K8S
+ ENCRYPT_KEY - secret key AES. 

> Also application works with config.yaml, where stored tokens to acces to namespaces. 
e.g.:
```Bash
Namespace1: Token1
Namespace2: Token2
```
---
## Docker

> Build Docker image from a [Dockerfile](https://github.com/7visij7/k8s-applyer-cleaner/blob/main/Dockerfile)
```
docker build -t IMAGENAME
```
> Start application
```
docker run -it --rm IMAGENAME
```
---
