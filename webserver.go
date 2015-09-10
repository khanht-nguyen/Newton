package main

import (
  "log"
  "net/http"

  "github.com/robertkrimen/otto"
)

func serveFiles(call otto.FunctionCall) otto.Value {
  http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
  })
  return otto.Value{}
}

func startServer(call otto.FunctionCall) otto.Value {
  r, e := call.Argument(0).ToString()

  if e != nil {
    log.Panic(e)
  } else {
    if r != "" {
      log.Println("Server is listening on: " + r)
      http.ListenAndServe(r, nil)
    } else {
      log.Panic("No port specified!")
    }
  }

  return otto.Value{}
}

func setRoute(call otto.FunctionCall) otto.Value {
  r := call.Argument(0)
  obj := call.Argument(1)

  if obj.IsFunction() {
    route_name, e := r.ToString()
    if e != nil {
      log.Panic(e)
    }

    routing := Route{route_name: route_name, route_func: obj}

    http.HandleFunc(route_name, routing.route)

  } else {
    panic("Parameter provided for setRoute is not a function!")
  }

  return otto.Value{}
}
