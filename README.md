# Saloncito API Gateway Submodule
This submodule functions as a single entry point for the application. It is responsible for communicating with the other microservices and returning the response to the client.

### Table of contents
- [How to run the project?](#how-to-run-the-project)
  - [Prerequisites](#prerequisites)
  - [Project cloning](#project-cloning)
  - [Environment variables](#environment-variables)
  - [Docker execution](#docker-execution)
  - [Proto files](#proto-files)

## How to run the project?
### Prerequisites
- [Docker](https://docs.docker.com/install/)
- [Docker compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/?utm_source=chatgpt.com#download)
- [Protoc](https://grpc.io/docs/protoc-installation/)

*Note: Keep in mind that the entire project was tested on WSL.*

### Project cloning
If you haven't already cloned the project, you can do so by running `git clone <url>`.
Each of the following steps can be executed by running `npm run setup` (this command should be executed only the first time).

### Environment variables
To configure the environment variables within the project, there is a command in the Makefile file called `make create-envs`. When executed, it will copy the ".env.example" file to a "app.env" file. If you need to change something, just modify the newly created "app.env" file.

### Docker execution
After configuring the environment variables, it's time to run the Docker containers. The necessary files are located in the ".dockers" folder. Run the makefile script `make compose`, which will execute the ".dockers/docker-compose.yml" file and start building the project. If necessary, there is also a command `make compose-build` that explicitly builds the project.

### Proto files
If you modified any proto file in the "sa_proto" submodule and need to generate the Go files, you can run `make proto`, and the Go files will be generated.