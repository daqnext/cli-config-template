package components

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/daqnext/cli-config-template/cli"
	elasticsearch "github.com/olivere/elastic/v7"
)

type ElasticSRetrier struct {
}

func (r *ElasticSRetrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	return 120 * time.Second, true, nil //retry after 2mins
}

/*
elasticsearch_addr
elasticsearch_username
*/

func InitElasticSearch() (*elasticsearch.Client, error) {

	elasticsearch_addr, elasticsearch_addr_err := cli.AppToDO.ConfigJson.GetString("elasticsearch_addr")
	if elasticsearch_addr_err != nil {
		return nil, errors.New("elasticsearch_addr [string] in config.json not defined," + elasticsearch_addr_err.Error())
	}

	elasticsearch_username, elasticsearch_username_err := cli.AppToDO.ConfigJson.GetString("elasticsearch_username")
	if elasticsearch_username_err != nil {
		return nil, errors.New("elasticsearch_username_err [string] in config.json not defined," + elasticsearch_username_err.Error())
	}

	elasticsearch_password, elasticsearch_password_err := cli.AppToDO.ConfigJson.GetString("elasticsearch_password")
	if elasticsearch_password_err != nil {
		return nil, errors.New("elasticsearch_password [string] in config.json not defined," + elasticsearch_password_err.Error())
	}

	ElasticSClient, err := elasticsearch.NewClient(
		elasticsearch.SetURL(elasticsearch_addr),
		elasticsearch.SetBasicAuth(elasticsearch_username, elasticsearch_password),
		elasticsearch.SetSniff(false),
		elasticsearch.SetHealthcheckInterval(30*time.Second),
		elasticsearch.SetRetrier(&ElasticSRetrier{}),
		elasticsearch.SetGzip(true),
	)
	if err != nil {
		return nil, err
	}

	return ElasticSClient, nil
}
