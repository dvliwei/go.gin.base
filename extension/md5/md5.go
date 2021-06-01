/**
 * @ClassName md5
 * @Description //TODO 
 * @Author liwei
 * @Date 2021/5/25 11:27
 * @Version go.translation.api V1.0
 **/

package md5

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"os"
)

// @Title StrMd5
// @Description 字符串md5加密
// @Return snowflake.ID
func StrMd5(str string) string{
	if str==""{
		return ""
	}
	h:=sha1.New()
	h.Write([]byte(str))
	hStr :=hex.EncodeToString(h.Sum(nil))
	//md5加盐
	nStr :=hStr+os.Getenv("AUTH_KEY")
	mh:=md5.New()
	mh.Write([]byte(nStr))
	return hex.EncodeToString(mh.Sum(nil))
}

/**
* @Title StrMd5NotConfKey
* @Description:
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/8
**/
func StrMd5NotConfKey(str string) string{
	if str==""{
		return ""
	}
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	return hex.EncodeToString(md5Ctx.Sum(nil))

}