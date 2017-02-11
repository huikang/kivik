package test

import (
	"testing"

	"github.com/flimzy/kivik"
)

func TestDBExists(t *testing.T) {
	s, err := kivik.New(TestDriver, TestServer)
	if err != nil {
		t.Fatalf("Error connecting to %s: %s\n", TestServer, err)
	}
	exists, err := s.DBExists("bogusDB")
	if err != nil {
		t.Fatalf("Failed to check DB existence: %s", err)
	}
	if !exists {
		t.Errorf("DBExists() should always return true on local databases")
	}
	// FIXME: Uncomment this once it's possible to create a database
	// exists, err = s.DBExists("_users")
	// if err != nil {
	// 	t.Fatalf("Failed to check DB existence: %s", err)
	// }
	// if !exists {
	// 	t.Errorf("DB '_users' should exist")
	// }
}

func TestRemoteDBExist(t *testing.T) {
	s, err := kivik.New(TestDriver, RemoteServer)
	if err != nil {
		t.Fatalf("Error connecting to %s: %s\n", RemoteServer, err)
	}
	exists, err := s.DBExists(RemoteServer + "/bogusDB")
	if err != nil {
		t.Fatalf("Failed to check DB existence: %s", err)
	}
	if exists {
		t.Errorf("DB 'bogusDB' should not exist")
	}

}