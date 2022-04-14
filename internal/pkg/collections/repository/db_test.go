package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	mylog "codex/internal/pkg/utils/log"
	"encoding/binary"
	"errors"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"strconv"
)

type testRow struct {
	inQuery    string
	bodyString string
	out        string
	status     int
	name       string
}

func MockDatabase() (*database.DBManager, pgxmock.PgxPoolIface, error) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		mylog.Error(errors.New("failed to create mock"))
	}
	return &database.DBManager{Pool: mock}, mock, err
}

var mockPersonPreview = domain.Actor{
	// Id:         1,
	// NameEn:     "Miley",
	// NameRus:    "Сайрус",
	// PictureUrl: "/miley.webp",
	// Career:     []string{"Актриса"},

	Id:           "1",
	Imgsrc:       "tales.webp",
	Name:         "Баба Яга",
	NameOriginal: "Baba Yaga",
	Career:       "Актёр",
	Height:       "160 см (без ступы)",
	Birthday:     "Неизвестно",
	Birthplace:   "Русь",
	Genres:       "Сказки",
	Total:        "500",
	
}

func TestGetSuccess(t *testing.T) {
	mdb, pool, err := MockDatabase()
	assert.Equal(t, nil, err, "create a mock")
	repository := InitColRep(mdb)
	defer pool.Close()
	mockCollections := domain.Collection{
		Title:        "Топ 256",
		Description: "Неповторимые must see",
		// type Feed struct {
		// 	Description string `json:"description"`
		// 	ImgSrc      string `json:"imgsrc"`
		// 	Page        string `json:"page"`
		// 	Num         string `json:"number"`
		// }
	}

	type Collection struct {
		Title       string       `json:"title"`
		Description string       `json:"description"`
		MovieList   []domain.MovieBasic `json:"movielist"`
	}

	countByte := make([]uint8, 8)
	binary.BigEndian.PutUint64(countByte, uint64(1))
	// numByte := make([]uint8, 8)

	Num, err := strconv.ParseUint("1", 10, 64)
	
	// binary.BigEndian.PutUint64(numByte, Num)
	rowsCount := pgxmock.NewRows([]string{"count"}).AddRow(countByte)
	rowsColl := pgxmock.NewRows([]string{"Title", "Description"}).AddRow([]uint8(mockCollections.Title), []uint8(mockCollections.Description))

	pool.ExpectBegin()
	pool.ExpectQuery(regexp.QuoteMeta(queryCountCollections)).WithArgs().WillReturnRows(rowsCount)
	pool.ExpectCommit()
	pool.ExpectBegin()
	pool.ExpectQuery(regexp.QuoteMeta(queryGetCollections)).WithArgs(10, 0).WillReturnRows(rowsColl)
	pool.ExpectCommit()

	actual, err := repository.GetCollection(Num)
	assert.NoError(t, err)
	assert.Equal(t, mockCollections, actual)
}
