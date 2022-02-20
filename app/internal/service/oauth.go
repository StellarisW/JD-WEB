package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"net/http"
	"os"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token" binding:"required"`
}

func OauthHandler(c *gin.Context) {
	clientID := g.Config.Auth.Oauth.ClientID
	clientSecret := g.Config.Auth.Oauth.ClientSecret
	//var rc *reqCode
	//err := r.ParseForm(&rc)
	//if err != nil {
	//	fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
	//	r.Response.WriteHeader(http.StatusBadRequest)
	//}
	code := c.PostForm("code")
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?"+
		"client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
	req.Header.Set("accept", "application/json")
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()
	var t AccessTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
	c.Writer.Header().Set("Location", "/login/oauth/hello/?access_token="+t.AccessToken)
	c.Writer.WriteHeader(http.StatusNotFound)
	//r.Middleware.Next()
}
