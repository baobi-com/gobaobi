# gobaobi
*基于Baobi官方API文档的Golang实现*

## 官方API文档地址
> https://baobiex.github.io/apidoc/

## 安装
```shell script
go get -u https://github.com/baobi-com/gobaobi
```


## 函数列表
* 设置BaseUrl
    * SetBaseUrl(url string) //无返回值

* 获取所有交易对 Pairs()
* 获取交易对当前行情 Ticker(region string, coin string)
* 获取所有交易对当前行情 AllTicker()
* 获取最新成交单 Orders(region string, coin string, params ...map[string]interface{})
* 获取深度 Depth(region string, coin string)
* 获取账户资产信息 Balance()
* 挂单 TrustAdd(region string, coin string, orderType string, amount string, price string)
* 查看订单详情 TrustView(region string, coin string, orderId string)
* 挂单列表 TrustList(region string, coin string, orderType string, params ...map[string]interface{})
* 取消挂单 TrustCancel(region string, coin string, orderId string)

## 函数返回
***函数返回类型为 ([]byte, error)***

## Demo
```golang

package main

import (
	"github.com/baobi-com/gobaobi"
	"fmt"
)

func main() {

	accessKey := ""
	secertKey := ""

	//var baobi = gobaobi.Baobi{AccessKey: accessKey, SecertKey: secertKey, BaseUrl: "http://api.baobi.com"} // BaseUrl为可选，已内置接口地址。
	var baobi = gobaobi.Baobi{AccessKey: accessKey, SecertKey: secertKey}

	// 获取所有交易对
	res, err := baobi.Pairs()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取交易对当前行情
	res, err = baobi.Ticker("usdt", "btc")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取所有交易对当前行情
	res, err = baobi.AllTicker()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取最新成交单
	res, err = baobi.Orders("USDT", "btc")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取最新成交单 带since可选参数
	res, err = baobi.Orders("USDT", "BTC", map[string]interface{}{"since": "11212"}) // since参数为可选
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取深度
	res, err = baobi.Depth("usdt", "btc")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 获取账户资产信息
	res, err = baobi.Balance()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 挂单
	res, err = baobi.TrustAdd("usdt", "btc", "sell", "0.01", "20000")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 查看订单详情
	res, err = baobi.TrustView("usdt", "btc", "378775677")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 挂单列表
	res, err = baobi.TrustList("usdt", "btc", "open", map[string]interface{}{"since": "0"}) // since 为可选
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// 取消挂单
	res, err = baobi.TrustCancel("usdt", "btc", "378775677")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

}

```