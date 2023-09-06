# velero-ui
Web UI for velero backups

Lot's of work in progress here.

## Install
```bash
helm repo add velero-ui https://hsmade.github.io/velero-ui/
kubectl create ns velero-ui
helm install velero-ui velero-ui/velero-ui --namespace velero-ui
```
