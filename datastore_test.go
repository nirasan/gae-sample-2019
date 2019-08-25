package main

import "testing"

func TestDatastoreClient_AddUser(t *testing.T) {
	c, err := NewDatastoreClient()
	if err != nil {
		t.Fatalf("failed to create client. %v", err)
	}

	id, err := c.AddUser(&User{Name: "user1", Age: 10})
	if err != nil {
		t.Fatalf("failed to add user. %v", err)
	}

	u, err := c.GetUser(id)
	if err != nil {
		t.Fatalf("failed to get user. %v", err)
	}

	if u.Name != "user1" || u.Age != 10 {
		t.Errorf("invalid user. %v", u)
	}
	t.Logf("%+v", u)

	users, err := c.ListUsers()
	if err != nil {
		t.Fatalf("failed to list users. %v", err)
	}
	if len(users) != 1 {
		t.Errorf("invalid user list. %v", users)
	}
	t.Logf("%+v", users)
}
