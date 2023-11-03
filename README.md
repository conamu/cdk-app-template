!! NOTE !!: Localstack integration will be removed in favor of faster development capability :)

This repository on GitHub is merely a mirror of my GitLab instance, so PRs will not be checked regularly.

## Todo:
- Lambdas have their own constructs, underlying infra has its own stack. This will contribute to faster deployments

# Welcome to the template for all things serverless!
This cdk-app template for golang serverless applications on aws is built for fast iteration and fast time to market

## Main Concepts

### Implicit routes by lambda name
The directory internal/app/lambda is the place where you put your lambda code.
The api routes and methods are built based on the lambda directories names.

So this structure in the template:
```
internal/
    app/
        lambda/
            ping-get/
            ping-post/
            pong-get/
```

will result in following api paths and methods:
```
GET /ping
POST /ping
GET /pong
```

### Automatic test environments
Based on the Branch of the MR a testing stack is built and deployed.
Through a manual-triggerable step in the CI of a MR the test stack can be teared down.

The link to the test env is posted into the MR to facilitate fast testing without visiting the AWS Dashboard.
API Key Authentication is also disabled for this reason on test envs.

### Fully local development Ready
Through the development of this template with localstack in mind you can build almost
everything locally before even deploying it once. Depending on what localstack supports of course.

You should still do final testing on AWS since localstack does not implement everything exactly like on AWS.

### Easy direct deployment
Just use the quick exports of AWS IAM Identity Center to set credentials for the account you want to use.
By default, this tempalte deploys every environment to the same account.

## Get Started
Fork this repository.

Set environment variables in your repo:
```
AWS_ACCESS_KEY_ID -> Access key id of the deployer user in your account
AWS_SECRET_ACCESS_KEY -> Secret Access Key of the deployer account
PROJECT_TOKEN -> GitLab Project Access token with api scope

To run automatic Acceptance tests through postman, get these details out of your postman client:
POSTMAN_ENV_ID 
POSTMAN_COLLECTION_ID
POSTMAN_API_KEY
```

### Local development deployment
1. `make localstack up` to start the localstack container. Input api key in compose file to use pro features (recommended)
2. `make local-deploy` to deploy your app to localstack.
3. `make localstack down` to stop localstack or to clear it in case something went wrong.

### Deployment to AWS Account
1. `make bootstrap` to set up tooling for the aws cdk. Do once per region and account.
2. `make deploy-staging` to deploy a staging environment.
3. `make destroy` to teardown the deployment.

By using a custom environment, you can deploy multiple staging/test environments (CI does this automatically):
```ENV=testing make deploy-staging```
