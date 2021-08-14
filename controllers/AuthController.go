package controllers

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/punokawan/Simple-Web-Dompet-Kilat/database"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/helpers"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/models"
	"github.com/punokawan/Simple-Web-Dompet-Kilat/security"

	// "github.com/punokawan/Simple-Web-Dompet-Kilat/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
var user = new(models.Users)
var input = new(models.Users)

func SignUp(c *fiber.Ctx) error {
	userCollection := database.MI.DB.Collection("users")
	ctx, _ := context.WithTimeout(ctx, 10*time.Second)

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return helpers.ResponseMsg(
			c,
			fiber.StatusUnprocessableEntity,
			false,
			"Failed to parse body",
			nil)
	}

	if govalidator.IsNull(input.Username) {
		return helpers.ResponseMsg(
			c,
			fiber.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrEmptyUsername).Message,
			nil)
	}

	if govalidator.IsNull(input.Password) {
		return helpers.ResponseMsg(
			c,
			fiber.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrEmptyPassword).Message,
			nil)
	}

	input.Email = helpers.NormalizeEmail(input.Email)
	if !govalidator.IsEmail(input.Email) {
		return helpers.ResponseMsg(
			c,
			fiber.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrInvalidEmail).Message,
			nil)
	}

	findEmail := userCollection.FindOne(ctx, bson.M{"email": input.Email})
	if err := findEmail.Err(); err != nil {
		if strings.TrimSpace(input.Password) == "" {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(helpers.NewJError(helpers.ErrEmptyPassword))
		}

		input.Password, err = security.EncryptPassword(input.Password)
		if err != nil {
			return helpers.ResponseMsg(
				c,
				fiber.StatusBadRequest,
				false,
				helpers.NewJError(err).Message,
				nil)
		}

		input.CreatedAt = time.Now()
		input.UpdatedAt = input.CreatedAt
		input.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		result, err := userCollection.InsertOne(ctx, input)

		if err != nil {
			return helpers.ResponseMsg(
				c,
				fiber.StatusInternalServerError,
				false,
				"user failed to register",
				nil)
		}

		return helpers.ResponseMsg(
			c,
			fiber.StatusCreated,
			true,
			"user registered successfully",
			result)
	}

	return helpers.ResponseMsg(
		c,
		fiber.StatusBadRequest,
		false,
		helpers.NewJError(helpers.ErrEmailAlreadyExists).Message,
		nil)
}

func SignIn(c *fiber.Ctx) error {
	userCollection := database.MI.DB.Collection("users")
	ctx, _ := context.WithTimeout(ctx, 10*time.Second)

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return helpers.ResponseMsg(
			c,
			fiber.StatusUnprocessableEntity,
			false,
			"Failed to parse body",
			nil)
	}

	if govalidator.IsNull(input.Password) {
		return helpers.ResponseMsg(
			c,
			fiber.StatusBadRequest,
			false,
			helpers.NewJError(helpers.ErrEmptyPassword).Message,
			nil)
	}

	input.Email = helpers.NormalizeEmail(input.Email)

	findUser := userCollection.FindOne(ctx, bson.M{"email": input.Email})
	if err := findUser.Err(); err != nil {
		return helpers.ResponseMsg(
			c,
			fiber.StatusUnauthorized,
			false,
			helpers.NewJError(helpers.ErrInvalidCredentials).Message,
			nil)
	}

	err := findUser.Decode(&user)
	if err != nil {
		return helpers.ResponseMsg(
			c,
			fiber.StatusNotFound,
			false,
			helpers.NewJError(err).Message,
			nil)
	}

	log.Printf("user : %s, \ninput: %s", user, input)
	err = security.VerifyPassword(user.Password, input.Password)
	if err != nil {
		return helpers.ResponseMsg(
			c,
			fiber.StatusUnauthorized,
			false,
			helpers.NewJError(helpers.ErrInvalidCredentials).Message,
			nil)
	}
	token, err := security.NewToken(user.ID.Hex())
	if err != nil {
		return helpers.ResponseMsg(
			c,
			fiber.StatusUnauthorized,
			false,
			helpers.NewJError(err).Message,
			nil)
	}
	response := fiber.Map{
		"user":  user,
		"token": strings.Join([]string{"Bearer ", token}, ""),
	}
	return helpers.ResponseMsg(
		c,
		fiber.StatusOK,
		true,
		"Login scuccessful",
		response)
}
