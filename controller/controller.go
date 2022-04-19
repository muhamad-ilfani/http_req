package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http_req/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUrl(c *gin.Context) {
	var result []model.Data
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"Data": result,
	})
}

func GetUrlID(c *gin.Context) {
	var result model.Data
	var ID = c.Param("id")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"Data": result,
	})
}

func PostUrl(c *gin.Context) {
	var data model.Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	reqJson, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Contect-type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	json.Unmarshal(body, &data)

	c.JSON(http.StatusOK, gin.H{
		"Data": data,
	})
}
