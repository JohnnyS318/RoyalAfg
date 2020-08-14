package database

import (
	"royalafg/pkg/shared/pkg/models"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type Users struct {
	logger *zap.SugaredLogger
	coll   *mgm.Collection
}

func NewUserDatabase(logger *zap.SugaredLogger) *Users {
	coll := mgm.Coll(&models.User{})

	logger.Warn("Collection")

	return &Users{
		logger: logger,
		coll:   coll,
	}
}

func (db *Users) CreateUser(user *models.User) error {
	err := user.Validate()

	if err != nil {
		return err
	}

	db.coll.Create(user)
	db.logger.Info("Inserted new User ", user.Username)
	return nil
}

func (db *Users) DeleteUser(user *models.User) error {
	db.coll.Delete(user)
	return nil
}

func (db *Users) FindById(id string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindByID(id, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *Users) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *Users) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
