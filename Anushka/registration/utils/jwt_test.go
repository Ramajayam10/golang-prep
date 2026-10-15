package utils

import "testing"

func TestGenerateToken(t *testing.T) {
	userID := 4
	email := "anushka@testing.com"

	token, err := GenerateToken(userID, email)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if token == "" {
		t.Errorf("Got an empty string")
	}
}

func TestValidateToken_valid(t *testing.T) {
	token, _ := GenerateToken(1, "anushka@test.com")

	claims, err := ValidateToken(token)
	if err != nil {
		t.Errorf("Expected no error , got %v", err)
	}

	if claims["email"] != "anushka@test.com" {
		t.Errorf("Expected email %s , got %s", "anushka@test.com", claims["email"])
	}
}

func TestValidateToken_Invalid(t *testing.T) {
	fakeToken := "this.is.a.fake.token"
	_, err := ValidateToken(fakeToken)

	if err == nil {
		t.Errorf("Expected an error for invalid token, but got nil")
	}
}
