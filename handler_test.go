package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpRequestMethods = []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace}

func httpReqMethodAllowed(httpReqMethodDefinedInRouter, testHttReqMethod string) bool {
	if httpReqMethodDefinedInRouter == testHttReqMethod {
		return true
	} else if httpReqMethodDefinedInRouter == http.MethodGet && testHttReqMethod == http.MethodHead {
		return true
	}
	return false
}

var okFn = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type dispatchMethod struct {
	method   string
	fn       http.HandlerFunc
	testName string
}

var givenTestDispatchMethods = []dispatchMethod{
	{
		method:   http.MethodGet,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Get",
	},
	{
		method:   http.MethodPost,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Post",
	},
	{
		method:   http.MethodPut,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Put",
	},
	{
		method:   http.MethodPatch,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Patch",
	},
	{
		method:   http.MethodDelete,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Delete",
	},
	{
		method:   http.MethodConnect,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Connect",
	},
	{
		method:   http.MethodOptions,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Options",
	},
	{
		method:   http.MethodTrace,
		fn:       okFn,
		testName: "Method_Defined_In_Router_Trace",
	},
}

func prepareDispatchMethodMap(d dispatchMethod) map[string]http.HandlerFunc {
	if d.method == http.MethodHead {
		d.method = http.MethodGet
	}
	return map[string]http.HandlerFunc{d.method: d.fn}
}
func TestDispatchMethods(t *testing.T) {
	for _, d := range givenTestDispatchMethods {
		for _, testHTTPReqMethod := range httpRequestMethods {
			t.Run(d.testName+"_Test_Method_"+testHTTPReqMethod, func(t *testing.T) {
				r := &http.Request{Method: testHTTPReqMethod}
				w := httptest.NewRecorder()
				DispatchMethods(prepareDispatchMethodMap(d))(w, r)
				if !httpReqMethodAllowed(d.method, testHTTPReqMethod) {
					if http.StatusMethodNotAllowed != w.Result().StatusCode {
						t.Errorf("expected status code[%d], actual status code[%d]", http.StatusMethodNotAllowed, w.Result().StatusCode)
					}
				} else if http.StatusOK != w.Result().StatusCode {
					t.Errorf("expected status code[%d], actual status code[%d]", http.StatusOK, w.Result().StatusCode)
				}
			})
		}
	}
}
