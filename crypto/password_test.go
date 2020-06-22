package crypto

import "testing"

func TestEncryptPassword(t *testing.T) {
	pwd := "Hello"
	pwd2 := "hello"

	// $2a$10$1iDZKLJ3HV7aGszdpZa0SenngmhfHsuRyg8b9PDkQyz2vIBjnQPVG
	encPwd, err := EncryptPassword(pwd)
	t.Logf("enc pwd: pwd=%v, err=%v", encPwd, err)

	in := []string{
		pwd,
		pwd2,
	}

	// validate:
	for _, item := range in {
		t.Logf("validate pwd: pwd=%v, ok=%v", item, ValidatePassword(encPwd, item))
	}
}
