package ads

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hurbcom/google-ads-go/auth"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

type GoogleAdsClient struct {
	credentials     *oauth2.Config
	token           *oauth2.Token
	conn            *grpc.ClientConn
	developerToken  string
	loginCustomerID string
	ctx             context.Context
}

type GoogleAdsClientParams struct {
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	DeveloperToken  string `json:"developer_token"`
	RefreshToken    string `json:"refresh_token"`
	LoginCustomerID string `json:"login_customer_id"`
}

// NewClient creates a new client with specified credential params
func NewClient(params *GoogleAdsClientParams) (*GoogleAdsClient, error) {
	credentials := auth.NewCredentials(params.ClientID, params.ClientSecret)
	initialToken := auth.NewPartialToken(params.RefreshToken)

	c := &GoogleAdsClient{
		credentials:     credentials,
		token:           initialToken,
		developerToken:  params.DeveloperToken,
		loginCustomerID: params.LoginCustomerID,
	}

	newToken, err := auth.RefreshToken(c.credentials, c.token)
	if err != nil {
		return nil, err
	}
	c.token = newToken

	conn, ctx, err := auth.NewGrpcConnection(c.token, c.developerToken, c.loginCustomerID)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.ctx = ctx

	return c, nil
}

// NewClientFromStorage creates a new client instance from specified "google-ads.json" file
func NewClientFromStorage(filepath string) (*GoogleAdsClient, error) {
	params, err := ReadCredentialsFile(filepath)
	if err != nil {
		return nil, err
	}
	client, err := NewClient(params)
	return client, err
}

// ReadCredentialsFile reads a credentials JSON file and returns the exported config
func ReadCredentialsFile(filepath string) (*GoogleAdsClientParams, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var g *GoogleAdsClientParams
	if err := json.Unmarshal(file, &g); err != nil {
		return nil, err
	}

	return g, nil
}

// Conn returns a pointer to the clients gRPC connection
func (g *GoogleAdsClient) Conn() *grpc.ClientConn {
	return g.conn
}

// Context returns the context of the client
func (g *GoogleAdsClient) Context() context.Context {
	return g.ctx
}

// TokenIsValid returns a bool indicating if the generated access token is valid
func (g *GoogleAdsClient) TokenIsValid() bool {
	return g.token.Valid()
}
