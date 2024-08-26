# Simple Ecommerce Transaction Example

Example project of send and receive payments

## Tech Stack
 - [Go 1.23](https://go.dev/doc/install)
 - [Gin](https://gin-gonic.com/)
 - [Postgres](https://www.postgresql.org)
 - [Docker](https://www.docker.com)
   - [Docker Compose](https://docs.docker.com/compose/)



## Installation

Follow these steps to install the project:

1. Clone the repository:
    ```sh
    git clone https://github.com/lfcifuentes/online-payment-platform
    ```
2. Navigate to the project directory:
    ```sh
    cd online-payment-platform
    ```
3. Navidate to Bank project directory:
    ```sh
    cd bank_simulator
    ```
4. Create .env file
5. Install the dependencies:
    ```sh
    go mod tidy
    ```
6. Navidate to Api project directory:
    ```sh
    cd api
    ```
7. Create .env file
8. Install the dependencies:
    ```sh
    go mod tidy
    ```
9. Navigate to home and Run docker-compose
    ```sh
    make docker_up
    ```


# Makefile
Here is a list of the available commands in the Makefile and their description:

Starts the Docker containers in the background.
```sh
make docker_up
```

Stops and removes the Docker containers.
```sh
make docker_down
```
