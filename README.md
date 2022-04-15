# kubedb-api-demo

1. Create a Kind cluster

```bash
> kind create cluster
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.23.4) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Have a nice day! 👋
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
