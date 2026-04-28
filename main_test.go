package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkAPIParallel(b *testing.B) {
	req := httptest.NewRequest("GET", "/user", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handler)
	// ajuste le niveau de parallélisme
	//b.SetParallelism(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
		}
	})
}
