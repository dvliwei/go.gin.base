/**
 * @ClassName translationController
 * @Description //TODO 
 * @Author liwei
 * @Date 2021/5/31 11:57
 * @Version go.translation.api V1.0
 **/

package translationController

import (
	"github.com/gin-gonic/gin"
	"go.translation.api/extension/translation.aws"
	"go.translation.api/extension/translation.aws/translation.action"
	"go.translation.api/http/controller"
	"github.com/astaxie/beego/validation"
)

type TranslationController struct {
	controller.GinBaseController
}

type translationResult struct {
	Message string `json:"message"`
	SourceLanguageCode string `json:"source_language_code"`
	TargetLanguageCode string `json:"target_language_code,omitempty"`
	TranResult string `json:"tran_result"`
}

func (this *TranslationController)ActionTranslation(ctx *gin.Context)  {
	valid :=validation.Validation{}
	valid.Required(ctx.Request.PostFormValue("message"),"message").Message("message is null")
	valid.Required(ctx.Request.PostFormValue("target_language_code"),"target_language_code").Message("target_language_code is null")
	var datas []string
	if valid.HasErrors(){
		for _,err := range valid.Errors {
			datas = append(datas,err.Message)
		}
		ctx.JSON(200,this.HTTP_ERROR_WITH_DATA(datas))
		return
	}

	var tranInput translation_aws.TranslationInput
	tranInput.Message = this.GetString("message","",ctx)
	tranInput.TargetLanguageCode = this.GetString("target_language_code","",ctx)
	tranInput.SourceLanguageCode = this.GetString("source_language_code","",ctx)
	result ,err:= translation_action.TranslationAction(tranInput)
	if err!=nil{
		ctx.JSON(200,controller.HTTP_ERROR_TRANSLATION_FAIL)
		return
	}
	var data  translationResult
	data.Message = tranInput.Message
	data.TargetLanguageCode = tranInput.TargetLanguageCode
	data.SourceLanguageCode = tranInput.SourceLanguageCode
	data.TranResult = result.TranslationResult
	ctx.JSON(200, 	this.HTTP_SUCCESS_WITH_DATA(data))
	return

}

/** 
* @Title ActionTranslationList
* @Description: 语言list
* @Param:  
* @return:  
* @Author: liwei
* @Date: 2021/5/31 
**/
func (this *TranslationController)ActionTranslationList(ctx *gin.Context)  {
	data :=translation_action.TranslationCodeList()
	ctx.JSON(200, 	this.HTTP_SUCCESS_WITH_DATA(data))
	return
}