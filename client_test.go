package go_http_client

import (
	"testing"
)

func build(s string) string { return "https://httpbin.org" + s }

func TestNew(t *testing.T) {
	_, err := New(build("/get"))
	if err != nil {
		t.Error(err)
	}
}

func TestNewMethod(t *testing.T) {
	client, err := NewMethod(build("/post"), "POST")
	if err != nil {
		t.Error(err)
	}
	if client.Request.Method != "POST" {
		t.Error("Method is not POST!")
	}
}

func TestClient_Do(t *testing.T) {
	client, err := New(build("/get"))
	if err != nil {
		t.Error(err)
	}
	cl, err := client.Do()
	if err != nil {
		t.Error(err)
	}
	if cl.StatusCode != 200 {
		t.Error(cl.Status)
	}
}

func TestClient_SetBasicAuth(t *testing.T) {
	client, err := New(build("/basic-auth/foo/bar"))
	if err != nil {
		t.Error(err)
	}
	client.SetBasicAuth("foo", "bar")
	cl, err := client.Do()
	if err != nil {
		t.Error(err)
	}
	if cl.StatusCode != 200 {
		t.Error(cl.Status)
	}
}
