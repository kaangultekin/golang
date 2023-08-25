package user

type IUserRepositoryForElasticsearch interface {
	SearchUsers(string) (interface{}, error)
}
