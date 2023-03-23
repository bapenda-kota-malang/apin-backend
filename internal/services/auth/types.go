package auth

type TokenType string

const AccessToken = "Access"
const RefreshToken = "Refresh"

type AuthInfo struct {
	Uuid             string
	User_Id          int
	Ref_Id           int
	Jabatan_Id       int
	BidangKerja_Kode string
}
