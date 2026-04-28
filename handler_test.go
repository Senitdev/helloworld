package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkAPI(b *testing.B) {
	req := httptest.NewRequest("GET", "/user", nil)
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		handler(rr, req)
	}
}
