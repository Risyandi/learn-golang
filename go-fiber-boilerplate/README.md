# Boilerplate Golang With Fiber

A Golang boilerplate project using the [Fiber](https://https://gofiber.io/) web framework. This project is set up with best practices for a scalable and maintainable web application.

## Description

This repository provides a starting point for building web applications in Golang using the [Fiber](https://https://gofiber.io/) framework. It includes a structured project layout, basic middleware, and configuration for environment variables.

## Installation

1. Clone the repository
   - `$ git clone https://gitlab.com/risyandi88/boilerplate-go`
   - `$ cd boilerplate-go`
2. Install dependency you needed, Make sure you have Go installed (version 1.20+ is recommended).
   - `$ go get` or `$ go install`
3. Setup environment variable in file `.env` with copying from example.
   - `$ cp .env.example .env`

## Usage

**1. Development**  
For hot-reload during development, you can use either Gin or Air.

- Gin
  - install: `$ go install github.com/codegangsta/gin@latest`
  - run the development server: `$ gin -i run main.go`
- Air
  - install: `$ go install github.com/cosmtrek/air@latest`
  - run the development serve: `$ air`

**2. Building and Running**  
For build the application you can follow this step

- build the application: `$ go build -o app`
- run the application: `$ ./app`

## Contributing

For people who want to make changes to our project, it's helpful to have some documentation on how to get started. Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/ your-feature-name`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License.
