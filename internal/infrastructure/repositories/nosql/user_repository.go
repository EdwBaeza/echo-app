package nosql

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/EdwBaeza/echo-app/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceUserRepository     sync.Once
	instanceUserRepository *Repository
)

// Repository by nosql (mongo)
type Repository struct {
	client  *mongo.Client
	context context.Context
}

//GetClient for mongodb
func (repository *Repository) GetClient() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017/echoapp"))
	repository.client = client
	repository.context, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = repository.client.Connect(repository.context)

	return err
}

//CloseClient for mongodb
func (repository *Repository) CloseClient() {
	repository.client.Disconnect(repository.context)
}

// NewUserRepository return a instance of repository
func NewUserRepository() *Repository {
	onceUserRepository.Do(func() {
		instanceUserRepository = new(Repository)
	})
	return instanceUserRepository
}

//Save user in mongodb
func (repository *Repository) Save(user domain.User) (domain.User, error) {
	collection := repository.client.Database("echoapp").Collection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalln("Error insert one: ", err)
		return user, err
	}
	log.Println("ID created user", result.InsertedID)

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

//Find user in mongodb
func (repository *Repository) Find(id string) (domain.User, error) {
	var user domain.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln("Error get id ", err.Error())
		return user, err
	}

	collection := repository.client.Database("echoapp").Collection("users")
	filter := bson.M{"_id": objectID}
	findOneError := collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, findOneError
}

//All user in mongodb
func (repository *Repository) All() ([]domain.User, error) {
	var users []domain.User

	collection := repository.client.Database("echoapp").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	defer cursor.Close(repository.context)
	for cursor.Next(repository.context) {
		var user domain.User
		if err = cursor.Decode(&user); err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}

	return users, err
}
