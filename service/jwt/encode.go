package jwtService

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ducthangng/geofleet/user-service/singleton"
	"github.com/golang-jwt/jwt/v5"
)

// SecretKey should be loaded from an Environment Variable in production
// e.g. os.Getenv("JWT_SECRET")

// UserClaims defines the custom data we want to store in the token

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GetSecretKey

var jwtKey string

func setKey() {
	if len(jwtKey) != 0 {
		return
	}

	cfg := singleton.GetConfig().Cookie.JWTKey
	jwtKey = cfg
}

// 1. ENCODE: Generate a new JWT token
func GenerateToken(userID, username, role string) (string, error) {
	if len(jwtKey) == 0 {
		setKey()
	}

	// Define the claims
	claims := UserClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Token expires in 15 minutes
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			// Token issued at now
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Issuer (optional)
			Issuer: "geofleet-service",
		},
	}

	// Create the token using HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	log.Println("key is: ", jwtKey)
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

// 2. COMPARE/VERIFY: Parse and validate the token
func ValidateToken(tokenString string) (*UserClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// SECURITY CHECK: Validate the Alg is what we expect (HMAC)
		// This prevents "None" algorithm attacks
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if token is valid and claims are of correct type
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// // Main function to demonstrate usage
// func main() {
// 	fmt.Println("--- JWT Demo ---")

// 	// Step 1: Login / Generate
// 	userID := "550e8400-e29b-41d4-a716-446655440000"
// 	token, err := GenerateToken(userID, "ducthangng", "admin")
// 	if err != nil {
// 		fmt.Println("Error generating:", err)
// 		return
// 	}
// 	fmt.Printf("Generated Token: %s\n\n", token)

// 	// Step 2: Middleware / Validate
// 	fmt.Println("--- Verifying Token ---")
// 	claims, err := ValidateToken(token)
// 	if err != nil {
// 		fmt.Printf("Error verifying: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Token Valid!\nUser ID: %s\nUser: %s\nRole: %s\nExpires: %s\n",
// 		claims.UserID, claims.Username, claims.Role, claims.ExpiresAt)
// }
