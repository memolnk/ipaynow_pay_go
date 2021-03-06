package ipaynow_pay

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type App struct {
	AppId  string
	AppKey string
	IsDev  bool
}
type OrderDetail struct {
	MhtOrderName   string
	MhtOrderDetail string
	MhtOrderAmt    int
	MhtGoodsTag    string
}

/**
 * 微信被扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param mhtSubAppId 微信子号对应多个公众号的时候必填,如果只对应一个公众号则不传
 * @param notifyUr 后台通知地址,详见2.2
 * @param channelAuthCode 支付码
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Wx_scan_05(app *App, orderDetail *OrderDetail, mhtSubAppId string, notifyUrl string, channelAuthCode string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, channelAuthCode, "", "", "05", mhtSubAppId, "", "", notifyUrl, "", "13", -1, mhtOrderNo)
}

/**
 * 支付宝被扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUr 后台通知地址
 * @param channelAuthCode 支付码
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Ali_scan_05(app *App, orderDetail *OrderDetail, notifyUrl string, channelAuthCode string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, channelAuthCode, "", "", "05", "", "", "", notifyUrl, "", "12", -1, mhtOrderNo)
}

/**
 * 手Q被扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUr 后台通知地址
 * @param channelAuthCode 支付码
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Handq_scan_05(app *App, orderDetail *OrderDetail, notifyUrl string, channelAuthCode string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, channelAuthCode, "", "", "05", "", "", "", notifyUrl, "", "25", -1, mhtOrderNo)
}

/**
 * 京东被扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUr 后台通知地址
 * @param channelAuthCode 支付码
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Jd_scan_05(app *App, orderDetail *OrderDetail, notifyUrl string, channelAuthCode string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, channelAuthCode, "", "", "05", "", "", "", notifyUrl, "", "04", -1, mhtOrderNo)
}

/**
 * 银联被扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUr 后台通知地址
 * @param channelAuthCode 支付码
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Union_scan_05(app *App, orderDetail *OrderDetail, notifyUrl string, channelAuthCode string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, channelAuthCode, "", "", "05", "", "", "", notifyUrl, "", "27", -1, mhtOrderNo)
}

/**
 * 微信主扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param mhtSubAppId 微信子号对应多个公众号的时候必填,如果只对应一个公众号则不传
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Wx_scan_08(app *App, orderDetail *OrderDetail, mhtSubAppId string, notifyUrl string, mhtOrderNo string) string {
	//最后参数0返回图片,data:..格式 。 1 返回支付链接
	return getPayResult(app, orderDetail, "", "", "", "05", mhtSubAppId, "", "", notifyUrl, "", "13", 0, mhtOrderNo)
}

/**
 * 支付宝主扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Ali_scan_08(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	//最后参数0返回图片,data:..格式 。 1 返回支付链接
	return getPayResult(app, orderDetail, "", "", "", "05", "", "", "", notifyUrl, "", "12", 0, mhtOrderNo)
}

/**
 * 手Q主扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Handq_scan_08(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	//最后参数0返回图片,data:..格式 。 1 返回支付链接
	return getPayResult(app, orderDetail, "", "", "", "05", "", "", "", notifyUrl, "", "25", 0, mhtOrderNo)
}

/**
 * 京东主扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Jd_scan_08(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	//最后参数0返回图片,data:..格式 。 1 返回支付链接
	return getPayResult(app, orderDetail, "", "", "", "05", "", "", "", notifyUrl, "", "04", 0, mhtOrderNo)
}

/**
 * 银联主扫支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Union_scan_08(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	//最后参数0返回图片,data:..格式 。 1 返回支付链接
	return getPayResult(app, orderDetail, "", "", "", "05", "", "", "", notifyUrl, "", "27", 0, mhtOrderNo)
}

/**
 * 微信公众号支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Wx_p_account(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	//最后参数为1返回支付要素
	return getPayResult(app, orderDetail, "", "", "", "0600", "", "", "", notifyUrl, frontNotifyUrl, "13", 0, mhtOrderNo)
}

/**
 * 支付宝公众号支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Ali_p_account(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	//最后参数为1返回支付要素
	return getPayResult(app, orderDetail, "", "", "", "0600", "", "", "", notifyUrl, frontNotifyUrl, "12", 0, mhtOrderNo)
}

/**
 * 手Q公众号支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Handq_p_account(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	//最后参数为1返回支付要素
	return getPayResult(app, orderDetail, "", "", "", "0600", "", "", "", notifyUrl, frontNotifyUrl, "25", 0, mhtOrderNo)
}

/**
 * 微信H5
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param consumerCreateIp 用户支付IP
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Wx_h5(app *App, orderDetail *OrderDetail, consumerCreateIp string, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", consumerCreateIp, "", "0601", "", "", "", notifyUrl, frontNotifyUrl, "13", 1, mhtOrderNo)
}

/**
 * 支付宝H5
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Ali_h5(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "0601", "", "", "", notifyUrl, frontNotifyUrl, "12", 1, mhtOrderNo)
}

/**
 * 银联H5
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Unionpay_h5(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "0601", "", "", "", notifyUrl, frontNotifyUrl, "20", 1, mhtOrderNo)
}

/**
 * 招行一网通H5
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Cmbywt_h5(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "0601", "", "", "", notifyUrl, frontNotifyUrl, "17", 1, mhtOrderNo)
}

/**
 * 手Q H5
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param frontNotifyUrl 前台页面跳转地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Handq_h5(app *App, orderDetail *OrderDetail, notifyUrl string, frontNotifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "0601", "", "", "", notifyUrl, frontNotifyUrl, "25", 1, mhtOrderNo)
}

/**
 * 支付宝网页web
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Ali_web(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "04", "", "", "", notifyUrl, "", "12", 0, mhtOrderNo)
}

/**
 * 银联网页web
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail 商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动)
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Union_web(app *App, orderDetail *OrderDetail, notifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "04", "", "", "", notifyUrl, "", "20", 0, mhtOrderNo)
}

/**
 * 微信小程序支付
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param orderDetail   商品名称,商品描述,商品价格(单位分),商品标记(用于营销活动),
 * @param consumerId  用户openId
 * @param notifyUrl 后台通知地址
 * @param mhtOrderNo 商户订单号,如果为空则自动生成商户订单号
 */
func Wx_app(app *App, orderDetail *OrderDetail, consumerId string, notifyUrl string, mhtOrderNo string) string {
	return getPayResult(app, orderDetail, "", "", "", "14", "", consumerId, "", notifyUrl, "", "13", 1, mhtOrderNo)
}

/**
* 商户被扫支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderScan05(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "05")
}

/**
* 商户主扫支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderScan08(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "08")
}

/**
* 商户公众号支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderPaccount(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "0600")
}

/**
* 商户H5支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderH5(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "0601")
}

/**
* 商户网页支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderWeb(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "04")
}

/**
* 商户微信小程序支付订单查询
* @param mhtOrderNo    商户订单号
* @param app appId(应用ID)和appKey ,
* 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
* @return
 */
func QueryOrderWxApp(mhtOrderNo string, app *App) string {
	return queryOrder(mhtOrderNo, app, "14")
}

/**
 * 商户支付订单查询
 * @param mhtOrderNo    商户订单号
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKey
 * @param deviceType    被扫05，主扫08，公众号传0600，h5传0601，网页04
 */
func queryOrder(mhtOrderNo string, app *App, deviceType string) string {
	var postMap = make(map[string]string)
	postMap["funcode"] = "MQ002"
	postMap["version"] = "1.0.0"
	postMap["deviceType"] = deviceType
	postMap["appId"] = app.AppId
	postMap["mhtOrderNo"] = mhtOrderNo
	postMap["mhtCharset"] = "UTF-8"
	postMap["mhtSignType"] = "MD5"

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]

	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/api/payment/"  
	} else {  
		posturl = "https://pay.ipaynow.cn/"  
	}  
	return post(posturl, content)
}

/**
 * 退款
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKeyey
 * @param mhtOrderNo    商户订单号
 * @param amount    退款金额(分)
 * @param reason    退款原因
 */
func RefundOrder(app *App, mhtOrderNo string, amount int, reason string) string {

	var postMap = make(map[string]string)

	postMap["funcode"] = "R001"
	postMap["appId"] = app.AppId
	postMap["mhtOrderNo"] = mhtOrderNo
	postMap["mhtRefundNo"] = getRandomString(20)
	postMap["amount"] = strconv.Itoa(amount)
	if reason != "" {
		postMap["reason"] = reason
	}
	postMap["mhtCharset"] = "UTF-8"

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature
	postMap["signType"] = "MD5"

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]

	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/refund_access/refundOrder"  
	} else {  
		posturl = "https://pay.ipaynow.cn/refund/refundOrder"  
	}  
	return post(posturl, content)
}

/**
 * 退款查询
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKeyey
 * @param mhtRefundNo   商户退款单号
 */
func RefundQuery(app *App, mhtRefundNo string) string {

	var postMap = make(map[string]string)

	postMap["funcode"] = "Q001"
	postMap["appId"] = app.AppId
	postMap["mhtRefundNo"] = mhtRefundNo
	postMap["mhtCharset"] = "UTF-8"

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature
	postMap["signType"] = "MD5"

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]

	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/refund_access/refundQuery"  
	} else {  
		posturl = "https://pay.ipaynow.cn/refund/refundQuery"  
	}  
	return post(posturl, content)
}

/**
 * 撤销(只能撤销当天的交易,且无论成功失败(逻辑包含退款))
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKeyey
 * @param mhtOrderNo    商户订单号
 * @param reason    退款原因
 */
func BackOrder(app *App, mhtOrderNo string, reason string) string {

	var postMap = make(map[string]string)

	postMap["funcode"] = "R002"
	postMap["appId"] = app.AppId
	postMap["mhtOrderNo"] = mhtOrderNo
	postMap["mhtRefundNo"] = getRandomString(20)
	if reason != "" {
		postMap["reason"] = reason
	}
	postMap["mhtCharset"] = "UTF-8"

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature
	postMap["signType"] = "MD5"

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]

	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/refund_access/refundOrder"  
	} else {  
		posturl = "https://pay.ipaynow.cn/refund/refundOrder"  
	}  
	return post(posturl, content)
}

/**
 * 撤销查询
 * @param app appId(应用ID)和appKey ,
 * 登录商户后台 : https://mch.ipaynow.cn ->商户中心->应用信息可以新增应用或查看appKeyey
 * @param mhtRefundNo   商户退款单号
 */
func BackQuery(app *App, mhtRefundNo string) string {

	var postMap = make(map[string]string)

	postMap["funcode"] = "Q002"
	postMap["appId"] = app.AppId
	postMap["mhtRefundNo"] = mhtRefundNo
	postMap["mhtCharset"] = "UTF-8"

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature
	postMap["signType"] = "MD5"

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]
	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/refund_access/refundQuery"  
	} else {  
		posturl = "https://pay.ipaynow.cn/refund/refundQuery"  
	}  
	return post(posturl, content)
}

func getPayResult(app *App, orderDetail *OrderDetail,
	channelAuthCode string, consumerCreateIp string, mhtSubMchId string, deviceType string, mhtSubAppId string, consumerId string, mhtReserved string, notifyUrl string, frontNotifyUrl string, payChannelType string, outputType int, mhtOrderNo string) string {

	var postMap = make(map[string]string)

	if channelAuthCode != "" {
		postMap["channelAuthCode"] = channelAuthCode
	}
	if consumerCreateIp != "" {
		postMap["consumerCreateIp"] = consumerCreateIp
	}
	postMap["funcode"] = "WP001"
	postMap["version"] = "1.0.0"
	postMap["mhtCurrencyType"] = "156"
	postMap["mhtOrderType"] = "01"
	postMap["mhtOrderTimeOut"] = "2000"
	postMap["mhtCharset"] = "UTF-8"
	postMap["mhtSignType"] = "MD5"
	postMap["mhtOrderStartTime"] = time.Now().Format("20060102150405")

	postMap["mhtLimitPay"] = "0"
	if outputType != -1 {
		postMap["outputType"] = strconv.Itoa(outputType)
	}
	if mhtSubMchId != "" {
		postMap["mhtSubMchId"] = mhtSubMchId
	}
	if orderDetail.MhtGoodsTag != "" {
		postMap["mhtGoodsTag"] = orderDetail.MhtGoodsTag
	}

	if mhtSubAppId != "" {
		postMap["mhtSubAppId"] = mhtSubAppId
	}
	if consumerId != "" {
		postMap["consumerId"] = consumerId
	}
	if mhtReserved != "" {
		postMap["mhtReserved"] = mhtReserved
	}

	postMap["appId"] = app.AppId

	if mhtOrderNo != "" {
		postMap["mhtOrderNo"] = mhtOrderNo
	} else {
		postMap["mhtOrderNo"] = getRandomString(13)
	}
	postMap["mhtOrderName"] = orderDetail.MhtOrderName
	postMap["mhtOrderAmt"] = strconv.Itoa(orderDetail.MhtOrderAmt)
	postMap["mhtOrderDetail"] = orderDetail.MhtOrderDetail
	if notifyUrl != "" {
		postMap["notifyUrl"] = notifyUrl
	}
	if frontNotifyUrl != "" {
		postMap["frontNotifyUrl"] = frontNotifyUrl
	}
	postMap["deviceType"] = deviceType
	postMap["payChannelType"] = payChannelType

	var keys []string
	for k := range postMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var postFormLinkReport = ""
	for _, k := range keys {
		postFormLinkReport += k + "=" + postMap[k] + "&"
	}
	postFormLinkReport = postFormLinkReport[0 : len(postFormLinkReport)-1]

	var mhtSignature = fmt.Sprintf("%x", md5.Sum([]byte(postFormLinkReport+"&"+fmt.Sprintf("%x", md5.Sum([]byte(app.AppKey))))))

	postMap["mhtSignature"] = mhtSignature
	postMap["appKey"] = app.AppKey

	var content = ""
	for km, vm := range postMap {

		l, e := url.Parse("?" + vm)
		if e != nil {
			fmt.Println(l, e)
		}

		content += km + "=" + l.Query().Encode()[0:len(l.Query().Encode())-1] + "&"

	}
	content = content[0 : len(content)-1]

	var posturl = ""  
	if app.IsDev {  
		posturl = "https://dby.ipaynow.cn/api/payment/"  
	} else {  
		posturl = "https://pay.ipaynow.cn/"  
	}  
	return post(posturl, content)
}

func post(url string, postcontent string) string {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(postcontent))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
