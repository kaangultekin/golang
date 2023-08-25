package user

import (
	"encoding/json"
	"errors"
	"golang/api/config"
	messageConstants "golang/api/constants/message"
	userModels "golang/api/models/user"
	resultStructs "golang/api/structs/result"
	"strings"
)

type UserRepositoryForElasticsearch struct{}

func (ur *UserRepositoryForElasticsearch) SearchUsers(query string) (interface{}, error) {
	var (
		resultForElasticsearch *resultStructs.ResultStructForElasticSearch
		users                  []userModels.User
	)

	jsonParams, jsonParamsErr := json.Marshal(map[string]interface{}{
		"size": 100,
		"sort": []map[string]interface{}{
			{
				"id": "desc",
			},
		},
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": query,
				"fields": []string{
					"name^3",
					"surname",
					"email",
				},
				"fuzziness": "AUTO",
				"operator":  "or",
			},
		},
	})

	if jsonParamsErr != nil {
		return nil, jsonParamsErr
	}

	res, err := config.ES.Search(
		config.ES.Search.WithIndex("users_index"),
		config.ES.Search.WithBody(
			strings.NewReader(
				string(jsonParams),
			),
		),
	)

	if err != nil {
		return nil, err
	}

	if resBodyErr := json.NewDecoder(res.Body).Decode(&resultForElasticsearch); resBodyErr != nil {
		return nil, resBodyErr
	}

	if resultForElasticsearch.Hits.Total.Value == 0 {
		return nil, errors.New(messageConstants.ErrUserNotFound)
	}

	for _, hit := range resultForElasticsearch.Hits.Hits {
		var user userModels.User

		sourceData, sourceDataErr := hit.Source.(map[string]interface{})

		if !sourceDataErr {
			continue
		}

		sourceJSON, sourceJSONErr := json.Marshal(sourceData)

		if sourceJSONErr != nil {
			continue
		}

		unmarshalSourceJson := json.Unmarshal(sourceJSON, &user)

		if unmarshalSourceJson != nil {
			continue
		}

		users = append(users, user)
	}

	return users, nil
}
