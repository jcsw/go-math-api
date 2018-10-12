package math

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/jcsw/go-math-api/pkg/driver/syslog"
	"github.com/jcsw/go-math-api/pkg/entity"
)

const (
	databaseName   = "admin"
	collectionName = "math-operation-log"
)

type operationLogCollection struct {
	ID         objectid.ObjectID `bson:"_id"`
	Operation  string            `bson:"operation"`
	Parameters []float64         `bson:"parameters"`
	Result     float64           `bson:"result"`
}

// RepositoryMongodb define the data mongodb repository
type RepositoryMongodb struct {
	MongoClient *mongo.Client
}

func (repository *RepositoryMongodb) operationLogCollection() (*mongo.Collection, error) {
	if repository.MongoClient == nil {
		return nil, errors.New("could not communicate with database")
	}
	return repository.MongoClient.Database(databaseName).Collection(collectionName, nil), nil
}

// PersistOperationLog function to persist OperationLog
func (repository *RepositoryMongodb) PersistOperationLog(request entity.MathOperationRequest, response entity.MathOperationResponse) error {

	operationLog := operationLogCollection{
		ID:         objectid.New(),
		Operation:  string(request.Operation),
		Parameters: request.Parameters,
		Result:     response.Result,
	}

	collection, err := repository.operationLogCollection()
	if err != nil {
		syslog.Error("p=math f=PersistOperationLog operationLog=%+v \n%v", operationLog, err)
		return err
	}

	if _, err := collection.InsertOne(nil, operationLog); err != nil {
		syslog.Error("p=math f=PersistOperationLog operationLog=%+v \n%v", operationLog, err)
		return err
	}

	syslog.Info("p=math f=PersistOperationLog operationLog=%+v", operationLog)
	return nil
}
