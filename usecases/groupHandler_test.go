package usecases_test

import (
	"github.com/AlecSmith96/dnd-scheduler/usecases"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
type GroupHandlerSuite struct {
	suite.Suite
	DB           *gorm.DB
	mock         sqlmock.Sqlmock
	groupHandler usecases.GroupHandler
}

func (s *GroupHandlerSuite) SetupSuite() {
	// var (
	// 	db  *sql.DB
	// 	err error
	// )
	// db, s.mock, err = sqlmock.New()
	// require.NoError(s.T(), err)

	// s.DB, err = gorm.Open(gorm.Config{}, db)
	// require.NoError(s.T(), err)

	// s.DB.LogMode(true)

	// s.groupHandler = *usecases.NewGroupHandler(s.DB)
}
