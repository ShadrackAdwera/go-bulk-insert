package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	db "github.com/ShadrackAdwera/go-bulk-insert/db/sqlc"
	"github.com/ShadrackAdwera/go-bulk-insert/worker"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type CovidFileData struct {
	FileData *multipart.FileHeader `form:"file" binding:"required"`
}

func (srv *Server) uploadFile(ctx *gin.Context) {
	var covidFile CovidFileData

	if err := ctx.ShouldBind(&covidFile); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
	}

	file, err := covidFile.FileData.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	defer file.Close()

	var covidData []worker.CovidData

	b, err := io.ReadAll(file)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	if err = json.Unmarshal(b, &covidData); err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	opts := []asynq.Option{
		asynq.MaxRetry(5),
	}

	err = srv.distro.DistributeData(ctx, &covidData, opts...)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("items found %d \n", len(covidData))})
}

type GetDataArgs struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (srv *Server) getData(ctx *gin.Context) {
	var params GetDataArgs

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, "Provide the query params")
		return
	}

	data, err := srv.store.ListCases(ctx, db.ListCasesParams{
		Limit:  params.PageSize,
		Offset: (params.PageID - 1) * params.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Records Found: %d", len(data)),
		"data":    data,
	})
}
