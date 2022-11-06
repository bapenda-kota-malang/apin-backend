package auth

type TokenType string

const AccessToken = "Access"
const RefreshToken = "Refresh"

type AuthInfo struct {
	Uuid    string
	User_Id int
	// Position int16
	Ref_Id int
}
