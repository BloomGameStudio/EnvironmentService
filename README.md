# Environment Service

The Bloom & LostLight Environment Service

## Quickstart Guide:

## Installation

### **Native Bare Metel**

        // Install Deps/Modules
        go mod download

        // Run the server on port 1323
        go run .

**Optional Live Code Reloading with Air**

Install Air via your preffered installation method: https://github.com/cosmtrek/air

        // Run the server on port 1323 with live code reloading
        air

### **Using Docker && DockerCompose**

---

### **Docker-Compose**

Chose your docker compose cli
Depending on what version you have or how you installed docker compose.

The examples will use the more wider used `docker-compose`

For more Information read: https://stackoverflow.com/questions/66514436/difference-between-docker-compose-and-docker-compose

The more wider used `docker-compose`.

        docker-compose <command>

The newer `docker compose`.

        docker compose <command>

### **Start the App and listen on port 1323**

Note: Depending on your system and context you may have to configure your image & container versions
view the official Docker Compose documentation on how Docker determines what and how it runs images & containers and how <docker-compose up> behaves.

TLDR: Docker and by extension Docker Compose will chose the latest container and if that does not exist the latest image to run your application.

**Run the latest version that was build from the branch main.**

Note: If you have build a later version or somehow else have a later version on your system a version that docker thinks is later than what was build from main it will most likely use that.
Which resulst in you not runing the version from main and not runing the intended version.
This will automatically be resolved for you if a new push to main happens.

        docker-compose up
        // CTRL + C to stop

**Run & Build the current state of the currently checkout out branch.**

Note: This will build a image and run and build a container which probably is a later version than the prebuild image built from the main branch.

        docker-compose up --build
        // CTRL + C to stop

        // If there are caching issues or some other problems or you want to be 100% sure that you run and have build the latest version of the current branch you can run:
        docker-compose up --build --force-recreate
        // This will recreate everything and might take longer.

---

### **Docker**

1.  Create a The docker volume for the database

        docker volume create enviormentservicevolume

2.  Run The Container

        // From Github Container Registry via Image
        // You can Replace the tag <main> at the end with whatever tag you want

                docker run --rm -p 1323:1323 -v enviormentservicevolume:/database ghcr.io/bloomgamestudio/environmentservice:main

        // Build it yourself locally with build tag/name then run it

                docker build -t environmentservice .

                docker run -p 1323:1323 -v enviormentservicevolume:/database environmentservice


        // Build it yourself locally without tag/name then run it

                docker build .

                docker run -p 1323:1323 -v enviormentservicevolume:/database <Containername>

---
