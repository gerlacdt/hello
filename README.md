## Summary

Simple golang server app with following endpoints:

* /hello
* /version
* /health

Used for testing purposes for:

* building a golang app
* create docker image
* publish docker image to dockerhub
* run docker image in different docker container management environments (nomad, kubernetes)


### Kubernetes instructions

This app comes with simple Kubernetes manifest. You can install the hello app with:

``` bash
# added a simple secret for testing purposes. It will be printed during startup of the app
kubectl apply -f k8s/secret.yaml

# install the pod with the hello container
kubectl apply -f k8s/deployment.yaml

# make the pods available in the cluster
kubectl apply -f k8s/service.yaml
```

#### Helm

For testing purposes the hello app comes with a simple helm
chart. Helm charts are more flexible than static Kubernetes yaml
manifests. You benefit from the helm template language which supports
variables and helm values. With these you can configure you helm
release (Kubernetes deployment). E.g. you can change the name of the
deployment or change the number of replicas without to touch the yaml
helm templates.

``` bash
# install helm chart
helm install --name myhello hello

# delete helm chart
helm delete --purge myhello
```
