package types

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CRYPTCOST       = 18
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 6
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func ValidateUser(u CreateUserParams) map[string]string {
	err := make(map[string]string)

	if len(u.FirstName) < 2 {
		err["firstName"] = "user firstName should be ateast 2 chars"
	}
	if len(u.LastName) < 2 {
		err["lastName"] = "user lastName should be ateast 2 chars"
	}
	if len(u.Password) < 6 {
		err["password"] = "user password should be ateast 6 chars"
	}

	if !isValidateEmail(u.Email) {
		err["email"] = "email format is invalid"
	}

	return err
}

func isValidateEmail(email string) bool {

	emailExpr := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailExpr.MatchString(email)
}
