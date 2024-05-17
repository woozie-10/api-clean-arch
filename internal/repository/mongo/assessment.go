package mongo

import (
	"context"

	"github.com/woozie-10/api-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssessmentRepository struct {
	db *mongo.Collection
}

func NewAssessmentRepository(db *mongo.Database, collection string) *AssessmentRepository {
	return &AssessmentRepository{
		db: db.Collection(collection),
	}
}

func (r AssessmentRepository) Get(ctx context.Context, username string) (*domain.Assessment, error) {
	var result *domain.Assessment
	filter := bson.M{"username": username}
	err := r.db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r AssessmentRepository) Add(ctx context.Context, assessment *domain.Assessment) error {
	filter := bson.M{"username": assessment.Username}
	update := bson.D{}
	for subject, grades := range assessment.Marks {
		update = append(update, bson.E{
			Key: "$push",
			Value: bson.D{
				{Key: "marks." + subject, Value: bson.D{{Key: "$each", Value: grades}}},
			},
		})
	}
	_, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
