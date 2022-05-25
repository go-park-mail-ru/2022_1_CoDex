package usrrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbUserRepository struct {
	dbm *database.DBManager
}

func InitUsrRep(manager *database.DBManager) domain.UserRepository {
	return &dbUserRepository{
		dbm: manager,
	}
}

func (ur *dbUserRepository) GetById(id uint64) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetById, id)
	if len(resp) == 0 {
		log.Warn("{GetById}")
		log.Error(domain.Err.ErrObj.NoUser)
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		log.Warn("{GetById} in query: " + queryGetById)
		log.Error(err)
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	row := resp[0]
	out := domain.User{
		Id:             cast.ToUint64(row[0]),
		Username:       cast.ToString(row[1]),
		Password:       "",
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		RepeatPassword: "",
	}

	return out, nil
}

func (ur *dbUserRepository) GetBookmarks(id uint64) ([]domain.Bookmark, error) {
	resp, err := ur.dbm.Query(queryUserExist, id)
	if err != nil {
		log.Warn("{GetBookmarks} in query: " + queryUserExist)
		log.Error(err)
		return []domain.Bookmark{}, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		log.Warn("{GetBookmarks}")
		log.Error(domain.Err.ErrObj.InvalidId)
		return []domain.Bookmark{}, domain.Err.ErrObj.InvalidId
	}

	resp, err = ur.dbm.Query(queryGetUserBookmarks, id)
	if err != nil {
		log.Warn("{GetBookmarks} in query: " + queryGetUserBookmarks)
		log.Error(err)
		return []domain.Bookmark{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.Bookmark{}, nil
	}

	out := make([]domain.Bookmark, 0)
	for i := range resp {
		out = append(out, domain.Bookmark{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Description: cast.ToString(resp[i][1]),
			Imgsrc:      cast.ToString(resp[i][2]),
			Public:      cast.ToBool(resp[i][3]),
		})
	}

	return out, nil
}

func (ur *dbUserRepository) UpdateUser(id uint64, upd domain.UpdUser) (domain.User, error) {
	_, err := ur.dbm.Query(queryUpdateUser, upd.Username, id)
	if err != nil {
		log.Warn("{UpdateUser} in query: " + queryUpdateUser)
		log.Error(err)
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	usr, err := ur.GetById(id)
	if err != nil {
		return domain.User{}, err
	}

	return usr, nil
}

func (ur *dbUserRepository) GetUserReviews(id uint64) ([]domain.UserReview, error) {
	// get rating
	resp, err := ur.dbm.Query(queryGetUserRatings, id)
	if err != nil {
		log.Warn("{GetUserReviews} in query: " + queryGetUserRatings)
		log.Error(err)
		return []domain.UserReview{}, domain.Err.ErrObj.InternalServer
	}
	out := make([]domain.UserReview, 0)
	if len(resp) != 0 {
		for i := range resp {
			out = append(out, domain.UserReview{
				MovieId: cast.IntToStr(cast.ToUint64(resp[i][0])),
				Type:    "1",

				Rating: cast.IntToStr(cast.ToUint64(resp[i][1])),

				Date:         "",
				FeedbackType: "",
				MovieTitle:   "",
			})
		}
	}

	// get comments
	resp, err = ur.dbm.Query(queryGetUserComments, id)
	if err != nil {
		log.Warn("{GetUserReviews} in query: " + queryGetUserComments)
		log.Error(err)
		return []domain.UserReview{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) != 0 {
		for i := range resp {
			tmp := domain.UserReview{
				MovieId: cast.IntToStr(cast.ToUint64(resp[i][0])),
				Type:    "2",

				Rating: "",

				Date:         cast.TimeToStr(cast.ToTime(resp[i][1]), true),
				FeedbackType: "",
				MovieTitle:   cast.ToString(resp[i][3]),
			}

			ft := cast.ToString(resp[i][2])
			if ft == "good" {
				tmp.FeedbackType = "1"
			} else if ft == "bad" {
				tmp.FeedbackType = "3"
			} else {
				tmp.FeedbackType = "2"
			}

			out = append(out, tmp)
		}
	}

	return out, nil
}

func (ur *dbUserRepository) UpdateAvatar(clientId uint64, url string) (domain.User, error) {
	_, err := ur.dbm.Query(queryUpdAvatarByUsID, clientId, url)
	if err != nil {
		return domain.User{}, err
	}

	updated, err := ur.GetById(clientId)
	if err != nil {
		return domain.User{}, err
	}

	return updated, err
}
