// internal/infra/container/file_container.go
package container

import (
	"project_virtual_internship_evermos/internal/package/controller"
)

func InitFileController() *controller.FileController {
	return controller.NewFileController()
}
