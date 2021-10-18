package nosql

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/EdwBaeza/echo-app/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "users"

var (
	onceUserRepository     sync.Once
	instanceUserRepository *UserRepository
)

type UserRepository struct {
	Repository
}

//BuildClient for mongodb
func (repository *UserRepository) BuildClient() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	repository.client = client
	repository.collection = client.Database(os.Getenv("MONGODB")).Collection(COLLECTION_NAME)
	repository.context, _ = context.WithTimeout(context.Background(), 60*time.Second)
	err = repository.client.Connect(repository.context)

	return err
}

//CloseClient for mongodb
func (repository *UserRepository) CloseClient() {
	repository.client.Disconnect(repository.context)
}

// NewUserRepository return a instance of repository
func NewUserRepository() *UserRepository {
	onceUserRepository.Do(func() {
		instanceUserRepository = new(UserRepository)
	})
	instanceUserRepository.BuildClient()

	return instanceUserRepository
}

//Save user in mongodb
func (repository *UserRepository) Save(user domain.User) (*domain.User, error) {
	result, err := repository.collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Println("Error insert one: ", err)
		return &user, err
	}
	log.Println("ID created user", result.InsertedID)

	createdUser, _ := repository.Find(result.InsertedID.(primitive.ObjectID).Hex())

	return createdUser, nil
}

//Find user in mongodb
func (repository *UserRepository) Find(id string) (*domain.User, error) {
	user := &domain.User{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error find by id ", err.Error())
		return user, err
	}

	filter := bson.M{"_id": objectID}
	findOneError := repository.collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, findOneError
}

//All user in mongodb
func (repository *UserRepository) All(pageSize int, pageNumber int) (*domain.UserPage, error) {
	userPage := &domain.UserPage{
		Data: []*domain.User{},
		Page: domain.Page{
			PageSize:   pageSize,
			PageNumber: pageNumber,
		},
	}
	userPage.SetLinks()

	cursor, err := repository.collection.Find(context.TODO(), bson.D{})

	defer cursor.Close(repository.context)
	for cursor.Next(repository.context) {

		user := &domain.User{}
		if err = cursor.Decode(user); err != nil {
			log.Println(err)
		}
		userPage.Data = append(userPage.Data, user)
	}
	startPage := pageSize * pageNumber
	endPage := startPage + pageSize
	userPage.Count = len(userPage.Data)
	userPage.Data = userPage.Data[startPage:endPage]

	return userPage, err
}
