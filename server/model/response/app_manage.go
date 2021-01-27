package response

import "gin-vue-admin/model"

type AppResponse struct {
	App *model.App `json:"app"`
}
