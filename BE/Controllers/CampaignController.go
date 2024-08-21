package Controllers

import (
	"net/http"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetKolsController(db *gorm.DB) func(*gin.Context) {
	return func(context *gin.Context) {
		var KolsVM ViewModels.KolViewModel
		var guid = uuid.New().String()

		// * Get Kols from the database based on the range of pageIndex and pageSize
		// * TODO: Implement the logic to get parameters from the request
		// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
		// @params: pageIndex
		// @params: pageSize
		var data DTO.GetSearchParam
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		// * Perform Logic Here
		// ! Pass the parameters to the Logic Layer
		kols, error := Logic.GetKolLogic(&data)
		if error != nil {
			KolsVM.Result = Const.UnSuccess
			KolsVM.ErrorMessage = error.Error()
			KolsVM.PageIndex = data.PageIndex // * change this to the actual page index from the request
			KolsVM.PageSize = data.PageSize   // * change this to the actual page size from the request
			KolsVM.Guid = guid
			context.JSON(http.StatusInternalServerError, KolsVM)
			return
		}

		// * Return the response after the logic is executed
		// ? If the logic is successful, return the response with HTTP Status OK (200)
		KolsVM.Result = Const.Success
		KolsVM.ErrorMessage = ""
		KolsVM.PageIndex = data.PageIndex // * change this to the actual page index from the request
		KolsVM.PageSize = data.PageSize   // * change this to the actual page size from the request
		KolsVM.Guid = guid
		KolsVM.KOL = kols
		KolsVM.TotalCount = int64(len(kols))
		context.JSON(http.StatusOK, KolsVM)
	}
}
