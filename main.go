package main

import (
	"errors"
	"math/rand"
	"time"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Player struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Password  string  `json:"-"`
	IPAddress string  `json:"ip_address"`
	OriginURL string  `json:"origin_url"`
	Wallet    float64 `json:"wallet"`
}

var players = make(map[string]Player)

func main() {
	e := echo.New()
	rand.Seed(time.Now().UnixNano())

	e.POST("/register", registerHandler)
	e.POST("/login", loginHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func registerHandler(c echo.Context) error {
	type RegisterRequest struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Password encryption failed"})
	}

	playerID := generatePlayerID()
	ip := c.RealIP()
	origin := c.Request().Header.Get("Origin")

	player := Player{
		ID:        playerID,
		Name:      req.Name,
		Phone:     req.Phone,
		Password:  string(hashedPassword),
		IPAddress: ip,
		OriginURL: origin,
		Wallet:    0.0,
	}

	players[player.Phone] = player

	return c.JSON(http.StatusCreated, player)
}

func loginHandler(c echo.Context) error {
	type LoginRequest struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid login request"})
	}

	player, err := findPlayerByPhone(req.Phone)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Player not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"player":  player,
	})
}

func findPlayerByPhone(phone string) (Player, error) {
	player, exists := players[phone]
	if !exists {
		return Player{}, errors.New("not found")
	}
	return player, nil
}

func generatePlayerID() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	id := ""
	for i := 0; i < 5; i++ {
		id += string(letters[rand.Intn(len(letters))])
	}
	for i := 0; i < 5; i++ {
		id += string(digits[rand.Intn(len(digits))])
	}
	return id
}
