package tools

import (
	"Atlassian/config"
	"Atlassian/models"
	m "Atlassian/models"
	"log"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

var gmailConfig m.Config
var jiraConfig m.Config
var driver *agouti.WebDriver

func StartDriver() *agouti.Page {
	jiraConfig = config.ReadConfig(m.JIRA)
	gmailConfig = config.ReadConfig(m.GMAIL)
	driver = agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}
	return sudo()
}

func StopDriver() {
	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
}

func sudo() (page *agouti.Page) {
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}
	jiraConfig := config.ReadConfig(models.JIRA)
	gmailConfig := config.ReadConfig(models.GMAIL)
	if err := page.Navigate(jiraConfig.JiraProfilePage); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	time.Sleep(5 * time.Second)
	oauthButton := page.FindByID("google-auth-button")
	oauthButton.Click()
	email := page.FindByID("identifierId")
	email.Fill(gmailConfig.AdminEmail)
	time.Sleep(2 * time.Second)
	b := page.FindByButton(jiraConfig.JiraNextButton)
	println(b.Text())
	b.Click()
	time.Sleep(5 * time.Second)
	password := page.FindByName("password")
	password.Fill(gmailConfig.AdminPassword)
	b = page.FindByButton(jiraConfig.JiraNextButton)
	println(b.Text())
	b.Click()
	time.Sleep(20 * time.Second)
	codeField := page.FindByID("code")
	code := GetVerificationCode()
	codeField.Fill(code)
	print(code + "\n")
	b = page.FindByID("login-submit")
	b.Submit()
	time.Sleep(5 * time.Second)
	return page
}

func ScrapEmailAddress(page *agouti.Page, accountId string) (emailAddress string) {
	if err := page.Navigate(jiraConfig.JiraProfilePage + accountId); err != nil {
		log.Fatal("Failed to navigate:", err)
	}
	time.Sleep(5 * time.Second)
	domainParts := strings.Split(gmailConfig.DomainParts, ",")
	emailAddress = getEmailAddress(page, domainParts...)
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
