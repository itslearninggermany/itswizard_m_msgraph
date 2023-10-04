package itswizard_m_msgraph

import (
	"errors"
	"fmt"
	//"github.com/segmentio/objconv/json"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	P "github.com/yaegashi/msgraph.go/ptr"
)

/////////////////////  PERSONS   /////////////////////  PERSONS   /////////////////////    PERSONS   /////////////////////    PERSONS   /////////////////////    PERSONS   /////////////////////
/*
Response all Users
*/
func (p *AADAction) GetAllUsers() ([]msgraph.User, error) {
	r := p.graphClient.Users().Request()
	users, err := r.Get(p.ctx)
	return users, err
}

/*
Get one user with username
*/
func (p *AADAction) GetUserWithUsername(username string) (msgraph.User, error) {
	r := p.graphClient.Users().Request()
	r.Filter(fmt.Sprintf("userPrincipalName eq '%s'", username))
	users, err := r.Get(p.ctx)
	if err != nil {
		return msgraph.User{}, err
	}
	if len(users) < 1 {
		return msgraph.User{}, errors.New("User does not exist")
	}
	return users[0], nil
}

/*
Get one User with AzureUserID
*/
func (p *AADAction) GetUserWithID(azureUserID string) (msgraph.User, error) {
	r := p.graphClient.Users().ID(azureUserID).Request()
	user, err := r.Get(p.ctx)
	if err != nil {
		//New
		var tmpUser msgraph.User
		return tmpUser, err
		/*
			return *user, err
		*/
	}
	return *user, nil
}

/*
GetMemberGroups
*/
/*
Add a License to the user
*/
type GetMemberGroupsResponse struct {
	OdataContext string   `json:"@odata.context"`
	Value        []string `json:"value"`
}

/*
func (p *AADAction) GetMemberGroups(azureUserID string) (groupIds []string, err error) {
	reqObj := map[string]interface{}{
		"securityEnabledOnly": "false",
	}
	b, err := p.graphClient.Users().ID(azureUserID).Request().JSONRequestNordmann(p.ctx, "POST", "/getMemberGroups", reqObj)
	if err != nil {
		return groupIds, err
	}
	var x GetMemberGroupsResponse
	err = json.Unmarshal(b, &x)
	groupIds = x.Value
	return groupIds, err
}
*/

/*
Create a new user
*/
func (p *AADAction) CreateUser(firstname, lastname, profile, password, username, mailNick string, otherMails []string) (azureUserId string, err error) {

	newUser := msgraph.User{
		DisplayName:  P.String(fmt.Sprint(firstname, " ", lastname)),
		GivenName:    P.String(firstname),
		Surname:      P.String(lastname),
		MailNickname: P.String(mailNick),
		PasswordProfile: &msgraph.PasswordProfile{
			ForceChangePasswordNextSignIn: P.Bool(true),
			Password:                      P.String(password),
		},
		UserPrincipalName: P.String(fmt.Sprint(username)),
		AccountEnabled:    P.Bool(true),
		UsageLocation:     P.String("DE"),
		JobTitle:          P.String(profile),
		OtherMails:        otherMails,
	}

	u, err := p.graphClient.Users().Request().Add(p.ctx, &newUser)
	if err != nil {
		return azureUserId, err
	}
	return *u.ID, nil
}

/*
Add a License to the user
*/
func (p *AADAction) AddLicense(azureUserID string, sKUId string) error {

	guid := msgraph.UUID(sKUId)
	reqObj := map[string]interface{}{
		"addLicenses":    []msgraph.AssignedLicense{msgraph.AssignedLicense{SKUID: &guid}},
		"removeLicenses": []msgraph.UUID{},
	}
	err := p.graphClient.Users().ID(azureUserID).Request().JSONRequest(p.ctx, "POST", "/assignLicense", reqObj, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
Delete a User with username
*/
func (p *AADAction) DeleteUser(azureUserId string) error {
	return p.graphClient.Users().ID(azureUserId).Request().Delete(p.ctx)
}

/*
Update a User
*/
func (p *AADAction) UpdateUser(azureUserId, firstname, lastname, profile, syncID, username, domain string) error {
	//PATCH /users/{id | userPrincipalName}
	reqObj := map[string]interface{}{
		"DisplayName":       fmt.Sprint(firstname, " ", lastname),
		"GivenName":         firstname,
		"Surname":           lastname,
		"JobTitle":          profile,
		"EmployeeID":        syncID,
		"userPrincipalName": fmt.Sprint(username, "@", domain),
	}

	err := p.graphClient.Users().ID(azureUserId).Request().JSONRequest(p.ctx, "PATCH", "", reqObj, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
Update a User
*/
func (p *AADAction) UpdateUserOtherEmails(azureUserId string, otherMails []string) error {
	//PATCH /users/{id | userPrincipalName}
	reqObj := map[string]interface{}{
		"OtherMails": otherMails,
	}

	err := p.graphClient.Users().ID(azureUserId).Request().JSONRequest(p.ctx, "PATCH", "", reqObj, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
Update a User
*/
func (p *AADAction) UpdateUserPassword(username, newPassword string) error {
	//PATCH /users/{id | userPrincipalName}
	pw := msgraph.PasswordProfile{
		ForceChangePasswordNextSignIn: P.Bool(false),
		Password:                      P.String(newPassword),
	}

	reqObj := map[string]interface{}{
		"passwordProfile": pw,
	}

	err := p.graphClient.Users().ID(username).Request().JSONRequest(p.ctx, "PATCH", "", reqObj, nil)
	if err != nil {
		return err
	}
	return nil
}
