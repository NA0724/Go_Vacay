package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key string
	val string
}

var theTests = []struct {
	name               string     //name of the test
	url                string     //url of the test
	method             string     //method of the test
	params             []postData //params of the test
	expectedStatusCode int        //status code of the test
}{
	{
		name:               "home",
		url:                "/",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "about",
		url:                "/about",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "contact",
		url:                "/contact",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "register",
		url:                "/register",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "login",
		url:                "/login",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "profile",
		url:                "/profile",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "search-availability",
		url:                "/search-availability",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "make-reservation",
		url:                "/make-reservation",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "executive-room",
		url:                "/executive-room",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "deluxe-room",
		url:                "/deluxe-room",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:               "premier-room",
		url:                "/premier-room",
		method:             "GET",
		params:             []postData{},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "postlogin",
		url:    "/login",
		method: "POST",
		params: []postData{
			{
				key: "email",
				val: "admin@admin.com",
			},
			{
				key: "password",
				val: "admin123",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "postRegister",
		url:    "/register",
		method: "POST",
		params: []postData{
			{
				key: "firstname",
				val: "neha",
			},
			{
				key: "lastname",
				val: "raj",
			},
			{
				key: "email",
				val: "raj@raj.com",
			},
			{
				key: "password",
				val: "raj123",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "post-search-availability",
		url:    "/search-availability",
		method: "POST",
		params: []postData{
			{
				key: "startdate",
				val: "2023-06-02",
			},
			{
				key: "enddate",
				val: "2023-06-04",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "post-search-availability-json",
		url:    "/search-availability-json",
		method: "POST",
		params: []postData{
			{
				key: "startdate",
				val: "2023-06-02",
			},
			{
				key: "enddate",
				val: "2023-06-04",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name:   "make-reservation",
		url:    "/make-reservation",
		method: "POST",
		params: []postData{
			{
				key: "firstname",
				val: "neha",
			},
			{
				key: "lastname",
				val: "raj",
			},
			{
				key: "email",
				val: "raj@raj.com",
			},
			{
				key: "phone",
				val: "5555555555",
			},
			{
				key: "startdate",
				val: "2023-06-02",
			},
			{
				key: "enddate",
				val: "2023-06-04",
			},
		},
		expectedStatusCode: http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected status code %d, got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, param := range test.params {
				values.Add(param.key, param.val)
			}
			response, err := testServer.Client().PostForm(testServer.URL+test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s, expected status code %d, got %d", test.name, test.expectedStatusCode, response.StatusCode)
			}
		}
	}
}
