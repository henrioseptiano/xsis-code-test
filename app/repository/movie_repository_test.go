package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
	"xsis-code-test/models/model"
)

func NewRepoMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestCreateMovie(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"movies\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()
	var reqMovie model.Movie
	repo.CreateMovie(reqMovie)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateMovie_Error(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	mock.ExpectBegin()
	mock.ExpectCommit()
	var reqMovie model.Movie
	err := repo.CreateMovie(reqMovie)
	assert.NotNil(t, err)
}

func TestUpdateMovie(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	expectedSQL := "UPDATE \"movies\" SET \"updated_at\"=.+ WHERE id =.+"
	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	var reqMovie model.Movie
	err := repo.UpdateMovie(1, reqMovie)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateMovie_Error(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	expectedSQL := "UPDATE \"movies\" WHERE id =.+"
	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	var reqMovie model.Movie
	err := repo.UpdateMovie(1, reqMovie)
	assert.NotNil(t, err)
}

func TestDeleteMovie(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	expectedSQL := "UPDATE \"movies\" SET \"deleted_at\"=.+ WHERE id =.+"
	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.DeleteMovie(1)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteMovie_Error(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	repo := NewAppRepository(db)
	expectedSQL := "UPDATE \"movies\" WHERE id =.+"
	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()
	err := repo.DeleteMovie(3)
	assert.NotNil(t, err)
}

func TestFindMovie_shouldFound(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	implObj := NewAppRepository(db)
	movies := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "beranakdalamkubur", "kubur dalam anak", 5, "ini.jpg", time.Now(), time.Now(), nil)

	expectedSQL := "SELECT (.+) FROM \"movies\" WHERE deleted_at is null"
	mock.ExpectQuery(expectedSQL).WillReturnRows(movies)
	_, res := implObj.ListMovie()
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindMovie_shouldNotFound(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	implObj := NewAppRepository(db)
	movies := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"})
	expectedSQL := "SELECT (.+) FROM \"movies\" WHERE deleted_at is null"
	mock.ExpectQuery(expectedSQL).WillReturnRows(movies)
	_, res := implObj.ListMovie()
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestImplementation_GetMovieById(t *testing.T) {
	sqlDB, db, mock := NewRepoMock(t)
	defer sqlDB.Close()

	implObj := NewAppRepository(db)
	movies := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "beranakdalamkubur", "kubur dalam anak", 5, "ini.jpg", time.Now(), time.Now(), nil)

	movieSQL := "SELECT (.+) FROM \"movies\" WHERE id =.+"
	mock.ExpectQuery(movieSQL).WillReturnRows(movies)
	_, res := implObj.GetMovie(1)
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}
