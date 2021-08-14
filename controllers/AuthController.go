package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/punokawan/Simple-Web-Dompet-Kilat/database"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/helpers"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/models"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/security"

	// "github.com/punokawan/Simple-Web-Dompet-Kilat/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/asaskevich/govalidator.v9"
)

// type AuthController interface {
// 	SignUp(ctx *fiber.Ctx) error
// SignIn(ctx *fiber.Ctx) error
// GetUser(ctx *fiber.Ctx) error
// GetUsers(ctx *fiber.Ctx) error
// PutUser(ctx *fiber.Ctx) error
// DeleteUser(ctx *fiber.Ctx) error
// }

var ctx = context.Background()
var users = new(models.Users)

func SignUp(c *fiber.Ctx) error {
	userCollection := database.MI.DB.Collection("users")
	ctx, _ := context.WithTimeout(ctx, 10*time.Second)

	if err := c.BodyParser(&users); err != nil {
		log.Println(err)
		return helpers.ResponseMsg(
			c,
			http.StatusUnprocessableEntity,
			false,
			"Failed to parse body",
			nil)
	}

	if govalidator.IsNull(users.Username) {
		return helpers.ResponseMsg(
			c,
			http.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrEmptyUsername).Message,
			nil)
	}

	if govalidator.IsNull(users.Password) {
		return helpers.ResponseMsg(
			c,
			http.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrEmptyPassword).Message,
			nil)
	}

	users.Email = helpers.NormalizeEmail(users.Email)
	if !govalidator.IsEmail(users.Email) {
		return helpers.ResponseMsg(
			c,
			http.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrInvalidEmail).Message,
			nil)
	}

	findEmail := userCollection.FindOne(ctx, bson.M{"email": users.Email})
	if err := findEmail.Err(); err != nil {
		if strings.TrimSpace(users.Password) == "" {
			return c.
				Status(http.StatusBadRequest).
				JSON(helpers.NewJError(helpers.ErrEmptyPassword))
		}

		users.Password, err = security.EncryptPassword(users.Password)
		if err != nil {
			return helpers.ResponseMsg(
				c,
				http.StatusBadRequest,
				false,
				helpers.NewJError(err).Message,
				nil)
		}

		users.CreatedAt = time.Now()
		users.UpdatedAt = users.CreatedAt
		result, err := userCollection.InsertOne(ctx, users)

		if err != nil {
			return helpers.ResponseMsg(
				c,
				http.StatusInternalServerError,
				false,
				"user failed to register",
				nil)
		}

		return helpers.ResponseMsg(
			c,
			http.StatusCreated,
			true,
			"user registered successfully",
			result)
	}

	return helpers.ResponseMsg(
		c,
		http.StatusBadRequest,
		false,
		helpers.NewJError(helpers.ErrEmailAlreadyExists).Message,
		nil)
}
