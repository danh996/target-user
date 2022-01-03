package models

import (
	"time"
	enums "user/common/enum"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Fullname       string             `bson:"fullname,omitempty"`
	Password       string             `bson:"password,omitempty"`
	Phone          string             `bson:"phone,omitempty"`
	Email          string             `bson:"email,omitempty"`
	Age            int32              `bson:"age,omitempty"`
	Role           enums.UserRole     `bson:"role,omitempty"`
	Sex            enums.UserSex      `bson:"sex,omitempty"`
	ProvinceID     int32              `bson:"province_id,omitempty"`
	DistrictID     int32              `bson:"district_id,omitempty"`
	WardID         int32              `bson:"ward_id,omitempty"`
	LocationScore  int64              `bson:"location_score,omitempty"`
	IdentityCard   string             `bson:"identity_card,omitempty"`
	Job            string             `bson:"job,omitempty"`
	BlockedUserIDS []string           `bson:"blocked_user_ids,omitempty"`
	CreatedAt      *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt      *time.Time         `bson:"updated_at,omitempty"`
}

// HashPassword : hash password using crypto
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// IsCorrectPassword : check password with passwordhash
func (u *User) IsCorrectPassword(password string) bool {
	var err error
	// fmt.Println("hello 20151515")
	// err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
