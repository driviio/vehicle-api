package db

import (
	"cloud.google.com/go/datastore"
	"fmt"
	"golang.org/x/net/context"
)

type VehicleLog struct {
	ID int64
	VehicleID int64
	Data string
}

type vehicleDB struct {
	client *datastore.Client
}

type VehicleDatabase interface {
	ListVehicleLog() ([]*VehicleLog, error)

	AddVehicleLog(b *VehicleLog) (id int64, err error)

	GetVehicleLog(id int64) (*VehicleLog, error)

	Close()
}

// Ensure vehicleDB conforms to the VehicleDatabase interface.
var _ VehicleDatabase = &vehicleDB{}

func NewVehicleDatabase(client *datastore.Client) (VehicleDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &vehicleDB{
		client: client,
	}, nil
}

// Close closes the database.
func (db *vehicleDB) Close() {
	db.client.Close()
}

func (db *vehicleDB) vehicleLogKey(id int64) *datastore.Key {
	return datastore.IDKey("VehicleLog", id, nil)
}

// GetVehicleLog retrieves a book by its ID.
func (db *vehicleDB) GetVehicleLog(id int64) (*VehicleLog, error) {
	ctx := context.Background()
	k := db.vehicleLogKey(id)
	book := &VehicleLog{}
	if err := db.client.Get(ctx, k, book); err != nil {
		return nil, fmt.Errorf("datastoredb: could not get Book: %v", err)
	}
	book.ID = id
	return book, nil
}

// AddVehicleLog saves a given ad, assigning it a new ID.
func (db *vehicleDB) AddVehicleLog(b *VehicleLog) (id int64, err error) {
	ctx := context.Background()
	k := datastore.IncompleteKey("VehicleLog", nil)
	k, err = db.client.Put(ctx, k, b)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put ad: %v", err)
	}
	return k.ID, nil
}

// ListVehicleLog returns a list of ads, ordered by title.
func (db *vehicleDB) ListVehicleLog() ([]*VehicleLog, error) {
	ctx := context.Background()
	ads := make([]*VehicleLog, 0)
	q := datastore.NewQuery("VehicleLog")

	keys, err := db.client.GetAll(ctx, q, &ads)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list ads: %v", err)
	}

	for i, k := range keys {
		ads[i].ID = k.ID
	}

	return ads, nil
}