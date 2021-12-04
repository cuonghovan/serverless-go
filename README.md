# Serverless Go example

This is an example code for deploying Go lambda functions using serverless framework.

## Environment variables

Configure enviroment variables in .envrc as bellow

```
export RELEASE_STAGE=<release stage>
export TENANT=<aws account id>
export AWS_REGION=<aws region>
```

# Deployment

## Build deployment packages

```
make
```

## Deploy

```
sls deploy -v
```
