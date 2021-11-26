package cowswap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gelfand/cowswap-go/types"
)

const (
	BaseURL = "https://protocol-mainnet.gnosis.io/api/v1/"
)

var errInvalidOrderKind = errors.New("invalid order kind")

var DefaultClient = NewClient("")

type Client struct {
	url    string
	client *http.Client
}

func NewClient(url string) *Client {
	if url == "" {
		url = BaseURL
	}

	return &Client{
		url: url,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

type Fee struct {
	ExpirationDate string
	Amount         string
}

func (c *Client) Fee(ctx context.Context, params types.FeeParams) (types.FeeResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"fee", nil)
	if err != nil {
		return types.FeeResponse{}, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("sellToken", params.SellToken)
	q.Add("buyToken", params.BuyToken)
	q.Add("amount", params.Amount)
	q.Add("kind", params.Kind)
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return types.FeeResponse{}, err
	}
	defer resp.Body.Close()

	var res types.FeeResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return types.FeeResponse{}, err
	}

	return res, nil
}

type CreateOrder struct {
	SellToken         string
	BuyToken          string
	Receiver          string
	SellAmount        string
	BuyAmount         string
	ValidTo           int64
	AppData           string
	FeeAmount         string
	Kind              string
	PartiallyFillable bool
	SellTokenBalance  string
	BuyTokenBalance   string
	SigningScheme     string
	Signature         string
	From              string
}

func (c *Client) CreateOrder(ctx context.Context, params types.OrderParams) (string, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.url+"orders", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res string
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res, nil
}

type OrderById struct {
	UID string
}

func (c *Client) OrderByID(ctx context.Context, uid string) (types.Order, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"orders/"+uid, nil)
	if err != nil {
		return types.Order{}, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return types.Order{}, err
	}
	defer resp.Body.Close()

	var res types.Order
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return types.Order{}, err
	}

	return res, nil
}

type OrdersParameters struct {
	Owner                      string
	SellToken                  string
	BuyToken                   string
	IncludeFullyExecuted       bool
	IncludeInvalidated         bool
	IncludeInsufficientBalance bool
	IncludePresignaturePending bool
	IncludeUnsupportedTokens   bool
	MinValidTo                 int64
}

func (c *Client) Orders(ctx context.Context, params types.OrdersParams) ([]types.Order, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"orders", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()

	if params.Owner != "" {
		q.Add("owner", params.Owner)
	}

	if params.SellToken != "" {
		q.Add("sellToken", params.SellToken)
	}

	if params.BuyToken != "" {
		q.Add("buyToken", params.BuyToken)
	}
	q.Add("includeFullyExecuted", strconv.FormatBool(params.IncludeFullyExecuted))
	q.Add("includeInvalidated", strconv.FormatBool(params.IncludeInvalidated))
	q.Add("includeInsufficientBalance", strconv.FormatBool(params.IncludeInsufficientBalance))
	q.Add("includePresignaturePending", strconv.FormatBool(params.IncludePresignaturePending))
	q.Add("includeUnsupportedTokens", strconv.FormatBool(params.IncludeUnsupportedTokens))
	q.Add("minValidTo", strconv.Itoa(int(params.MinValidTo)))
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res []types.Order
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func Orders(ctx context.Context, params types.OrdersParams) ([]types.Order, error) {
	return DefaultClient.Orders(ctx, params)
}
