package openapiv2

import (
	_ "embed"
)

//go:embed v1/service.swagger.json
var V1ServiceSwaggerJSON []byte
