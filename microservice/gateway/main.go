package main

import (
	"context"
	"gateway/infrastruce"
	"gateway/service"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net/http"
	"time"
)

func main() {

	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://60.205.12.92:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer client.Close()

	r := gin.Default()

	// 网关接口，转发到 ServiceA A
	r.GET("/serviceA/info/:id/:num", func(c *gin.Context) {
		id := c.Param("id")
		num := c.Param("num")
		resp, _ := client.Get(context.Background(), "/services/serviceA/info")
		if resp.Kvs == nil {
			log.Println("配置加载出错")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "配置加载出错"})
			return
		}
		target := string(resp.Kvs[0].Value)
		re, err := service.AService2(c, target, id, num)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, re)
	})
	r.GET("/serviceA/buy/:id/:num", func(c *gin.Context) {
		id := c.Param("id")
		num := c.Param("num")
		resp, _ := client.Get(context.Background(), "/services/serviceA/buy")
		if resp.Kvs == nil {
			log.Println("配置加载出错")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "配置加载出错"})
			return
		}
		target := string(resp.Kvs[0].Value)
		re, err := service.AService1(c, target, id, num)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, re)
	})

	// 网关接口，转发到 ServiceA B
	r.POST("/serviceB", func(c *gin.Context) {
		resp, _ := client.Get(context.Background(), "/services/serviceB")
		if resp.Kvs == nil {
			log.Println("配置加载出错")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "配置加载出错"})
			return
		}
		target := string(resp.Kvs[0].Value)
		var req infrastruce.RequestB
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "解析请求失败"})
		}
		re, err := service.BService(c, req, target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, re)

	})

	r.Run(":8082")
}
