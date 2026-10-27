package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string, pepper string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+pepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Compare(hashedPassword string, password string, pepper string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+pepper))
}
