
# Go API Example

This project is a simple example of a Go API using the [Gorilla Mux](https://github.com/gorilla/mux) router package.

## Requirements

- Go 1.16 or higher
- Gorilla Mux: This package is used for routing and can be installed using the following command:

  ``
  go get -u github.com/gorilla/mux
``

### Project Structure

```
goapiexample/
│
├── api/
│   └── example/
│       └── helloworld.go  # Example handler
│   └── router.go          # Router setup
│
├── .gitignore             # Ignored files
├── go.mod                 # Go module file
├── go.sum                 # Dependencies
├── LICENSE                # License information
├── main.go                # Main entry point of the application
└── README.md              # This file
```

