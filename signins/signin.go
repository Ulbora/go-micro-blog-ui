package signins

// SystemType SystemType
type SystemType int64

const (
	//LinkedIn LinkedIn
	LinkedIn SystemType = iota
	//Google Google
	Google
	//Facebook Facebook
	Facebook
)

// Signin Signin
type Signin interface {
	Authorization()
	AccessToken() *TokenResponse
}

// SystemSignin SystemSignin
type SystemSignin struct {
	//linked in
	LinkedInAuthParams  *AuthParams
	LinkedInTokanParams *TokenParams
}

// AuthParams AuthParams
type AuthParams struct {
	ResponseType string
	ClientID     string
	RedirectURI  string
	State        string
	Scope        string
	SysType      SystemType
}

// TokenParams TokenParams
type TokenParams struct {
	GrantType    string
	AuthCode     string
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// TokenResponse TokenResponse
type TokenResponse struct {
	AccessToken           string
	ExpiresIn             int
	RefreshToken          string
	RefreshTokenExpiresIn int
	Scope                 string
}
