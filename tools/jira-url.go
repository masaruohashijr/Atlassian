package tools

import (
	"Atlassian/config"
	"Atlassian/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sclevine/agouti"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var gmailConfig models.Config

func ScrapEmailAddress(accountId string) (emailAddress string) {

	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}
	jiraConfig := config.ReadConfig(models.JIRA)
	gmailConfig := config.ReadConfig(models.GMAIL)
	navTo := jiraConfig.JiraProfilePage + accountId
	if err := page.Navigate(navTo); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	time.Sleep(5 * time.Second)
	oauthButton := page.FindByID("google-auth-button")
	oauthButton.Click()
	email := page.FindByID("identifierId")
	email.Fill(gmailConfig.AdminEmail)
	time.Sleep(2 * time.Second)
	b := page.FindByButton("Próxima")
	println(b.Text())
	b.Click()
	time.Sleep(5 * time.Second)
	password := page.FindByName("password")
	password.Fill(gmailConfig.AdminPassword)
	b = page.FindByButton("Próxima")
	println(b.Text())
	b.Click()
	time.Sleep(20 * time.Second)
	codeField := page.FindByID("code")
	code := GetVerificationCode()
	codeField.Fill(code)
	print(code + "\n")
	b = page.FindByID("login-submit")
	b.Submit()
	time.Sleep(10 * time.Second)
	//html, _ := page.HTML()
	//println(html)
	println("***************************")
	time.Sleep(10 * time.Second)
	//html, _ = page.HTML()
	//println(html)
	println("***************************")
	time.Sleep(10 * time.Second)
	domainParts := strings.Split(gmailConfig.DomainParts, ",")
	emailAddress = getEmailAddress(page, domainParts...)
	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
	return emailAddress
}

func getEmailAddress(page *agouti.Page, domains ...string) (emailAddress string) {
	html, _ := page.HTML()
	for _, d := range domains {
		ini := strings.Index(html, d)
		if ini == -1 {
			println("I din't found " + d + " !!!")
			continue
		}
		left, right := html[:ini], html[ini:]
		closingLeft := strings.LastIndex(left, ">")
		openingRight := strings.Index(right, "<")
		emailAddress = left[closingLeft+1:] + right[:openingRight]
		println("I found " + emailAddress + "!!!")
		break
	}
	return
}

func GetVerificationCode() (response string) {
	b, err := ioutil.ReadFile("config/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Messages.List(user).Q("subject:'Verificação de e-mail da Atlassian'").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Messages) == 0 {
		fmt.Println("No labels found.")
		return ""
	}
	for _, l := range r.Messages {
		//println(i, l.Id)
		m, err := srv.Users.Messages.Get(user, l.Id).Do()
		srv.Users.History.List(user)

		if err != nil {
			log.Fatalf("Unable to retrieve labels: %v", err)
		}
		/*
			for _, header := range m.Payload.Headers {
				if header.Name == "Subject" {
					fmt.Println(header.Value)
				}
			}
		*/
		for _, part := range m.Payload.Parts {
			if part.MimeType == "text/html" {
				data, _ := base64.URLEncoding.DecodeString(part.Body.Data)
				html := string(data)
				i := strings.Index(html, "margin-top:24px\">")
				i += 17
				response = fmt.Sprintf(html[i : i+8])
				break
			}
		}
		if response != "" {
			break
		}
	}
	return response
}

type message struct {
	size    int64
	gmailID string
	date    string // retrieved from message header
	snippet string
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "config/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
