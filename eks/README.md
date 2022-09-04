# EKS demo

## install

```terminal
brew tap weaveworks/tap
brew install weaveworks/tap/eksctl
eksctl version

> 0.110.0
```

## create cluster

```terminal
eksctl create cluster --name test-cluster --region ap-northeast-1
```

## update kubeconfig

```terminal
aws eks update-kubeconfig --name test-cluster
less ~/.kube/config
```

## deploy

```terminal
//nodeの確認
kubectl get node

//deploymentを作成
kubectl create deploy nginx --image=nginx --replicas=3
kubectl get pod

//ロードバランサーでServiceを作成
kubectl create service loadbalancer nginx --tcp=80
```

## リソースの削除

```terminal
kubectl delete svc nginx
kubectl delete deploy nginx

//eksクラスターの削除
eksctl delete cluster --name test-cluster --region ap-northeast-1

//kubeconfigも戻す
vim ~/.kube/config

currentConfig -> docker-desktop
```