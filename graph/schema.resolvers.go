package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pusher-Labs/midgard-graphql/graph/generated"
	"github.com/Pusher-Labs/midgard-graphql/graph/model"
)

func (r *queryResolver) Health(ctx context.Context) (*model.Health, error) {
	health := &model.Health{}

	resp, err := DefaultHttpClient.Get("http://18.158.69.134:8080/v1/health")
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&health)
	if err != nil {
		return nil, err
	}

	return health, nil
}

func (r *queryResolver) Stakers(ctx context.Context) ([]*string, error) {
	var stakers []*string

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/stakers", BaseUrl))
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&stakers)
	if err != nil {
		return nil, err
	}

	return stakers, nil
}

func (r *queryResolver) Staker(ctx context.Context, address string) (*model.Staker, error) {
	staker := &model.Staker{}

	fmt.Printf("address is %s", address)

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("http://18.158.69.134:8080/v1/stakers/%s", address))
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&staker)
	if err != nil {
		return nil, err
	}

	return staker, nil
}

func (r *queryResolver) Txs(ctx context.Context, address *string, txid *string, asset *string, typeArg []*string, offset int, limit int) (*model.Transactions, error) {
	transactions := &model.Transactions{}

	baseUrl := fmt.Sprintf("%s/txs", BaseUrl)

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	offsetString := strconv.Itoa(offset)
	limitString := strconv.Itoa(limit)

	q := req.URL.Query()
	q.Add("offset", offsetString)
	q.Add("limit", limitString)

	if address != nil {
		a := *address
		q.Add("address", a)
	}

	if txid != nil {
		t := *txid
		q.Add("address", t)
	}

	if asset != nil {
		a := *asset
		q.Add("asset", a)
	}

	if typeArg != nil {

		var args []string

		for _, transactionType := range typeArg {
			arg := *transactionType
			args = append(args, arg)
		}

		// args := *typeArg
		typesString := strings.Join(args[:], ",")
		q.Add("type", typesString)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Get(req.URL.String())
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *queryResolver) Assets(ctx context.Context, asset []*string) ([]*model.Asset, error) {
	var assets []string
	var body []*model.Asset

	for _, a := range asset {
		arg := *a
		assets = append(assets, arg)
	}

	assetString := strings.Join(assets[:], ",")

	baseUrl := fmt.Sprintf("%s/assets", BaseUrl)

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("asset", assetString)

	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Get(req.URL.String())
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) Pools(ctx context.Context) ([]*string, error) {
	var pools []*string

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/pools", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

func (r *queryResolver) PoolData(ctx context.Context, asset []*string) ([]*model.Pool, error) {
	var assets []string
	var body []*model.Pool

	for _, a := range asset {
		arg := *a
		assets = append(assets, arg)
	}

	assetString := strings.Join(assets[:], ",")

	baseUrl := fmt.Sprintf("%s/pools/detail", BaseUrl)

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("asset", assetString)

	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Get(req.URL.String())
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) StakerPoolData(ctx context.Context, address string, asset []*string) ([]*model.StakerPoolData, error) {
	var body []*model.StakerPoolData
	var assets []string

	for _, a := range asset {
		arg := *a
		assets = append(assets, arg)
	}

	assetString := strings.Join(assets[:], ",")

	baseUrl := fmt.Sprintf("%s/stakers/%s/pools", BaseUrl, address)

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("asset", assetString)

	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Get(req.URL.String())
	if err != nil {
		panic(fmt.Errorf("error fetching stuff request"))
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) PoolAddresses(ctx context.Context) (*model.PoolAddresses, error) {
	var body *model.PoolAddresses

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/thorchain/pool_addresses", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) Network(ctx context.Context) (*model.NetworkData, error) {
	var body *model.NetworkData

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/network", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) Nodes(ctx context.Context) ([]*model.Node, error) {
	var body []*model.Node

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/nodes", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) Constants(ctx context.Context) (*model.Constants, error) {
	var body *model.Constants

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/thorchain/constants", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *queryResolver) LastBlock(ctx context.Context) (*model.LastBlock, error) {
	var body *model.LastBlock

	resp, err := DefaultHttpClient.Get(fmt.Sprintf("%s/thorchain/lastblock", BaseUrl))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
