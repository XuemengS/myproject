docker build -t xuemengs/stsb:v6 .
docker push xuemengs/stsb:v6
kubectl set image deployment/hello-app hello-app=xuemengs/stsb:v6

