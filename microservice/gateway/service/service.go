package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway/infrastruce"
	"io"
	"log"
	"net/http"
)

func AService1(c context.Context, target string, id string, num string) (infrastruce.ResponseA, error) {
	queryURL := fmt.Sprintf("%s/order/%s/%s", target, id, num)
	resp2, err := http.Get(queryURL)
	if err != nil {
		log.Println("查询购买状态失败:", err)

		return infrastruce.ResponseA{}, err
	}
	defer resp2.Body.Close()
	body, err := io.ReadAll(resp2.Body)
	if err != nil {
		return infrastruce.ResponseA{}, err
	}
	var response infrastruce.ResponseA
	err = json.Unmarshal(body, &response)
	if err != nil {
		return infrastruce.ResponseA{}, err
	}
	return response, nil
}
func AService2(c context.Context, target string, id string, num string) (infrastruce.ResponseA, error) {
	queryURL := fmt.Sprintf("%s/order/%s/%s", target, id, num)
	resp2, err := http.Get(queryURL)
	if err != nil {
		log.Println("查询购买状态失败:", err)

		return infrastruce.ResponseA{}, err
	}
	defer resp2.Body.Close()
	body, err := io.ReadAll(resp2.Body)
	if err != nil {
		return infrastruce.ResponseA{}, err
	}
	var response infrastruce.ResponseA
	err = json.Unmarshal(body, &response)
	if err != nil {
		return infrastruce.ResponseA{}, err
	}
	return response, nil
}
func BService(c context.Context, r infrastruce.RequestB, target string) (infrastruce.ResponseB, error) {
	var response infrastruce.ResponseB
	data, err := json.Marshal(r)
	if err != nil {
		return response, fmt.Errorf("JSON 编码失败: %w", err)
	}
	req, err := http.NewRequestWithContext(c, http.MethodPost, target+"/order", bytes.NewBuffer(data))
	if err != nil {
		return response, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 发出请求
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return response, fmt.Errorf("请求服务B失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("读取响应失败: %w", err)
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf("解析响应失败: %w", err)
	}

	return response, nil

}
