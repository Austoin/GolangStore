package config

import "testing"

func TestLoadUsesDefaultValues(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("APP_NAME", "")
	t.Setenv("MYSQL_HOST", "")
	t.Setenv("MYSQL_PORT", "")
	t.Setenv("MYSQL_USER", "")
	t.Setenv("MYSQL_PASSWORD", "")
	t.Setenv("MYSQL_DATABASE", "")
	t.Setenv("REDIS_HOST", "")
	t.Setenv("REDIS_PORT", "")
	t.Setenv("REDIS_PASSWORD", "")

	conf := Load()

	if conf.App.Env != "local" {
		t.Fatalf("expected default app env, got %s", conf.App.Env)
	}

	if conf.MySQL.Host != "127.0.0.1" {
		t.Fatalf("expected default mysql host, got %s", conf.MySQL.Host)
	}

	if conf.Redis.Port != "6379" {
		t.Fatalf("expected default redis port, got %s", conf.Redis.Port)
	}
}

func TestLoadUsesEnvironmentValues(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	t.Setenv("APP_NAME", "DemoStore")
	t.Setenv("MYSQL_HOST", "mysql")
	t.Setenv("MYSQL_PORT", "3307")
	t.Setenv("MYSQL_USER", "demo")
	t.Setenv("MYSQL_PASSWORD", "secret")
	t.Setenv("MYSQL_DATABASE", "demo_store")
	t.Setenv("REDIS_HOST", "redis")
	t.Setenv("REDIS_PORT", "6380")
	t.Setenv("REDIS_PASSWORD", "redis-secret")

	conf := Load()

	if conf.App.Name != "DemoStore" {
		t.Fatalf("expected app name DemoStore, got %s", conf.App.Name)
	}

	if conf.MySQL.Database != "demo_store" {
		t.Fatalf("expected mysql database demo_store, got %s", conf.MySQL.Database)
	}

	if conf.Redis.Password != "redis-secret" {
		t.Fatalf("expected redis password redis-secret, got %s", conf.Redis.Password)
	}
}
