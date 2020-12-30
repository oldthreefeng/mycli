package apisix

import (
	"testing"
)

func Test_initRoutes(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initRoutes()
		})
	}
}

func TestUpdateSSL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateSSL()
		})
	}
}

func Test_getApisixAllssl(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getApisixAllssl()
		})
	}
}
