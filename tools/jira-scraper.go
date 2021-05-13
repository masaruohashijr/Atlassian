package tools

import (
	"Atlassian/config"
	"Atlassian/models"
	"log"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

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
	println(navTo)
	if err := page.Navigate(navTo); err != nil {
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
