# Golang with Iris Framework  
in this repository we will learning how to used the Iris framework. learn more about Iris framework can visit this sites https://www.iris-go.com/.

**Requirement :**

- Golang version 1.20 or above
- Library Iris framework version 12.2.0
    `$ go get github.com/kataras/iris/v12@latest`  
    learn more at [Iris official documentation](https://www.iris-go.com/docs)

**Troubleshooting :**

- Perform clean of your go modules cache  
    `$ go clean --modcache`
- Network error during installation Iris  
    `$ go env -w GOPROXY=https://goproxy.io,direct`