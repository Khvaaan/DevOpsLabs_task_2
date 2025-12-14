package mongo

import (
  "context"
  "time"

  "go-news/pkg/storage"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
  client *mongo.Client
  col    *mongo.Collection
}

func New(uri string) (*DB, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
  if err != nil {
    return nil, err
  }

  // without defer client.Disconnect(...) here

  db := client.Database("news_app")
  col := db.Collection("tasks")

  // index id (unical'niy)
  _, _ = col.Indexes().CreateOne(ctx, mongo.IndexModel{
    Keys:    bson.D{{Key: "id", Value: 1}},
    Options: options.Index().SetUnique(true),
  })

  return &DB{client: client, col: col}, nil
}

func (m *DB) Tasks() ([]storage.Task, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  cur, err := m.col.Find(ctx, bson.D{})
  if err != nil {
    return nil, err
  }
  defer cur.Close(ctx)

  var tasks []storage.Task
  for cur.Next(ctx) {
    var t storage.Task
    if err := cur.Decode(&t); err != nil {
      return nil, err
    }
    tasks = append(tasks, t)
  }
  return tasks, cur.Err()
}

func (m *DB) AddTask(t storage.Task) error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  _, err := m.col.InsertOne(ctx, t)
  return err
}

func (m *DB) UpdateTask(t storage.Task) error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  res, err := m.col.UpdateOne(
    ctx,
    bson.M{"id": t.ID},
    bson.M{"$set": bson.M{
      "title":      t.Title,
      "done":       t.Done,
      "created_at": t.CreatedAt,
    }},
  )
  if err != nil {
    return err
  }
  if res.MatchedCount == 0 {
    return storage.ErrNotFound
  }
  return nil
}

func (m *DB) DeleteTask(t storage.Task) error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  res, err := m.col.DeleteOne(ctx, bson.M{"id": t.ID})
  if err != nil {
    return err
  }
  if res.DeletedCount == 0 {
    return storage.ErrNotFound
  }
  return nil
}
