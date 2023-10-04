package itswizard_m_msgraph

import (
	"fmt"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	P "github.com/yaegashi/msgraph.go/ptr"
	"log"
	"strings"
)

/////////////////////  GROUPS   /////////////////////    GROUPS   /////////////////////    GROUPS   /////////////////////    GROUPS   /////////////////////    GROUPS   /////////////////////

/*
Todo: Create Group
*/
func (p *AADAction) CreateGroup(groupname, synckey, schoolname string) (azureGroupId string, err error) {
	azureGroupName := fmt.Sprint(schoolname, " ", groupname)
	azureGroupNameMail := fmt.Sprint(schoolname, "_", groupname)
	// Default group: you should replace it with your own
	var defaultGroup = &msgraph.Group{
		Description:                 P.String(synckey),
		DisplayName:                 P.String(azureGroupName),
		MailEnabled:                 P.Bool(true),
		MailNickname:                P.String(Normalise(strings.ToLower(strings.Replace(azureGroupNameMail, " ", "", -1)))),
		SecurityEnabled:             P.Bool(false),
		Visibility:                  P.String("Private"),
		ResourceProvisioningOptions: []string{"Team"},
		GroupTypes:                  []string{"Unified"},
	}
	g, err := p.graphClient.Groups().Request().Add(p.ctx, defaultGroup)
	if err != nil {
		return "", err
	}
	return *g.ID, err
}

/*
Delete a Group
*/
func (p *AADAction) DeleteGroup(azureGroupId string) error {
	return p.graphClient.Groups().ID(azureGroupId).Request().Delete(p.ctx)
}

/*
Show Groups
*/
func (p *AADAction) ShowGroups() ([]msgraph.Group, error) {
	r := p.graphClient.Groups().Request()
	return r.Get(p.ctx)
}

/*
Show Group
*/
func (p *AADAction) ShowGroup(azureGroupId string) (group msgraph.Group, err error) {
	r := p.graphClient.Groups().ID(azureGroupId).Request()
	out, err := r.Get(p.ctx)
	if err != nil {
		log.Fatal(err)
	}
	group = *out
	return
}
