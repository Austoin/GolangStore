package redis

import (
	"testing"

	"github.com/austoin/GolangStore/pkg/config"
)

func TestBuildAddr(t *testing.T) {
	conf := config.Redis{
		Host:     "redis",
		Port:     "6379",
		Password: "secret",
	}

	addr := BuildAddr(conf)
	if addr != "redis:6379" {
		t.Fatalf("expected redis:6379, got %s", addr)
	}
}

func TestNewClientUsesConfig(t *testing.T) {
	conf := config.Redis{
		Host:     "redis",
		Port:     "6379",
		Password: "secret",
	}

	client := NewClient(conf)
	options := client.Options()

	if options.Addr != "redis:6379" {
		t.Fatalf("expected addr redis:6379, got %s", options.Addr)
	}

	if options.Password != "secret" {
		t.Fatalf("expected password secret, got %s", options.Password)
	}
}
