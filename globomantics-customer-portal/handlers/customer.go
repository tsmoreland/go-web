package handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tsmoreland/go-web/globomantics-customer-portal/data"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

var (
	tpl            *template.Template
	c              *http.Cookie
	oauthConfig    *oauth2.Config
	clientId       string
	clientSecret   string
	loginUrl       string
	logoutUrl      string
	oauthLogoutUrl string
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	clientId = os.Getenv("AZURE_CLIENT_ID")
	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	loginUrl = os.Getenv("AZUREB2C_LOGIN_REDIRECT_URL")
	logoutUrl = os.Getenv("AZUREB2C_LOGOUT_REDIRECT_URL")
	authUrl := os.Getenv("AZURE_AUTH_URL") // https://(azure domain).b2clogin.com/(azure domain).onmicrosoft.com..)
	tokenUrl := os.Getenv("AZURE_TOKEN_URL")
	oauthLogoutUrl = os.Getenv("AZURE_LOGOUT_URL")
	if anyAreEmpty(clientId, clientSecret, loginUrl, logoutUrl, authUrl, tokenUrl, oauthLogoutUrl) {
		log.Fatalln("One or more required environment variables are not set")
	}

	oauthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		Scopes: []string{"openid", clientId, "offline_access"},
	}
}

func Customers(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("customer")
	_, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		log.Println("No cookie found.  Redirecting to home page")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	customer := data.GetCustomerByIdOrNil(customerId)
	if customer != nil {
		if err := tpl.ExecuteTemplate(w, "customer.gohtml", customer); err != nil {
			log.Fatalln("Not able to call the template", err)
		}
	} else {
		// use a not found template
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err == http.ErrNoCookie {
			urlstring := oauthConfig.AuthCodeURL("thisstate")
			http.Redirect(w, r, urlstring, http.StatusSeeOther)
			return
		}
		v := r.URL.Query()
		v.Add("customer", cookie.Value)
		r.URL.RawQuery = v.Encode()
		next.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != "thisstate" {
		http.Redirect(w, r, "/home?status=nosuccess", http.StatusSeeOther)
		return
	}
	if r.FormValue("error") == "access_denied" {
		http.Redirect(w, r, "/home?status=nosuccess", http.StatusSeeOther)
		return
	}
	oauthtoken, err := oauthConfig.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalln("Token is not valid") // probably just want to redirect in this case and log
	}
	accessToken, err := jwt.Parse(oauthtoken.AccessToken, nil)

	var customerId string
	var isNewCustomer string

	customerClaims := accessToken.Claims.(jwt.MapClaims) // type assertion, ensures Claims is a jwt.MapClaims
	customerId = fmt.Sprintf("%v", customerClaims["sub"])
	email := fmt.Sprintf("%v", customerClaims["emails"])
	email = strings.TrimLeft(email, "[")
	email = strings.TrimRight(email, "]")

	isNewCustomer = fmt.Sprintf("%v", customerClaims["newUser"])
	if isNewCustomer == "true" {
		log.Println("new customer logging in")
	}

	customer := &data.Customer{
		CustomerID:   customerId,
		FirstName:    fmt.Sprintf("%v", customerClaims["given_name"]),
		LastName:     fmt.Sprintf("%v", customerClaims["family_name"]),
		Address:      fmt.Sprintf("%v", customerClaims["country"]),
		Phone:        "-----------",
		Email:        email,
		SubType:      fmt.Sprintf("%v", customerClaims["extension_Subscription"]),
		Active:       true,
		CreationDate: time.Now().Format("January 2, 2006"),
	}

	if ok, err := data.AddCustomer(customer); !ok || err != nil {
		log.Println(err)
	}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    customerId,
		HttpOnly: true,
		Secure:   false, // should be true but for testing purposes we're leaving as false - ok for local dev
	}

	http.SetCookie(w, cookie)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	u, _ := url.Parse(oauthLogoutUrl)
	q := u.Query()
	q.Add("post_logout_redirect_uri", logoutUrl)
	u.RawQuery = q.Encode()

	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func anyAreEmpty(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}
