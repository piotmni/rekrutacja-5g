#### Docs

##### Dockerfiles

dockerfiles are stored in dockerfiles/


There are stored dockerfiles to build open5gs, open5gs-webui and ueransim

all of them are build and pushed to docker-hub registry piotmni/

##### Charts

List of charts:
```
- open5gs-5gcore
- open5gs-component-base
- open5gs-webui
- ueransim
```

open5gs-component-base is chart created to use as wrapper to start all required NFs

It is used as a dependency with specified aliases in open5gs-5gcore

Components:
```
- open5gs-amfd
- open5gs-ausfd
- open5gs-bsfd
- open5gs-nrfd
- open5gs-nssfd
- open5gs-pcfd
- open5gs-smfd
- open5gs-udmd
- open5gs-udrd
- open5gs-upfd
```

ueransim chart is only for testing purposes

How to run this is stored in [docs/README.md](./docs/README.md)
