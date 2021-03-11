package localstorage

import (
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserLocalStorage(t *testing.T) {
	storage := NewUserLocalStorage()

	user := &models.User{
		Username:      "let_robots_reign",
		Email:         "sample@ya.ru",
		Password:      "1234",
		MoviesWatched: 4,
		ReviewsNumber: 2,
	}

	t.Run("CreateUser", func(t *testing.T) {
		err := storage.CreateUser(user)
		assert.NoError(t, err)
	})

	t.Run("SuccessfulGetByUsername", func(t *testing.T) {
		gotUser, err := storage.GetUserByUsername("let_robots_reign")
		assert.NoError(t, err)
		assert.Equal(t, user, gotUser)
	})

	t.Run("CheckPassword", func(t *testing.T) {
		checkPass, err := storage.CheckPassword("1234", user)
		assert.NoError(t, err)
		assert.True(t, checkPass)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		_, err := storage.UpdateUser(user, models.User{
			Username: "let_robots_reign",
			Email:    "corrected@ya.ru",
			Password: "12345",
		})
		assert.NoError(t, err)
		updatedUser, err := storage.GetUserByUsername("let_robots_reign")
		assert.NoError(t, err)
		assert.Equal(t, "corrected@ya.ru", updatedUser.Email)
		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte("12345")))
	})
}

func TestLocalStorageErrors(t *testing.T) {
	storage := NewUserLocalStorage()

	user := &models.User{
		Username:      "let_robots_reign",
		Email:         "sample@ya.ru",
		Password:      "1234",
		MoviesWatched: 4,
		ReviewsNumber: 2,
	}

	t.Run("UnsuccessfulGetByUsername", func(t *testing.T) {
		_, err := storage.GetUserByUsername("unknown")
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("UpdateNonExistentUser", func(t *testing.T) {
		nonExistentUser := &models.User{
			Username: "nonexistent_user",
			Email:    "corrected@ya.ru",
			Password: "12345",
		}
		_, err := storage.UpdateUser(nonExistentUser, models.User{
			Username: "nonexistent_user",
			Email:    "new_email@ya.ru",
			Password: "qwerty",
		})

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("UpdateWrongUser", func(t *testing.T) {
		nonExistentUser := models.User{
			Username: "nonexistent_user",
			Email:    "corrected@ya.ru",
			Password: "12345",
		}
		_, err := storage.UpdateUser(user, nonExistentUser)

		assert.Error(t, err)
		assert.Equal(t, "username doesn't match", err.Error())
	})
}
