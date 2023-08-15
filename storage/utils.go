package storage

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
)

const PasswordEncryptionCost = 15

func NewId() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}.String()
	}
	return id.String()
}

func GenerateRandomString(length int) string {
	vowels := []rune{'e', 'u', 'i', 'o', 'a'}
	consonants := []rune{'q', 'r', 't', 'p', 's', 'd', 'g', 'h', 'k', 'z', 'x', 'v', 'b', 'n', 'm'}

	str := strings.Builder{}

	for i := 0; i < length; i += 2 {
		str.WriteRune(consonants[rand.Intn(len(consonants))])
		if i != length-1 {
			str.WriteRune(vowels[rand.Intn(len(vowels))])
		}
	}

	return str.String()
}

func GenerateRandomName() string {
	adjectives := []string{"playing", "running", "jumping", "thinking", "flying", "sleeping", "singing", "dancing", "painting", "walking", "dreaming", "coding", "sitting", "drinking", "swimming", "reading", "chatting", "rolling", "trolling", "smiling"}
	animals := []string{"cat", "dog", "mouse", "tiger", "lion", "wolf", "fox", "rabbit", "lizard", "leopard", "turtle", "elephant", "panda", "bear", "raccoon", "camel", "dinosaur", "hamster", "bird", "fish"}

	adjIdx := rand.Intn(len(adjectives))
	animalIdx := rand.Intn(len(animals))

	return adjectives[adjIdx] + "-" + animals[animalIdx]
}

///////////////////////////////////////////////////////////////////////////////

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordEncryptionCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

///////////////////////////////////////////////////////////////////////////////

func DefaultSettings(userId string) Settings {
	settings := Settings{
		UserId:           userId,
		Language:         "en",
		EmulatorSettings: make(map[string]EmulatorSettings),
	}
	for _, system := range PlatformIds {
		settings.EmulatorSettings[system] = DefaultEmulatorSettings(system)
	}
	return settings
}
