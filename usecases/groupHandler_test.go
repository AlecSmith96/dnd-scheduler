package usecases_test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/AlecSmith96/dnd-scheduler/usecases"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
type GroupHandlerSuite struct {
	suite.Suite
	DB           *gorm.DB
	sqlmock      sqlmock.Sqlmock
	groupHandler usecases.GroupHandler
}

func (s *GroupHandlerSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.sqlmock, err = sqlmock.New()
	if err != nil {
		log.Fatalf(err.Error(), "sql mock failed")
	}

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error(), "gorm open failed")
	}
	require.NoError(s.T(), err)

	s.groupHandler = usecases.GroupHandler{
		DB: s.DB,
	}
}

func (s *GroupHandlerSuite) TestGetAllGroups_HappyPath() {
	// https://blog.questionable.services/article/testing-http-handlers-go/

	// Given
	req, err := http.NewRequest("GET", "/api/group", nil)
	if err != nil {
		log.Fatalf("Unable to create GET request: %v", err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(s.groupHandler.GetAllGroups)

	groupId := uuid.New()
	s.sqlmock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "groups"`)).
		WillReturnRows(s.sqlmock.NewRows([]string{"id", "name"}).
			AddRow(groupId, "Group Name"))

	handler.ServeHTTP(responseRecorder, req)

	expectedResponse := fmt.Sprintf(
		"{\"groups\":[{\"id\":\"%s\",\"name\":\"Group Name\",\"sessions\":null}]}\n", 
		groupId)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code, "response return non-200 status")
	assert.Equal(s.T(), expectedResponse, responseRecorder.Body.String(), "unexpected response body")
}
