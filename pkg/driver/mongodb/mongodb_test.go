// +build integration

package mongodb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jcsw/go-math-api/pkg/driver/properties"
)

func TestShouldInitializeMongoDBSession(t *testing.T) {

	properties.AppProperties =
		properties.Properties{
			MongoDB: properties.MongoDBProperties{
				Hosts:     []string{"localhost:27017"},
				Database:  "admin",
				Username:  "go-math-api",
				Password:  "admin",
				Timeout:   500,
				PoolLimit: 1,
			}}

	InitializeMongoDBSession()
	defer CloseMongoDBSession()

	if assert.True(t, IsMongoDBSessionAlive()) {

		mongoDBSession := RetrieveMongoDBSession()
		defer mongoDBSession.Close()

		err := mongoDBSession.Ping()
		assert.Nil(t, err)
	}

}
