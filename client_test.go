package cowswap

import (
	"context"
	"fmt"
	"testing"

	"github.com/gelfand/cowswap-go/types"
)

var client = NewClient("")

func TestClient_Fee(t *testing.T) {
	ctx := context.Background()
	_, err := client.Fee(ctx, types.FeeParams{
		SellToken: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		BuyToken:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Amount:    "100000000",
		Kind:      "buy",
	})
	if err != nil {
		t.Errorf("Client.Fee() error = %v, wantErr = false", err)
	}
}

func TestClient_Orders(t *testing.T) {
	ctx := context.Background()
	resp, err := client.Orders(ctx, types.OrdersParams{
		Owner:                      "0x7BB8536Bac7d6e5c4E352Ce50be8968d5d6cd445",
		IncludeFullyExecuted:       true,
		IncludeInvalidated:         true,
		IncludeInsufficientBalance: true,
		IncludePresignaturePending: true,
		IncludeUnsupportedTokens:   true,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}

func TestOrders(t *testing.T) {
	ctx := context.Background()
	_, err := Orders(ctx, types.OrdersParams{
		Owner:                      "0x7BB8536Bac7d6e5c4E352Ce50be8968d5d6cd445",
		IncludeFullyExecuted:       true,
		IncludeInvalidated:         true,
		IncludeInsufficientBalance: true,
		IncludePresignaturePending: true,
		IncludeUnsupportedTokens:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestOrderByID(t *testing.T) {
	ctx := context.Background()

	resp, err := client.OrderByID(ctx, "0x54d77f468bfd87405000734b689e1549d54e4106e39db9c2a2ed4793493b5c107bb8536bac7d6e5c4e352ce50be8968d5d6cd44561906a59")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
