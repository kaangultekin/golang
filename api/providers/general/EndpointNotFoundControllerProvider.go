package general

import (
	generalControllers "golang/api/controllers/general"
)

func EndpointNotFoundControllerProvider() *generalControllers.EndpointNotFoundController {
	return &generalControllers.EndpointNotFoundController{}
}
