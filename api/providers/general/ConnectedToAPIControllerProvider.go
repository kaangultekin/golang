package general

import (
	generalControllers "golang/api/controllers/general"
)

func ConnectedToAPIControllerProvider() *generalControllers.ConnectedToAPIController {
	return &generalControllers.ConnectedToAPIController{}
}
