##### Dockerfiles

Dockerfiles used to build open5gs, open5gs-webui and ueransim are stored in [dockerfiles](./dockerfiles)

All of them are build and pushed to docker-hub registry piotmni/

##### Charts

List of charts:
```
- open5gs-5gcore
- open5gs-component-base
- open5gs-webui
- ueransim
```

open5gs-component-base is chart created to use as wrapper to start all required NFs.

open5gs-5gcore is chart that mainly has a dependecies to start NFs and mongodb. Also allows to start debugging container.

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

Example with steps how to run this is stored in [docs/README.md](./docs/README.md)
