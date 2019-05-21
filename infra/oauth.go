package infra

import (
	"fmt"
	"golang.org/x/oauth2"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	oauthConfig      *oauth2.Config
	oauthStateString = randSeq(10)
)

func init() {
	endpoint := oauth2.Endpoint{
		AuthURL:  "https://api.secure.mercedes-benz.com/oidc10/auth/oauth/v2/authorize",
		TokenURL: "https://api.secure.mercedes-benz.com/oidc10/auth/oauth/v2/token",
	}
	oauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("CALLBACK_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"mb:vehicle:status:general mb:user:pool:reader"},
		Endpoint:     endpoint,
	}
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func HandleOauthLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GetToken(state string, code string) (*oauth2.Token, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	return token, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
