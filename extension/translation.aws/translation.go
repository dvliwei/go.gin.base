/**
 * @ClassName translation
 * @Description //TODO 
 * @Author liwei
 * @Date 2021/5/31 12:04
 * @Version go.translation.api V1.0
 **/

package translation_aws


const Translation_Text_Redis = "translation.text.redis."

/**
* @Title TranslationInput
* @Description:  翻译请求参数
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/5/31
**/
type TranslationInput struct {
	Message string
	TargetLanguageCode string
	SourceLanguageCode string
}

/**
* @Title TranslationOut
* @Description:  翻译结果
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/5/31
**/
type TranslationOut struct {
	TranslationResult string
}

/**
* @Title LanguageCode
* @Description:语言code列表
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/5/31
**/
type LanguageCode struct {
	LanguageCode string
	ZhCnName string
}