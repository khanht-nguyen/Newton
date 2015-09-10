package main

import (
  "net/http"
  "log"
  "io"

  "github.com/robertkrimen/otto"
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
  g_vm.Set("setHeader", ro.setHeader)

  ro.route_func.Call(ro.route_func)
}

func (ro Route) setHeader(call otto.FunctionCall) otto.Value {
  k, e := call.Argument(0).ToString()
  v, e2 := call.Argument(1).ToString()

  if e != nil {
    log.Fatal(e)
  } else if e2 != nil {
    log.Fatal(e2)
  } else {
    if k != "undefined" && v != "undefined" {
      ro.w.Header().Set(k, v)
    } else {
      io.WriteString(ro.w, "")
    }
  }
  return otto.Value{}
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
