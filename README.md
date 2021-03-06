# kubedb-api-demo

1. Create a Kind cluster

```bash
> kind create cluster
Creating cluster "kind" ...
 â Ensuring node image (kindest/node:v1.23.4) đŧ
 â Preparing nodes đĻ
 â Writing configuration đ
 â Starting control-plane đšī¸
 â Installing CNI đ
 â Installing StorageClass đž
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Have a nice day! đ
```

2. Register Postgres CRD

```
> kubectl create -f https://raw.githubusercontent.com/kubedb/apimachinery/v0.26.0/crds/kubedb.com_postgreses.yaml
customresourcedefinition.apiextensions.k8s.io/postgreses.kubedb.com created
```

3. Create Postgres object

```
> k apply -f pg.yaml
postgres.kubedb.com/quick-postgres created
```

4. List objects

```
> go run main.go

Using Generated client
default/quick-postgres
Using kubebuilder client
default/quick-postgres
```
