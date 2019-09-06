package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

type datastoreClient struct {
	client *datastore.Client
}

func NewDatastoreClient() (*datastoreClient, error) {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &datastoreClient{client}, nil
}

func (c *datastoreClient) GetUser(id int64) (*User, error) {
	ctx := context.Background()
	k := datastore.IDKey("User", id, nil)
	u := &User{}
	if err := c.client.Get(ctx, k, u); err != nil {
		return nil, fmt.Errorf("datastoreClient: could not get User: %v", err)
	}
	u.ID = id
	return u, nil
}

func (c *datastoreClient) AddUser(u *User) (int64, error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("User", nil)
	k, err := c.client.Put(ctx, k, u)
	if err != nil {
		return 0, fmt.Errorf("datastoreClient: could not put User: %v", err)
	}
	return k.ID, nil
}

func (c *datastoreClient) DeleteUser(id int64) error {
	ctx := context.Background()
	k := datastore.IDKey("User", id, nil)
	err := c.client.Delete(ctx, k)
	if err != nil {
		return fmt.Errorf("datastoreClient: could not delete User: %v", err)
	}
	return nil
}

func (c *datastoreClient) ListUsers() ([]*User, error) {
	ctx := context.Background()
	users := make([]*User, 0)
	q := datastore.NewQuery("User").
		Order("Name")

	keys, err := c.client.GetAll(ctx, q, &users)

	if err != nil {
		return nil, fmt.Errorf("datastoreClient: could not list users: %v", err)
	}

	for i, k := range keys {
		users[i].ID = k.ID
	}

	return users, nil
}
