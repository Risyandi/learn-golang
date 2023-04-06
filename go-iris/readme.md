# Golang with Iris Framework  
In this repository we will learning how to used the Iris framework, project structured in Iris framework and if you want to learn more about Iris framework can visit this sites https://www.iris-go.com/.

**Requirement :**

- Golang version minimum 1.20 or above
- Library Iris framework version 12.2.0  
    `$ go get github.com/kataras/iris/v12@latest`  
    learn more at [Iris official documentation](https://www.iris-go.com/docs)

**Troubleshooting :**

- Perform clean of your go modules cache  
    `$ go clean --modcache`
- Network error during installation Iris  
    `$ go env -w GOPROXY=https://goproxy.io,direct`

**Running :**  

- Running your Iris framework  
    `$ go run main.go --config=server.yml`