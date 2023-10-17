# CodePix
CodePix simulates transfers of different amounts of money via PIX, Brazil's Central Bank instant payments system.

## Table of Content
* [How to prepare the environment](#how-to-prepare-the-environment)
* [How to execute the application](#how-to-execute-the-application)

## How to prepare the environment
In order to get all of our services up and running we are going to use Docker. <br>
Fire up a terminal and let's get going.
* Clone the repo, i'm using SSH but you can use HTTP or the GitHub CLI:
```bash
git clone git@github.com:gabrielparaizo/CodePix.git
```
* With Docker installed on your machine, just run the following command:
```bash
docker compose up -d
```

## How to execute the application
* Access the application container:
```bash
docker exec -it codepix-app-1 bash
```