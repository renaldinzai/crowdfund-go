package campaign

import (
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mocks := []Campaign{
		Campaign{
			ID:               1,
			UserID:           1,
			Name:             "Campaign 1",
			ShortDescription: "Short Description 1",
			Description:      "Description 1",
			Perks:            "Perks 1",
			BackerCount:      10,
			GoalAmount:       10,
			CurrentAmount:    100,
			Slug:             "Slug 1",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		Campaign{
			ID:               2,
			UserID:           2,
			Name:             "Campaign 2",
			ShortDescription: "Short Description 2",
			Description:      "Description 2",
			Perks:            "Perks 2",
			BackerCount:      10,
			GoalAmount:       10,
			CurrentAmount:    100,
			Slug:             "Slug 2",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "short_description", "description", "perks",
		"backer_count", "goal_amount", "current_amount", "slug", "created_at", "updated_at"}).
		AddRow(mocks[0].ID, mocks[0].UserID, mocks[0].Name, mocks[0].ShortDescription, mocks[0].Description, mocks[0].Perks,
			mocks[0].BackerCount, mocks[0].GoalAmount, mocks[0].CurrentAmount, mocks[0].Slug, mocks[0].CreatedAt, mocks[0].UpdatedAt).
		AddRow(mocks[1].ID, mocks[1].UserID, mocks[1].Name, mocks[1].ShortDescription, mocks[1].Description, mocks[1].Perks,
			mocks[1].BackerCount, mocks[1].GoalAmount, mocks[1].CurrentAmount, mocks[1].Slug, mocks[1].CreatedAt, mocks[1].UpdatedAt)

	query := "SELECT * FROM `campaigns`"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := NewRepository(gdb)

	if _, err = r.FindAll(); err != nil {
		t.Errorf("expected no error, got error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mocks := []Campaign{
		Campaign{
			ID:               1,
			UserID:           1,
			Name:             "Campaign 1",
			ShortDescription: "Short Description 1",
			Description:      "Description 1",
			Perks:            "Perks 1",
			BackerCount:      10,
			GoalAmount:       10,
			CurrentAmount:    100,
			Slug:             "Slug 1",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	imageMocks := []CampaignImage{
		CampaignImage{
			ID:         1,
			CampaignID: 1,
			FileName:   "image-1",
			IsPrimary:  1,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		CampaignImage{
			ID:         2,
			CampaignID: 1,
			FileName:   "image-2",
			IsPrimary:  0,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "short_description", "description", "perks",
		"backer_count", "goal_amount", "current_amount", "slug", "created_at", "updated_at"}).
		AddRow(mocks[0].ID, mocks[0].UserID, mocks[0].Name, mocks[0].ShortDescription, mocks[0].Description, mocks[0].Perks,
			mocks[0].BackerCount, mocks[0].GoalAmount, mocks[0].CurrentAmount, mocks[0].Slug, mocks[0].CreatedAt, mocks[0].UpdatedAt)

	imageRows := sqlmock.NewRows([]string{"id", "campaign_id", "file_name", "created_at", "updated_at"}).
		AddRow(imageMocks[0].ID, imageMocks[0].CampaignID, imageMocks[0].FileName, imageMocks[0].CreatedAt, imageMocks[0].UpdatedAt)

	query := "SELECT * FROM `campaigns` WHERE user_id = ?"

	imageQuery := "SELECT * FROM `campaign_images` WHERE `campaign_images`.`campaign_id` = ? AND campaign_images.is_primary = 1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta(imageQuery)).WithArgs(1).WillReturnRows(imageRows)

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := NewRepository(gdb)

	if _, err = r.FindByUserID(1); err != nil {
		t.Errorf("expected no error, got error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
