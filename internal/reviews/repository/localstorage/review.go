package localstorage

import (
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/models"
	"strconv"
	"sync"
)

type ReviewLocalStorage struct {
	reviews   map[string]*models.Review
	currentID uint64
	mutex     sync.Mutex
}

func NewReviewLocalStorage() *ReviewLocalStorage {
	return &ReviewLocalStorage{
		reviews: make(map[string]*models.Review),
		currentID: 1,
	}
}

func (storage *ReviewLocalStorage) CreateReview(review *models.Review) error {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	review.ID = strconv.FormatUint(storage.currentID, 10)
	storage.reviews[review.ID] = review
	storage.currentID++

	return nil
}

func (storage *ReviewLocalStorage) CheckIfExists(username string, review *models.Review) bool {
	userReviews := storage.GetUserReviews(username)
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	for _, r := range userReviews {
		if r.MovieID == review.MovieID {
			return true
		}
	}
	return false
}

func (storage *ReviewLocalStorage) GetUserReviews(username string) []*models.Review {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	var userReviews []*models.Review

	for _, review := range storage.reviews {
		if review.Author == username {
			userReviews = append(userReviews, review)
		}
	}
	return userReviews
}

func (storage *ReviewLocalStorage) GetMovieReviews(movieID string) []*models.Review {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	var movieReviews []*models.Review

	for _, review := range storage.reviews {
		if review.MovieID == movieID {
			movieReviews = append(movieReviews, review)
		}
	}
	return movieReviews
}
