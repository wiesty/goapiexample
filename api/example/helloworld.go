package example

import (
    "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Go!"))
}
