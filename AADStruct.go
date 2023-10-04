package itswizard_m_msgraph

import (
	"context"
	"github.com/jinzhu/gorm"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"regexp"
)

type AADAction struct {
	tentantID    *string
	clientID     *string
	clientSecret *string
	httpClient   *http.Client
	ctx          context.Context
	graphClient  *msgraph.GraphServiceRequestBuilder
	domain       string
}

/*
Creates an Connector to Azure Active Directory
*/
func NewAADAction(tentantID, clientID, clientSecret string) *AADAction {
	p := new(AADAction)
	//	flag.StringVar(p.tentantID, "tenant-id", tentantID, "Tenant ID")
	//	flag.StringVar(p.clientID, "client-id", clientID, "Client ID")
	//	flag.StringVar(p.clientSecret, "client-secret", clientSecret, "Client Secret")
	//	flag.Parse()
	p.ctx = context.Background()
	m := msauth.NewManager()
	scopes := []string{msauth.DefaultMSGraphScope}
	ts, err := m.ClientCredentialsGrant(p.ctx, tentantID, clientID, clientSecret, scopes)
	if err != nil {
		log.Fatal(err)
	}

	p.httpClient = oauth2.NewClient(p.ctx, ts)
	p.graphClient = msgraph.NewClient(p.httpClient)
	return p
}

/*
change äöüß
*/
func Normalise(input string) string {
	regExp := regexp.MustCompile("[äöüß]")

	umlauts := map[string]string{
		"ä": "ae",
		"ö": "oe",
		"ü": "ue",
		"ß": "ss",
	}

	replaceUmlauts := func(str string) string {
		return umlauts[str]
	}

	return regExp.ReplaceAllStringFunc(input, replaceUmlauts)
}

/*
Export an interface to a json
*/
func Dump(o interface{}) {
	enc := jsonx.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

type AadErrorLog struct {
	gorm.Model
	Type           string `gorm:"not null; type:VARCHAR(500)"`
	OrganisationID uint
	InstiutionID   uint
	Error          string `gorm:"not null; type:VARCHAR(500)"`
}

type AadLog struct {
	gorm.Model
	Type           string `gorm:"not null; type:VARCHAR(500)"`
	OrganisationID uint
	InstiutionID   uint
}
