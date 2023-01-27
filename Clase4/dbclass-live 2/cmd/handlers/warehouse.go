package handlers

import (
	"dbtest/internal/warehouse"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	sv warehouse.Service
}

func (wh *Warehouse) GetReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		var id int; var err error
		ids, ok := ctx.GetQuery("id")
		if ok {
			id, err = strconv.Atoi(ids)
			if err != nil {
				ctx.JSON(400, nil)
				return
			}
		}

		// process
		report, err := wh.sv.GetReport(id)
		if err != nil {
			ctx.JSON(500, nil)
			return
		}

		// response
		ctx.JSON(200, report)
	}
}