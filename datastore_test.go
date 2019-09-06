package main

import (
	"context"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestDatastoreClient_CRUDUser(t *testing.T) {
	c, err := NewDatastoreClient()
	if err != nil {
		t.Fatalf("failed to create client. %v", err)
	}

	q := datastore.NewQuery("User").KeysOnly()
	keys, err := c.client.GetAll(context.Background(), q, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, k := range keys {
		if err := c.client.Delete(context.Background(), k); err != nil {
			t.Fatal(err)
		}
	}

	id, err := c.AddUser(&User{Name: "user1", Age: 10})
	if err != nil {
		t.Fatalf("failed to add user. %v", err)
	}

	id, err = c.AddUser(&User{Name: "user2", Age: 20})
	if err != nil {
		t.Fatalf("failed to add user. %v", err)
	}

	u, err := c.GetUser(id)
	if err != nil {
		t.Fatalf("failed to get user. %v", err)
	}

	if u.Name != "user2" || u.Age != 20 {
		t.Errorf("invalid user. %v", u)
	}

	users, err := c.ListUsers()
	if err != nil {
		t.Fatalf("failed to list users. %v", err)
	}
	if len(users) != 2 {
		t.Errorf("invalid user list. %v", users)
	}

	err = c.DeleteUser(u.ID)
	if err != nil {
		t.Fatalf("failed to delete user. %v", err)
	}

	users, err = c.ListUsers()
	if err != nil {
		t.Fatalf("failed to list users. %v", err)
	}
	if len(users) != 1 {
		t.Errorf("invalid user list. %+v", users)
	}

	t.Logf("%+v", users)
}
