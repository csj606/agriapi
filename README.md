# AgriAPI - Ongoing Project

[![Network Tests](https://github.com/csj606/agriapi/actions/workflows/network_tests.yml/badge.svg)](https://github.com/csj606/agriapi/actions/workflows/network_tests.yml)

### Overview

AgriAPI is an RESTful API built using Golang Gin to provide users when queried with a list of in-season crops or current planting windows. The API offers two methods of deployment.
You could choose to locally deploy the service using a Docker container, or deploy the API to AWS using a Terraform script. Currently, the API only provides data for the state of Iowa.

### Endpoints

/ping - mainly exists as a testing endpoint for development purposes
/in-season - returns a list of crops which are in season at this specific time
/planting - returns a list of crops which are currently in their planting windows

### Stack

- Router - Written in Golang
- IaC + Cloud Service - AWS and Terraform
- Testing Libraries - testify

The application is also containerized using Docker and a CI pipeline has been made in GitHub Actions.

### Important Notes

The Terraform script is written assuming a private SSH key exists in your cloud environment named cloud-private-key and also exists in your saved SSH keys. The Terraform script exists mostly for experimental purposes. It uses a provisioner which typically are not recommended for usage in a production environment.

### Roadmap

The MVP version for this project has been completed. I may come back and implement the following:
1. A more complex API with data from more states, such as Kansas and Nebraska
2. Add to the Terraform script to provide VPC + security rules and other AWS components.



