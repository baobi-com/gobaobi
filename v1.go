package gobaobi

import "strings"

func (this *Baobi) Pairs() ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/pairs"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) Ticker(region string, coin string) ([]byte, error) {
	data := map[string]interface{}{}
	data["coin"] = strings.ToLower(coin)
	path := "/api/v1/ticker/region/" + strings.ToLower(region)
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) AllTicker() ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/allticker"
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) Orders(region string, coin string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	data["coin"] = strings.ToLower(coin)
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	path := "/api/v1/orders/region/" + strings.ToLower(region)
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) Depth(region string, coin string) ([]byte, error) {
	data := map[string]interface{}{}
	data["coin"] = strings.ToLower(coin)
	path := "/api/v1/depth/region/" + strings.ToLower(region)
	body, err := this.getRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) Balance() ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/balance"
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) TrustList(region string, coin string, orderType string, params ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/trust_list/region/" + strings.ToLower(region)
	data["coin"] = strings.ToLower(coin)
	data["type"] = strings.ToLower(orderType)
	if len(params) > 0 {
		for k, v := range params[0] {
			data[k] = v
		}
	}
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) TrustView(region string, coin string, orderId string) ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/trust_view/region/" + strings.ToLower(region)
	data["coin"] = strings.ToLower(coin)
	data["id"] = orderId
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) TrustCancel(region string, coin string, orderId string) ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/trust_cancel/region/" + strings.ToLower(region)
	data["coin"] = strings.ToLower(coin)
	data["id"] = orderId
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) TrustAdd(region string, coin string, orderType string, amount string, price string) ([]byte, error) {
	data := map[string]interface{}{}
	path := "/api/v1/trust_add/region/" + strings.ToLower(region)
	data["coin"] = strings.ToLower(coin)
	data["type"] = strings.ToLower(orderType)
	data["amount"] = amount
	data["price"] = price
	body, err := this.postRequest(path, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
