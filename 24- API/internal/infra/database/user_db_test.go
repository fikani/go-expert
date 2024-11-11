package database

import (
	"app-example/internal/entity"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type UserDBTestSuite struct {
	suite.Suite
	userDB *User
	db     *sql.DB
}

func (suite *UserDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	err = CreateUserTables(db)
	if err != nil {
		panic(err)
	}

	suite.db = db
	suite.userDB = NewUser(db)
}

func (suite *UserDBTestSuite) TearDownTest() {
	suite.db.Close()
}

func (suite *UserDBTestSuite) TestUserDB_Create_ShouldPersistUser() {
	user, _ := entity.NewUser("John Doe", "john@mail.com", "password")
	err := suite.userDB.Create(user)
	suite.Nil(err)

	foundUser, err := suite.userDB.FindByID(user.ID)
	suite.Nil(err)
	suite.Equal(user.ID, foundUser.ID)
	suite.Equal(user.Name, foundUser.Name)
	suite.Equal(user.Email, foundUser.Email)
	suite.Equal(user.Password, foundUser.Password)
	suite.NotEmpty(foundUser.Password)
}

func (suite *UserDBTestSuite) TestUserDB_FindByEmail_ShouldReturnUser() {
	user, _ := entity.NewUser("John Doe", "john@mail.com", "password")
	err := suite.userDB.Create(user)
	suite.Nil(err)

	foundUser, err := suite.userDB.FindByEmail(user.Email)
	suite.Nil(err)
	suite.Equal(user.ID.String(), foundUser.ID.String())
	suite.Equal(user.Name, foundUser.Name)
	suite.Equal(user.Email, foundUser.Email)
	suite.Equal(user.Password, foundUser.Password)
	suite.NotEmpty(foundUser.Password)
}

func TestUserDBSuite(t *testing.T) {
	suite.Run(t, new(UserDBTestSuite))
}
