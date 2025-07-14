package tweet

import (
	"context"
	"time"

	"github.com/samirgattas/microblog/internal/core/domain"
	"github.com/samirgattas/microblog/internal/core/port/repository"
	"github.com/samirgattas/microblog/lib/customerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tweetID = int64(0) // This is the entity id

type tweetRepository struct {
	tweetDB *mongo.Database
}

func NewTweetRepository(tweetDB *mongo.Database) repository.TweetRepository {
	return &tweetRepository{
		tweetDB: tweetDB,
	}
}

func (r *tweetRepository) Save(ctx context.Context, tweet *domain.Tweet) error {
	now := time.Now()
	tweet.CreatedAt = &now
	res, err := r.tweetDB.Collection("Tweet").InsertOne(ctx, tweet)
	if err != nil {
		return err
	}

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return customerror.NewInternalServerError("no inserted")
	}
	tweet.ID = insertedID.Hex()
	return nil
}

func (r *tweetRepository) Get(ctx context.Context, tweetID int64) (*domain.Tweet, error) {
	tweet := domain.Tweet{}

	// Filter by id
	filter := bson.D{{"_id", tweetID}}
	res := r.tweetDB.Collection("Tweet").FindOne(ctx, filter)

	if err := res.Decode(tweet); err != nil {
		if err == mongo.ErrNoDocuments {
			return &domain.Tweet{}, customerror.NewNotFoundError("followed")
		}
		return &domain.Tweet{}, err
	}

	return &tweet, nil
}

// TODO: implement pagination
func (r *tweetRepository) Search(ctx context.Context, params domain.TweetSearchParams) (domain.TweetsSearchResult, error) {
	// Sort by created_at field. 1 -> ASC / -1 -> DESC
	options := options.Find().SetSort(bson.D{{"created_at", -1}})
	options.Limit = &params.Limit
	options.SetSkip(params.Offset)
	// Filter
	filter := bson.D{{"user_id", bson.D{{"$in", params.UserIDs}}}}
	cursor, err := r.tweetDB.Collection("Tweet").Find(ctx, filter, options)
	if err != nil {
		return domain.TweetsSearchResult{}, err
	}

	tweets := []domain.Tweet{}
	if err = cursor.All(ctx, &tweets); err != nil {
		panic(err)
	}

	paging := domain.Paging{
		Total:  10,
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	result := domain.TweetsSearchResult{
		Results: tweets,
		Paging:  paging,
	}
	return result, nil
}
