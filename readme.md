## ginerr: A Go package for standardized error handling in Gin

The `ginerr` package provides a simple and effective way to manage errors in your Gin applications. It defines a consistent error structure and includes middleware to automatically handle and return errors to clients in a standardized format.

### Features

- **Standardized Error Structure:**  Defines a `HighLevelError` type that wraps HTTP status codes and error messages.
- **Common Error Definitions:**  Provides predefined instances of `HighLevelError` for common HTTP status codes.
- **Gin Middleware:** Includes a `GinErrorHandlerMiddleware` to automatically handle and format errors thrown in Gin routes.
- **Convenience Functions:** Provides a simple `AbortAndError` function to gracefully handle errors and abort requests.
- **Error Extraction:**  The `ExtractError` function can be used to extract `HighLevelError` from any error.

### Installation

Use the following command to install the package:

```bash
go get github.com/vloldik/ginerr
```

### Usage

1. **Import the Package:**

   ```go
   import (
       "github.com/vloldik/ginerr"
       "github.com/gin-gonic/gin"
   )
   ```

2. **Create and Use Errors:**

   ```go
   func handleRequest(c *gin.Context) {
       // ... your logic
       if err := someError(); err != nil {
           ginerr.AbortAndError(c, ginerr.NewHighLevelError(http.StatusInternalServerError, "Something went wrong"))
           return
       }
       // ...
   }
   ```

3. **Use Predefined Errors:**

   ```go
   func handleRequest(c *gin.Context) {
       // ... your logic
       if err := someError(); err != nil {
           ginerr.AbortAndError(c, ginerr.InternalServerError)
           return
       }
       // ...
   }
   ```

4. **Use Error Extraction for Custom Errors:**

   ```go
   func handleRequest(c *gin.Context) {
       // ... your logic
       if err := someError(); err != nil {
           highLevelErr := ginerr.ExtractError(err)
           ginerr.AbortAndError(c, highLevelErr)
           return
       }
       // ...
   }
   ```

5. **Use Middleware:**

   ```go
   func main() {
       r := gin.Default()
       // ... your route setup
       // Register the middleware
       r.Use(ginerr.GinErrorHandlerMiddleware)
       // ... your route setup
       r.Run()
   }
   ```

### Example

```go
package main

import (
    "github.com/vloldik/ginerr"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(ginerr.GinErrorHandlerMiddleware)

    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        // ... logic to fetch user by ID

        if id == "invalid" {
            ginerr.AbortAndError(c, ginerr.NotFound)
            return
        }
        // ... Handle success
    })

    r.Run()
}
```

### Documentation

**HighLevelError**

```go
type HighLevelError struct {
    Code    int
    Message string
}
```

Represents a standardized error with an HTTP status code and an error message.

**Predefined Errors:**

- **DefaultInternalError:** `http.StatusInternalServerError`, "internal server error"
- **BadRequest:** `http.StatusBadRequest`, "missed parameter, incorrect or malformed data"
- **NotFound:** `http.StatusNotFound`, "not found"
- **Unauthorized:** `http.StatusUnauthorized`, "unauthorized"
- **Forbidden:** `http.StatusForbidden`, "forbidden"
- **Conflict:** `http.StatusConflict`, "conflict"
- **TooManyRequests:** `http.StatusTooManyRequests`, "too many requests"
- **InternalServerError:** `http.StatusInternalServerError`, "internal server error"
- **NotImplemented:** `http.StatusNotImplemented`, "not implemented"
- **ServiceUnavailable:** `http.StatusServiceUnavailable`, "service unavailable"
- **GatewayTimeout:** `http.StatusGatewayTimeout`, "gateway timeout"
- **BadGateway:** `http.StatusBadGateway`, "bad gateway"
- **ProxyError:** `http.StatusBadGateway`, "proxy error"
- **UnknownError:** `http.StatusInternalServerError`, "unknown error"

**ExtractError**:

```go
func ExtractError(err error) HighLevelError
```

Extracts a `HighLevelError` from any error, returning `DefaultInternalError` if no `HighLevelError` is found.

**GinErrorHandlerMiddleware:**

```go
func GinErrorHandlerMiddleware(context *gin.Context)
```

Middleware that automatically handles errors thrown in Gin routes. It extracts the `HighLevelError` and returns a JSON response to the client.

**AbortAndError:**

```go
func AbortAndError(c *gin.Context, err error)
```

Aborts the request and sets the error to be handled by the `GinErrorHandlerMiddleware`.

### Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.

### License

The ginerr package is released under the MIT license.