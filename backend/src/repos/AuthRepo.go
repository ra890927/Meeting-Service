package repos

import (
	"meeting-center/src/models"

	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(user *models.User) (*models.User, *string, error)
	Logout(user *models.User, token *string) error
}

type authRepo struct {
	dataBase    *gorm.DB
	redisClient *redis.Client
}

func NewAuthRepo(authRepoArgs ...authRepo) AuthRepo {
	if len(authRepoArgs) == 1 {
		return AuthRepo(&authRepo{
			dataBase:    authRepoArgs[0].dataBase,
			redisClient: authRepoArgs[0].redisClient,
		})
	} else if len(authRepoArgs) == 0 {
		db, err := gorm.Open(sqlite.Open("../sqlite.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		}
		red := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		_, err = red.Ping(red.Context()).Result()
		if err != nil {
			panic("Failed to connect to redis")
		}
		return AuthRepo(&authRepo{
			dataBase:    db,
			redisClient: red,
		})
	} else {
		panic("too many arguments")
	}
}

func (ar authRepo) Login(user *models.User) (*models.User, *string, error) {
	// check if the user with the given email exists
	var existingUser models.User
	result := ar.dataBase.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		return nil, nil, result.Error
	}
	// check if the password is correct
	// hash the password of the input user's password
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return nil, nil, err
	}

	// generate a token, check if such a token exists in the redis database,
	// if it does, generate a new token
	// if it does not, save the token in the redis database
	cnt := 0
	hashString := ""
	for {
		hash, err := bcrypt.GenerateFromPassword([]byte(existingUser.Email+time.Now().String()+string(cnt)), bcrypt.DefaultCost)
		hashString = string(hash)
		// check if the hash exists in the redis database
		_, err = ar.redisClient.Get(ar.redisClient.Context(), hashString).Result()
		if err != nil {
			break
		}
		cnt++
	}
	// save the token in the redis database
	err = ar.redisClient.Set(ar.redisClient.Context(), hashString, existingUser.ID, 0).Err()
	if err != nil {
		return nil, nil, err
	}

	// return the user with the token
	return &existingUser, &hashString, nil
}

func (ar authRepo) Logout(user *models.User, token *string) error {
	// check if the token exists in the redis database
	_, err := ar.redisClient.Get(ar.redisClient.Context(), *token).Result()
	if err != nil {
		return err
	}
	// delete the token from the redis database
	err = ar.redisClient.Del(ar.redisClient.Context(), *token).Err()
	if err != nil {
		return err
	}

	return nil
}
