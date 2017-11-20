package userLogin

//LoginUser ...
// Attempts to login a user, given uname and pass.
func LoginUser(uname string, pass string) (bool, error) {
	// See if the password is valid
	res, err := ValidatePass(uname, pass)

	if err != nil {
		return false, err
	}

}
