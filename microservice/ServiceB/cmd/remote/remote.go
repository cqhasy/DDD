package remote

import (
	"ServiceB/infrastructure/etcdx"
	"io"
	"net/http"
)

const BuyHandlerPath = "interface/ServiceA/BuyHandlerPath"

func BuyRemote() string {
	var re etcdx.PathRemote
	r, err := etcdx.FindService(BuyHandlerPath, re)
	if err != nil {
		return ""
	}
	return r.Path
}
func GetOtherServiceWithParams(p string, params ...string) ([]byte, error) {
	p = "http://" + p
	for _, v := range params {
		p = p + "/" + v
	}
	req, err := http.Get(p)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
