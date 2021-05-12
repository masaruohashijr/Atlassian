# Atlassian

Dear reader,
I share an application developed to solve the following problem:

# Problem Definition
Migrate a Jira Cloud software project to the Jira Service Desk so that the company could save some of the Jira subscriptions for its customers monthly charged.

# Proposed solution
Develop an application that automates the entire migration process that consists of:
* Filter all unresolved issues from the old project in Jira Software (Cloud) and register them with the same title in Jira Service Management (Cloud).
* Filter all users of the old project's client group in Jira Software (Cloud) and register them with the same email address in the client group in Jira Service Management (Cloud).

# Restriction: 
By default, Jira REST Api does not allow consultation of email addresses. Jira users must personally (the administrator cannot) configure the visibility of their profile data, such as: email addresses, local times, locations and photos, making them public.

# Opportunity: 
Only through the browser, the administrator can still view the email addresses of Jira Cloud users. So it would be possible to think of a scraping of that information.

# Developed Product Components
- Integration between desktop application written in Go and Jira Software (Cloud) using Token api.
- Integration between desktop application written in Go and Jira Service Management (Cloud) using api Token.
- Gmail API service consumer client.
- Web crawler to open a page and navitage throgh authentication, verification code and scraping the email address.
- Fill in the email address and password of the Jira Admin.
- Wait for the verification code sent to the Admin inbox.
- Fill in the verification code into the confirmation form of Jira.
- Navigate to each user's profile by collecting the email address.
- Registration of old issues in the new Jira Service Desk project.
- Registration of the email addresses of the old users in the Jira Service Desk Client Group for the new project.
- Inactivation of these users of the old project in Jira Software (Cloud).
