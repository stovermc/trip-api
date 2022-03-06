package mongo

import (
	"context"

	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Trip struct {
	ID string `bson:"identifier"`
	Name string `bson:"name"`
}

type tripRepository struct {
	client *mongo.Client
	db string
}

func NewTripRepository(c *mongo.Client, db string) app.TripRepository {
 return &tripRepository{
	 client: c,
	 db: db,
 }
}

func (tr *tripRepository) Add(ctx context.Context, t domain.Trip) error {
	doc := &Trip{}
	doc.FromModel(t)
	filter := bson.D{{"identifier", t.ID()}}

	err := tr.collection().FindOne(ctx, filter).Err()
	if err == nil {
		return app.ErrTripAlreadyExists
	}

	if err != mongo.ErrNoDocuments {
		return err
	}

	_, err = tr.collection().InsertOne(ctx, doc)

	return err
}

func (tr *tripRepository) Get(ctx context.Context, id string) (domain.Trip, error) {
	doc := &Trip{}
	err := tr.collection().FindOne(ctx, bson.D{{"identifier", id}}).Decode(doc)
	if err == mongo.ErrNoDocuments{
		return nil, app.ErrTripNotFound
	}
	if err != nil {
		return nil, err
	}

	return doc.ToModel(), nil
}

func (tr *tripRepository) Save(ctx context.Context, t domain.Trip) error {
	doc := &Trip{}
	doc.FromModel(t)
	filter := bson.D{{"identifier", t.ID()}}

	err := tr.collection().FindOneAndReplace(ctx, filter, doc).Decode(&Trip{})
	if err == mongo.ErrNoDocuments {
		return app.ErrTripNotFound
	}

	return err
}

func (tr *tripRepository) collection() *mongo.Collection {
	return tr.client.Database(tr.db).Collection("trips")
}

func (doc *Trip) ToModel() domain.Trip {
		return domain.NewTrip(doc.ID, doc.Name)
}

func (doc *Trip) FromModel(model domain.Trip) {
		doc.ID = model.ID()
		doc.Name = model.Name()
}
