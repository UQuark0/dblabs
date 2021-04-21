package common

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/volatiletech/null"
	"math/rand"
)

type Product struct {
	Id   int
	Name string
}

type Fat struct {
	Id   int
	Type string
}

type Info struct {
	Id        int
	ProductId int
	FatId     int
	Price     float64

	Product Product `json:"-"`
	Fat     Fat `json:"-"`
}

type Sale struct {
	InfoId    int
	Amount    int
	Sold      null.Time
	Purchased null.Time

	Info Info `json:"-"`
}

type User struct {
	Id int
	Username string
	Password string
	Salt string
	AllowSelect bool
	AllowUpdate bool
	AllowInsert bool
	AllowDelete bool
}

func generateSalt() string {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const saltLength = 16
	salt := ""
	for i := 0; i < saltLength; i++ {
		salt += string(alphabet[rand.Intn(len(alphabet))])
	}
	return salt
}

func hashAndSalt(password, salt string) string {
	p1 := sha256.Sum256([]byte(password))
	p2 := hex.EncodeToString(p1[:]) + salt
	p3 := sha256.Sum256([]byte(p2))
	return hex.EncodeToString(p3[:])
}

func (u *User) ValidatePassword(password string) bool {
	return hashAndSalt(password, u.Salt) == u.Password
}

func (u *User) GeneratePassword(password string) {
	u.Salt = generateSalt()
	u.Password = hashAndSalt(password, u.Salt)
}