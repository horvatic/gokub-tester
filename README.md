# gokub-tester
A project to be deployed, and test kubernetes deployment. 

## Using
Run the following command to deploy
```
kubectl apply -k deploy/dev
```

or replica set run 

```
kubectl apply -k deploy/prod
```
The ${TAG} will need to be replaced with a gokub image tag in deploy/base/deployment.yaml before an apply can be used 
