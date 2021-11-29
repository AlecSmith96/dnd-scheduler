package usecases_test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"

	"github.com/AlecSmith96/dnd-scheduler/usecases"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PlayerHandlerSuite struct {
	suite.Suite
	db      *gorm.DB
	sqlmock sqlmock.Sqlmock
	playerHandler usecases.PlayerHandler
}

func (s *PlayerHandlerSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.sqlmock, err = sqlmock.New()
	if err != nil {
		log.Fatalln(err.Error(), "sql mock failed")
	}

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error(), "gorm open failed")
	}
	require.NoError(s.T(), err)

	s.playerHandler = usecases.PlayerHandler{
		DB: s.db,
	}
}

func (s *PlayerHandlerSuite) TestGetAllPlayers_HappyPath() {
	// Given
	req, err := http.NewRequest("GET","/api/players", nil)
	if err != nil {
		log.Fatalln("Unable to create GET request: ", err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(s.playerHandler.GetAllPlayers)

	playerId := uuid.New()
	groupId := uuid.New()
	sessionID := uuid.New()
	from, _ := time.Parse("2006-01-02T15:04:05.0000000Z", "2021-11-18T21:54:23.2332927Z")
	to, _ := time.Parse("2006-01-02T15:04:05.0000000Z", "2021-11-19T00:54:23.2332927Z")
	s.sqlmock.ExpectQuery("SELECT * FROM \"groups\"").
		WillReturnRows(s.sqlmock.NewRows([]string{"id", "name"}).
		AddRow(groupId, "Group Name"))
}
