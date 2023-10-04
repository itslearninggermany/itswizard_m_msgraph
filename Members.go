package itswizard_m_msgraph

import (
	"fmt"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	"io/ioutil"
	"net/http"
)

/////////////////////   Membership   /////////////////////   Membership   /////////////////////   Membership   /////////////////////   Membership   /////////////////////   Membership
/*
TODO: Return all members from a group
*/

type Tmp struct {
	OdataContext  string         `json:"@odata.context"`
	OdataNextLink string         `json:"@odata.nextLink"`
	OdataType     string         `json:"@odata.type"`
	ID            string         `json:"id"`
	Value         []msgraph.User `json:"value"`
}

func (p *AADAction) GetAllMembersOfAGroup(azureGroupID string) ([]msgraph.User, error) {
	r := p.graphClient.Groups().ID(azureGroupID).Members().Request()
	var out Tmp
	err := r.JSONRequest(p.ctx, "GET", "", nil, &out)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(len(out.Value),UnPtrString(out.Value[0].DisplayName))
	//	fmt.Println(out.OdataNextLink)

	q, err := http.NewRequest("GET", out.OdataNextLink, nil)

	client := &http.Client{}
	resp, err := client.Do(q)
	fmt.Println(err)

	//	fmt.Println(resp.Status)
	//	fmt.Println(resp.Header)

	defer resp.Body.Close()
	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bodys))
	return out.Value, nil
}

/*
Add a User to a Group
*/
func (p *AADAction) AddMemberToAGroup(azureGroupId string, azureUserID string) error {
	reqObj := map[string]interface{}{
		"@odata.id": p.graphClient.DirectoryObjects().ID(azureUserID).Request().URL(),
	}
	r := p.graphClient.Groups().ID(azureGroupId).Members().Request()
	err := r.JSONRequest(p.ctx, "POST", "/$ref", reqObj, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
Delete a Member from a Group
*/
func (p *AADAction) DeleteMemberFromAGroup(azureGroupID string, azureUserID string) error {
	r := p.graphClient.Groups().ID(azureGroupID).Members().ID(azureUserID).Request()
	err := r.JSONRequest(p.ctx, "DELETE", "/$ref", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
