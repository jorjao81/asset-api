package service

import (
	"net/http"

	"github.com/unrolled/render"
)

type Asset struct {
	Code string `json:"code"`
}

func createGetAssetsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		assets := []Asset{
			Asset{
				Code: "This is a test",
			},
		}

		formatter.JSON(w, http.StatusOK, assets)
	}
}
