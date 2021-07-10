package gmailfilter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/oauth2"
	googleoauth "golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Config is the configuration structure used to instantiate the Google
// provider.
type Config struct {
	Credentials           string
	ImpersonatedUserEmail string

	client           *http.Client
	gmailService     *gmail.Service
	context          context.Context
	terraformVersion string
	userAgent        string
	scopes           []string

	tokenSource oauth2.TokenSource
}

var DefaultClientScopes = []string{
	gmail.GmailLabelsScope,
	gmail.GmailSettingsBasicScope,
}

func (c *Config) LoadAndValidate(ctx context.Context) error {
	if len(c.scopes) == 0 {
		c.scopes = DefaultClientScopes
	}
	tokenSource, err := c.getTokenSource(c.scopes)
	if err != nil {
		return err
	}
	c.tokenSource = tokenSource

	cleanCtx := context.WithValue(ctx, oauth2.HTTPClient, cleanhttp.DefaultClient())

	// 1. OAUTH2 TRANSPORT/CLIENT - sets up proper auth headers
	client := oauth2.NewClient(cleanCtx, tokenSource)

	// 2. Logging Transport - ensure we log HTTP requests to GCP APIs.
	loggingTransport := logging.NewTransport("Google", client.Transport)

	// 3. Retry Transport - retries common temporary errors
	// Keep order for wrapping logging so we log each retried request as well.
	// This value should be used if needed to create shallow copies with additional retry predicates.
	// See ClientWithAdditionalRetries
	retryTransport := NewTransportWithDefaultRetries(loggingTransport)

	// Set final transport value.
	client.Transport = retryTransport

	// This timeout is a timeout per HTTP request, not per logical operation.
	client.Timeout = 30 * time.Second

	tfUserAgent := terraformUserAgent(c.terraformVersion)
	providerVersion := fmt.Sprintf("terraform-provider-gmailfilter/%s", Version)
	userAgent := fmt.Sprintf("%s %s", tfUserAgent, providerVersion)

	c.client = client
	c.context = ctx
	c.userAgent = userAgent

	gmailService, err := gmail.NewService(ctx, option.WithHTTPClient(c.client))
	if err != nil {
		return nil
	}
	c.gmailService = gmailService
	return nil
}

func (c *Config) getTokenSource(clientScopes []string) (oauth2.TokenSource, error) {
	if c.Credentials != "" && c.ImpersonatedUserEmail != "" {
		contents, _, err := readFromPathOrContent(c.Credentials)
		if err != nil {
			return nil, fmt.Errorf("Error loading credentials: %s", err)
		}

		//creds, err := googleoauth.CredentialsFromJSON(context.Background(), []byte(contents), clientScopes...)
		//if err != nil {
		//	return nil, fmt.Errorf("Unable to parse credentials from '%s': %s", contents, err)
		//}
		//log.Printf("[INFO] Authenticating using configured Google JSON 'credentials'...")
		//log.Printf("[INFO]   -- Scopes: %s", clientScopes)
		// return creds.TokenSource, nil

		var serviceAccount serviceAccountFile
		if err := parseJSON(&serviceAccount, contents); err != nil {
			return nil, fmt.Errorf("error parsing credentials %q: %s", contents, err)
		}

		conf := jwt.Config{
			Email:      serviceAccount.ClientEmail,
			PrivateKey: []byte(serviceAccount.PrivateKey),
			Scopes:     clientScopes,
			TokenURL:   "https://oauth2.googleapis.com/token",
		}
		conf.Subject = c.ImpersonatedUserEmail
		return conf.TokenSource(context.Background()), nil
	}

	log.Printf("[INFO] Authenticating using DefaultClient...")
	log.Printf("[INFO]   -- Scopes: %s", clientScopes)
	return googleoauth.DefaultTokenSource(context.Background(), clientScopes...)
}

func readFromPathOrContent(poc string) (string, bool, error) {
	if len(poc) == 0 {
		return poc, false, nil
	}

	path := poc
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, true, err
		}
	}

	if _, err := os.Stat(path); err == nil {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return string(contents), true, err
		}
		return string(contents), true, nil
	}

	return poc, false, nil
}

type serviceAccountFile struct {
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientId     string `json:"client_id"`
}

func parseJSON(result interface{}, contents string) error {
	r := strings.NewReader(contents)
	dec := json.NewDecoder(r)

	return dec.Decode(result)
}
