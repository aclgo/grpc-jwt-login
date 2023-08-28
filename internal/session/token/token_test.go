package token

import (
	"testing"
	"time"
)

func TestCreateJwtToken(t *testing.T) {
	fakeUserID := "my-fake-user-id"
	fakeUserRole := "admin"
	ttl := time.Hour

	tkAction := NewtokenAction("my-test-secret")

	tokenString, err := tkAction.NewToken("access", fakeUserID, fakeUserRole, ttl)
	if err != nil {
		t.Error(err)
	}

	tokenParsed, err := tkAction.ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	if !tokenParsed.Valid {
		t.Error("token invalid")
	}

	claims, err := tkAction.GetClaims(tokenParsed)
	if err != nil {
		t.Error("GetClaims")
	}

	if claims["id"].(string) != fakeUserID {
		t.Error("mistach user id")
	}

	if claims["role"].(string) != fakeUserRole {
		t.Error("mistach role")
	}

	time.Sleep(time.Second)

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		t.Error("token expired")
	}
}
