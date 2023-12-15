package localserver

import (
	"go-gql-dynamodb-template/handler/graph/model"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

type DDBMnager struct {
	DB *dynamo.DB
}

func (d *DDBMnager) Migration() error {
	var tableMap = map[string]interface{}{}
	tableMap[os.Getenv("TODO_TABLE_NAME")] = &model.Todo{}
	tableMap[os.Getenv("USER_TABLE_NAME")] = &model.User{}

	for table, modelInstance := range tableMap {
		exist, err := d.TableExists(table)
		if err != nil {
			return errors.Wrap(err, "TableExists")
		}
		if exist {
			log.Printf("Table %s already exists", table)
			continue
		}

		if err := d.TableCreate(table, modelInstance); err != nil {
			return errors.Wrap(err, "TableCreate")
		}
	}
	return nil
}

func (d *DDBMnager) TableExists(tableName string) (bool, error) {
	if _, err := d.DB.Table(tableName).Describe().Run(); err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			return false, nil
		}
		return false, errors.Wrap(err, "dynamo.DB.Table.Describe")
	}
	return true, nil
}

func (d *DDBMnager) TableCreate(tableName string, model interface{}) error {
	if err := d.DB.CreateTable(tableName, model).Run(); err != nil {
		return  errors.Wrapf(err, "dynamo.DB.Table.Describe(%s)", tableName)
	}
	return nil
}
