/**
 * @ClassName translation_action
 * @Description //TODO 翻译
 * @Author liwei
 * @Date 2021/5/31 12:03
 * @Version go.translation.api V1.0
 **/

package translation_action

import (
	"github.com/aws/aws-sdk-go/service/translate"
	"go.translation.api/extension/dbLog"
	"go.translation.api/extension/md5"
	"go.translation.api/extension/redis"
	"go.translation.api/extension/translation.aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

var translationCacheKey string

var fTranInput translation_aws.TranslationInput

/**
* @Title TranslationAction
* @Description:  文本翻译
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/5/31
**/
func TranslationAction(tranInput  translation_aws.TranslationInput)(translation_aws.TranslationOut,error){
	var result translation_aws.TranslationOut
	translationCacheKey = os.Getenv("REDIS_PREFIX")+ md5.StrMd5(translation_aws.Translation_Text_Redis+tranInput.Message+tranInput.TargetLanguageCode)
	if !redis.HasRedis(translationCacheKey){
		fTranInput = tranInput
		resultStr,err:=awsTranslation()
		if err!=nil{
			dbLog.ServerErrorMap("aws翻译失败", map[string]interface{}{
				"tranInput":tranInput,
				"errr":err,
				"resultStr":resultStr,
			})
			redis.SetDataToRedisWithSecond(translationCacheKey,resultStr,86400*90)
		}
	}
	resultStr,err :=redis.GetRedisDataByKey(translationCacheKey)
	if err!=nil{
		dbLog.ServerErrorMap("redis获取翻译失败", map[string]interface{}{
			"tranInput":tranInput,
			"errr":err,
			"resultStr":resultStr,
			"translationCacheKey":translationCacheKey,
		})
		return result,err
	}
	result.TranslationResult =  resultStr
	return  result,nil
}

var translateSession *translate.Translate

/**
* @Title awsTranslation
* @Description:  aws翻译
* @Param:
* @return:
* @Author: liwei
* @Date: 2021/5/31
**/
func awsTranslation()(string,error)  {
	access_key:=os.Getenv("AWS_ACCESS_KEY")
	secret_key:=os.Getenv("AWS_SECRET_KEY")
	translateSession = translate.New(session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Region: aws.String("us-east-1"), // Frankfurt
	})))
	sourceLanguageCode:=fTranInput.SourceLanguageCode
	if sourceLanguageCode==""{
		sourceLanguageCode="auto"
	}
	targetLanguageCode:=fTranInput.TargetLanguageCode
	text:=fTranInput.Message
	response, err := translateSession.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(sourceLanguageCode),
		TargetLanguageCode: aws.String(targetLanguageCode),
		Text: aws.String(text),
	})
	if err != nil {
		return "",err
	}
	resultStr:=*response.TranslatedText //翻译文本结果
	return resultStr,nil
}



func TranslationCodeList()(map[string]string){
	var result map[string]string
	result["af"]="南非荷兰语"
	result["sq"]="阿尔巴尼亚语"
	result["am"]="阿姆哈拉语"
	result["ar"]="阿拉伯语"
	result["HY"]="亚美尼亚"
	result["az"]="阿塞拜疆语"
	result["bn"]="孟加拉语"
	result["bs"]="波斯尼亚语"
	result["ca"]="加泰罗尼亚语"
	result["zh"]="简体中文"
	result["zh-TW"]="繁体中文"
	result["cs"]="捷克语"
	result["da"]="丹麦语"
	result["fa-AF"]="达里语"
	result["nl"]="荷兰语"
	result["en"]="英语"
	result["et"]="爱沙尼亚语"
	result["et"]="荷兰语"
	result["fa"]="波斯语"
	result["tl"]="菲律宾塔加洛语"
	result["fi"]="芬兰语"
	result["fr"]="法语"
	result["fr-CA"]="法语（加拿大）"
	result["ka"]="格鲁吉亚语"
	result["de"]="德语"
	result["el"]="希腊语"
	result["gu"]="古吉拉特"
	result["ht"]="海地克里奥尔"
	result["ha"]="豪萨语"
	result["he"]="希伯来语"
	result["hi"]="印地语"
	result["hu"]="匈牙利语"
	result["is"]="冰岛语"
	result["id"]="印度尼西亚语"
	result["it"]="意大利语"
	result["ja"]="日语"
	result["kn"]="卡纳达"
	result["is"]="哈萨克斯坦"
	result["kk"]="冰岛语"
	result["ko"]="韩语"
	result["lv"]="拉脱维亚语"
	result["lt"]="立陶宛"
	result["mk"]="马其顿"
	result["ms"]="马来语"
	result["ml"]="马拉雅拉姆"
	result["mt"]="马耳他语"
	result["mn"]="蒙古"
	result["no"]="挪威语"
	result["fa"]="波斯语"
	result["ps"]="普什图语"
	result["pl"]="波兰语"
	result["pt"]="葡萄牙语"
	result["ro"]="罗马尼亚语"
	result["ru"]="俄语"
	result["sr"]="塞尔维亚语"
	result["si"]="僧伽罗语"
	result["sk"]="斯洛伐克语"
	result["sl"]="斯洛文尼亚语"
	result["so"]="索马里语"
	result["es"]="西班牙语"
	result["es-MX"]="西班牙语（墨西哥）"
	result["sw"]="斯瓦希里语"
	result["sv"]="瑞典语"
	result["tl"]="塔加洛语"
	result["ta"]="泰米尔语"
	result["te"]="泰卢固语"
	result["th"]="泰语"
	result["tr"]="土耳其语"
	result["uk"]="乌克兰语"
	result["ur"]="乌尔都语"
	result["uz"]="乌兹别克"
	result["vi"]="越南语"
	result["cy"]="威尔士语"

	return result
}