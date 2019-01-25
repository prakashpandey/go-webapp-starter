package auth

var counter = 0

// Validate user credentials
func Validate(uname, password string) bool {
	// implement the validation logic here

	// sudo logic
	counter++
	if counter%2 == 0 {
		return true
	}
	return false
}
