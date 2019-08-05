package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Redirect() {
	this.Data["json"] = map[string]interface{}{"msg":"Token失效或无效"}
	this.ServeJSON()
}

func (this *MainController) VerifyToken() {
	this.Data["json"] = "Token验证成功"
	this.ServeJSON()
}

func (this *MainController)GetToken()  {
	keys:=make(map[string]interface{})
	data := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &keys);err!=nil{
		this.Data["json"]=map[string]interface{}{"msg":"jsonErr"}
	}

	//1. aud 标识token的接收者.
	//2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
	//3. jti 是自定义的id号
	//4. iat 签名发行时间.
	//5. iss 是签名的发行者.
	//6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
	//7. sub 签名面向的用户
	//userInfo:=make(map[string]interface{})
	if keys["name"].(string)=="admin"{
		//userInfo["exp"]=strconv.FormatInt(time.Now().Unix()+60480,10)// 1周 exp 过期时间.通常与Unix UTC时间做对比过期后token无效604800
		// Add returns the time t+d.
		//userInfo["exp"]=time.Now().Add(time.Duration(time.Second*20)).Unix()	//exp 过期时间.通常与Unix UTC时间做对比过期后token无效
		//userInfo["iat"]=time.Now().Unix()	//iat 签名发行时间.
		//userInfo["sub"]="lpgsy"		//sub 签名面向的用户

		token:=this.createToken(beego.AppConfig.String("jwtkey"))
		this.Data["json"]=map[string]interface{}{"Authorization":token,"msg":beego.AppConfig.String("jwtkey")}
	}else {
		this.Data["json"]=map[string]interface{}{"Error":"身份验证未通过，不能签发Token","msg":"fail"}
	}
	this.ServeJSON()

}
//h1, _ := time.ParseDuration("-10h10m")
//eighteenBefore:=time.Now().Add(h1)	//18小时10分钟之前
func (this *MainController) createToken(key string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	//h, _ := time.ParseDuration("1s")
	//days:=time.Now().Add(20*h)
	claims:=jwt.MapClaims{	//创建 MapClaims map[string]interface{}
		// Add returns the time t+d.
		//"exp":strconv.FormatInt(time.Now().Unix()+60480,10),// 1周 exp 过期时间.通常与Unix UTC时间做对比过期后token无效604800
		//"exp":strconv.FormatInt(days.Unix(),10),
		"exp":strconv.FormatInt(time.Now().AddDate(0,0,7).Unix(),10),	//exp 过期时间.通常与Unix UTC时间做对比过期后token无效
		"iat":time.Now().Unix(),	//iat 签名发行时间.
		"sub":"BQP",	//sub 签名面向的用户
	}
	//claims:=make(jwt.MapClaims)  //创建 MapClaims map[string]interface{}
	//for index,value:=range userInfo{
	//	claims[index]=value
	//}

	token.Claims=claims
	signedToken, _ := token.SignedString([]byte(key))
	return signedToken
}

//*************************************************************************************************//
func CreateTokens()(tokenss string,err error){
	//自定义claim
	claim := jwt.MapClaims{
		"nbf":      int64(time.Now().Unix() - 1000),
		"exp":      int64(time.Now().Unix() + 3600),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err  = token.SignedString([]byte(viper.GetString("jwtkey")))
	return
}
