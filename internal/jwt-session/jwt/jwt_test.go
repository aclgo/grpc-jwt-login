package jwt

import (
	"testing"
	"time"
)

func TestFunctionsJwtToken(t *testing.T) {
	fakeUserID := "my-fake-user-id"
	fakeUserRole := "admin"
	ttl := time.Nanosecond

	jwtObj := NewJwt("my-test-secret")

	tokenString, err := jwtObj.NewToken("access", fakeUserID, fakeUserRole, ttl)
	if err != nil {
		t.Error(err)
	}

	tokenParsed, err := jwtObj.ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}

	if !tokenParsed.Valid {
		t.Error("token invalid")
	}

	claims, err := jwtObj.GetClaims(tokenParsed)
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
