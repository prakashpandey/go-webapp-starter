package main

import (
	"net/http"
	"testing"
)

const baseUrl = "http://localhost"
const port = 8284

var testRolesData = []struct {
	api    string
	method string
	role   []string
}{
	{
		api:    "/user",
		method: http.MethodPost,
		role:   []string{},
	},
	{
		api:    "/user",
		method: http.MethodGet,
		role:   []string{"this"},
	},
	{
		api:    "/admin",
		method: http.MethodDelete,
		role:   []string{"super_admin", "admin"},
	},
}

func Test_userRoles(t *testing.T) {
	for _, d := range testRolesData {
		switch d.method {
		case http.MethodPost:
			http.Post()
		case http.MethodGet:
			http.Get(baseUrl + d.api + ":" + string(port))
		case http.MethodDelete:

		}
	}
}
