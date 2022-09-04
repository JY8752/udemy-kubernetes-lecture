# udemy サンプル

## MySQL

### namespace

```terminal
kubectl create ns database
or 
kubectl apply -f mysql-namespace.yml
```

### deployment

```terminal
(kubectl create deployment mysql --image=myswl:5.7 --dry-run=client -o yaml > mysql-deployment.yml)
環境変数の設定を追加する
kubectl apply -f mysql-deployment.yml -n database

kubectl exec -it -n database mysql-d8b464d5c-xhmnh -- mysql -uroot -p
```

### service

```terminal
kubectl create service clusterip mysql --tcp=3306 --dry-run=client -o yaml > mysql-service.yml
```

## sample app

### namespace

```terminal
namespace=yama
kubectl create ns $namespace
```

### deployment

```terminal
//deployment
kubectl create deployment sample-app --image=ghcr.io/nakamasato/github-actions-practice:pr-141 --dry-run=client -o yaml > sample-app-deployment.yml

//configMapRefとsecretRefの項目を追加する

//環境変数MYSQL_HOSTに設定する値は同じnamsespaceであればmysqlと記載すれば良いが、今回は別のnamespaceからの参照なので
//mysql.database.svc.cluster.localとなる

//configMap
kubectl create cm sample-app --from-env-file=env.txt --dry-run=client -o yaml > sample-app-configmap.yml

//secret
kubectl create secret generic sample-app --from-literal=MYSQL_PASSWORD=password --dry-run=client -o yaml > sample-app-secret.yml

//apply
kubectl apply -f sample-app-configmap.yml,sample-app-secret.yml,sample-app-deployment.yml -n $namespace
```

### service

```terminal
kubectl create service clusterip sample-app --tcp=80 --dry-run=client -o yaml > sample-app-service.yml
```

## request

```
kuebectl port-forward svc/sample-app 8080:80 -n $namespace

curl localhost:8080
```