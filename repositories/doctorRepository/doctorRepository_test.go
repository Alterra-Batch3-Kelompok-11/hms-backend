package doctorRepository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"hms-backend/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type doctorTestSuite struct {
	suite.Suite

	repository *doctorRepository
	mocking    sqlmock.Sqlmock
	dbMock     *sql.DB
}

func (s *doctorTestSuite) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	s.NoError(err)

	dbGormMock, errOpen := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbMock,
	}))
	s.NoError(errOpen)

	s.mocking = mock
	s.dbMock = dbMock
	s.repository = &doctorRepository{dbGormMock}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(doctorTestSuite))
}

func (s *doctorTestSuite) TestGetAll() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE `doctors`.`deleted_at` IS NULL")).
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name     string
		expected []models.Doctor
	}{
		{
			"berhasil",
			Expecteds,
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetAll()

		if s.NoError(err) {
			s.Equal(len(testCase.expected), len(resRep))
			s.Equal(testCase.expected[0].ID, resRep[0].ID)
			s.Equal(testCase.expected[0].UserId, resRep[0].UserId)
			s.Equal(testCase.expected[0].LicenseNumber, resRep[0].LicenseNumber)
			s.Equal(testCase.expected[0].Email, resRep[0].Email)
			s.Equal(testCase.expected[0].Phone, resRep[0].Phone)
			s.Equal(testCase.expected[0].MaritalStatus, resRep[0].MaritalStatus)
		} else {
			fmt.Println("wassa ", err.Error())
		}
	}
}
func (s *doctorTestSuite) TestGetById() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE ID = ? AND `doctors`.`deleted_at` IS NULL ORDER BY `doctors`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  models.Doctor
	}{
		{
			"Success",
			1,
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetById(testCase.id)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestGetByUserId() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE user_id = ? AND `doctors`.`deleted_at` IS NULL ORDER BY `doctors`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  models.Doctor
	}{
		{
			"Success",
			1,
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetByUserId(testCase.id)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestGetByLicenseNumber() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE license_number = ? AND `doctors`.`deleted_at` IS NULL ORDER BY `doctors`.`id` LIMIT 1")).
		WithArgs("1234567890").
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name          string
		licenseNumber string
		isHasData     bool
		expected      models.Doctor
	}{
		{
			"Success",
			"1234567890",
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetByLicenseNumber(testCase.licenseNumber)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestGetByLicenseNumberOther() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE (license_number = ? AND ID != ?) AND `doctors`.`deleted_at` IS NULL ORDER BY `doctors`.`id` LIMIT 1")).
		WithArgs("1234567890", 1).
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name          string
		id            uint
		licenseNumber string
		isHasData     bool
		expected      models.Doctor
	}{
		{
			"Success",
			1,
			"1234567890",
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetByLicenseNumberOther(testCase.licenseNumber, testCase.id)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestGetBySpecialityId() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE speciality_id = ? AND `doctors`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(Rows).RowsWillBeClosed()

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  models.Doctor
	}{
		{
			"Success",
			1,
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.GetBySpecialityId(testCase.id)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep[0].ID)
				s.Equal(testCase.expected.UserId, resRep[0].UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep[0].LicenseNumber)
				s.Equal(testCase.expected.Email, resRep[0].Email)
				s.Equal(testCase.expected.Phone, resRep[0].Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep[0].MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestCount() {
	s.mocking.
		ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `doctors` WHERE `doctors`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(1)).RowsWillBeClosed()

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  int64
	}{
		{
			"Success",
			1,
			true,
			1,
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.Count()

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected, resRep)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestCreate() {
	s.mocking.ExpectBegin()

	s.mocking.ExpectExec(regexp.QuoteMeta("INSERT INTO `doctors`")).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)

	s.mocking.ExpectCommit()

	s.mocking.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `doctors` WHERE `doctors`.`deleted_at` IS NULL")).
		WillReturnRows(Rows)

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  models.Doctor
	}{
		{
			"Success",
			1,
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.Create(testCase.expected)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}

	}
}
func (s *doctorTestSuite) TestUpdate() {
	s.mocking.ExpectBegin()

	s.mocking.ExpectExec(regexp.QuoteMeta("UPDATE `doctors` SET `id`=?,`updated_at`=?,`user_id`=?,`speciality_id`=?,`license_number`=?,`phone`=?,`email`=? WHERE ID = ? AND `doctors`.`deleted_at` IS NULL")).
		WithArgs(Expecteds[0].ID, AnyTime{}, Expecteds[0].UserId, Expecteds[0].SpecialityId, Expecteds[0].LicenseNumber, Expecteds[0].Phone, Expecteds[0].Email, Expecteds[0].ID).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)

	s.mocking.ExpectCommit()

	var testCases = []struct {
		name      string
		id        uint
		isHasData bool
		expected  models.Doctor
	}{
		{
			"Success",
			1,
			true,
			Expecteds[0],
		},
	}

	for _, testCase := range testCases {
		resRep, err := s.repository.Update(testCase.id, testCase.expected)

		if testCase.isHasData {
			if s.NoError(err) {
				s.Equal(testCase.expected.ID, resRep.ID)
				s.Equal(testCase.expected.UserId, resRep.UserId)
				s.Equal(testCase.expected.LicenseNumber, resRep.LicenseNumber)
				s.Equal(testCase.expected.Email, resRep.Email)
				s.Equal(testCase.expected.Phone, resRep.Phone)
				s.Equal(testCase.expected.MaritalStatus, resRep.MaritalStatus)
			}
		} else {
			s.Equal("record not found", err.Error())
		}
	}
}
func (s *doctorTestSuite) TestDelete() {

	s.mocking.ExpectBegin()

	s.mocking.ExpectExec(regexp.QuoteMeta("UPDATE `doctors` SET `deleted_at`=? WHERE `doctors`.`id` = ? AND `doctors`.`deleted_at` IS NULL")).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)).
		WillReturnError(nil)

	s.mocking.ExpectCommit()

	var testCases = []struct {
		name      string
		id        uint
		isSuccess bool
	}{
		{
			"Success",
			1,
			true,
		},
	}

	for _, testCase := range testCases {
		err := s.repository.Delete(testCase.id)

		if testCase.isSuccess {
			s.NoError(err)
		} else {
			s.Equal("record not found", err.Error())
		}
	}
}

var Rows = sqlmock.NewRows([]string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
	"user_id",
	"speciality_id",
	"license_number",
	"profile_pic",
	"birth_date",
	"phone",
	"marital_status",
	"email",
}).AddRow(
	1,
	nil,
	nil,
	nil,
	1,
	1,
	"1234567890",
	"",
	nil,
	"0812121213",
	false,
	"fulan@mail.com",
)

var NullRows = sqlmock.NewRows([]string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
	"user_id",
	"speciality_id",
	"license_number",
	"profile_pic",
	"birth_date",
	"phone",
	"marital_status",
	"email",
})

var Expecteds = []models.Doctor{
	{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		UserId:        1,
		SpecialityId:  1,
		LicenseNumber: "1234567890",
		ProfilePic:    "",
		BirthDate:     time.Time{},
		Phone:         "0812121213",
		MaritalStatus: false,
		Email:         "fulan@mail.com",
	},
}
