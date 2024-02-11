package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "menu/domain"
)

var result struct {
	Id    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title"`
	Uuid  string             `bson:"uuid"`
}

type MenuMongoRepository struct {
	mongoDao *MongoDao
}

func NewMenuMongoRepository(mongoDao *MongoDao) *MenuMongoRepository {
	return &MenuMongoRepository{
		mongoDao,
	}
}

func (mmr *MenuMongoRepository) Add(menu *Menu) (*Menu, error) {
	newMenu := NewMenu(
		NewCafeId(NewUuid()),
		NewTitle(menu.GetTitle().GetValue()),
	)

	data := mmr.mongoDao.Query("sample", "movies")
	doc := bson.M{"title": newMenu.GetTitle().GetValue(), "uuid": newMenu.GetCafeId().GetValue()}
	_, err := data.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	return newMenu, nil
}

func (mmr *MenuMongoRepository) Find(title Title) (*Menu, error) {
	data := mmr.mongoDao.Query("sample", "movies")
	filter := bson.M{"title": title.GetValue()}
	err := data.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return NewMenu(
		NewCafeId(NewUuidWithUuid(result.Uuid)),
		NewTitle(result.Title),
	), nil
}

func (mmr *MenuMongoRepository) FindById(uuid Uuid) *Menu {
	data := mmr.mongoDao.Query("sample", "movies")
	filter := bson.M{"uuid": uuid.GetValue()}
	err := data.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil
	}
	return NewMenu(
		NewCafeId(NewUuidWithUuid(result.Uuid)),
		NewTitle(result.Title),
	)
}

func (mmr *MenuMongoRepository) FindAllBy(page int, limit int, sort string) ([]*Menu, MetaData) {
	var menus []*Menu
	data := mmr.mongoDao.Query("sample", "movies")
	countTotal, err := data.CountDocuments(context.TODO(), bson.D{}, options.Count().SetHint("_id_"))
	if err != nil {
		return nil, MetaData{}
	}

	var sortValue = 1
	if sort == "desc" {
		sortValue = -1
	}
	limitInt64 := int64(limit)
	skip := int64(page)*limitInt64 - limitInt64
	pageTotal := int(countTotal) / limit
	opts := options.FindOptions{Limit: &limitInt64, Skip: &skip, Sort: bson.M{"title": sortValue}}
	cursor, err := data.Find(context.Background(),
		bson.D{{}},
		&opts,
	)
	if err != nil {
		return nil, MetaData{}
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&result)
		if err != nil {
			return nil, MetaData{}
		}
		menus = append(menus, NewMenu(
			NewCafeId(NewUuidWithUuid(result.Uuid)),
			NewTitle(result.Title),
		))
	}

	if err = cursor.Err(); err != nil {
		return nil, MetaData{}
	}

	return menus, NewMetaData(page, "ASC", pageTotal, int(countTotal), limit)
}
