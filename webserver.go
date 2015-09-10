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
}

func (ro Route) route(w http.ResponseWriter, r *http.Request) {
  v, e := ro.route_func.Call(ro.route_func, nil)

  if e != nil {
    log.Fatal(e)
  } else {
    s, e := v.ToString()

    if e != nil {
      log.Fatal(e)
    } else {
      io.WriteString(w, s)
    }
  }
}

func writeMessage(call otto.FunctionCall) otto.Value {
  obj, e := call.Argument(0).Export()
  r, e := call.Argument(1).ToString()

  if e != nil {
      log.Fatal("writeMessage failed!")
    } else {
      f_obj := obj.(http.ResponseWriter)

      io.WriteString(f_obj, r)
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
