package main

import (
  "log"
  "net/http"
  "github.com/robertkrimen/otto"
  "io"
)

type Route struct {
  route_name string
  route_func otto.Value
  w http.ResponseWriter
  r *http.Request
}

func (ro Route) route(w http.ResponseWriter, r *http.Request) {
  ro.w = w
  ro.r = r
  g_vm.Set("writeMessage", ro.writeMessage)
  ro.route_func.Call(ro.route_func)
}

func (ro Route) writeMessage(call otto.FunctionCall) otto.Value {
  r, e := call.Argument(0).ToString()

  if e != nil {
    log.Fatal(e)
  } else {
    if r != "undefined" {
      io.WriteString(ro.w, r)
    } else {
      io.WriteString(ro.w, "")
    }
  }
  return otto.Value{}
}

func startServer(call otto.FunctionCall) otto.Value {
  r, e := call.Argument(0).ToString()

  if e != nil {
    panic(e)
  } else {
    if r != "" {
      log.Println("Server is listening on: " + r)
      http.ListenAndServe(r, nil)
    } else {
      panic("No port specified!")
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
      panic(e)
    }

    routing := Route{route_name: route_name, route_func: obj}

    http.HandleFunc(route_name, routing.route)

  } else {
    panic("Parameter provided for setRoute is not a function!")
  }

  return otto.Value{}
}
