package handler

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sarthak7509/goftodo/config"
	"github.com/sarthak7509/goftodo/database"
	"github.com/sarthak7509/goftodo/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password, err
	}
	return string(bytes), nil
}

func comparePasswords(hashedPassword []byte, enteredPassword string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(enteredPassword))
	return err
}

func generateToken() (string, error) {
	secretKey := []byte(config.Config("SECRET_KEY"))
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Handler to signUpUser
func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	payload := new(models.UserSignUp)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	hashedPassword, err := HashPassword(payload.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	token, err := generateToken()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	newUser := models.User{
		Id:       uuid.New(),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		ApiKey:   token,
	}

	errs := db.Create(&newUser).Error
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": models.UserSignUpResponse(newUser)})
}

func SingIn(c *fiber.Ctx) error {
	db := database.DB
	payload := new(models.UserSignIn)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user := models.User{}

	db.Find(&user, "Email=?", payload.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	err = comparePasswords([]byte(user.Password), payload.Password)

	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong Password"})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    user.ApiKey,
		Path:     "/",
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": models.UserSignInResponse(user)})
}

func ProtectedViewSet(c *fiber.Ctx,user models.User) error{
	return c.Status(fiber.StatusAccepted).JSON(models.UserSignUpResponse(user))
}
