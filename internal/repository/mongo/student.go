package mongo

import (
	"context"

	"github.com/woozie-10/api-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentRepository struct {
	db *mongo.Collection
}

func NewStudentRepository(db *mongo.Database, collection string) *StudentRepository {
	return &StudentRepository{
		db: db.Collection(collection),
	}
}

func (r StudentRepository) Get(ctx context.Context) ([]*domain.Student, error) {
	var students []*domain.Student
	filter := bson.M{}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (r StudentRepository) GetByUsername(ctx context.Context, username string) (*domain.Student, error) {
	var student *domain.Student
	filter := bson.M{"username": username}
	err := r.db.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (r StudentRepository) Add(ctx context.Context, student *domain.Student) error {
	_, err := r.db.InsertOne(ctx, student)
	if err != nil {
		return err
	}
	return nil
}

func (r StudentRepository) Delete(ctx context.Context, username string) error {
	filter := bson.M{"username": username}
	_, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r StudentRepository) GetByGroup(ctx context.Context, group string) ([]*domain.Student, error) {
	var students []*domain.Student
	filter := bson.M{"group": group}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (r StudentRepository) GetByCourse(ctx context.Context, course string) ([]*domain.Student, error) {
	var students []*domain.Student
	filter := bson.M{"course": course}
	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (r StudentRepository) Update(ctx context.Context, username string, newStudent *domain.Student) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": newStudent}
	opts := options.Update().SetUpsert(true)
	_, err := r.db.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil

}
