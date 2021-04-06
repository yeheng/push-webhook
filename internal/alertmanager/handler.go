package alertmanager

import (
	"bytes"
	"encoding/json"
	"github.com/YeHeng/push-webhook/app"
	common "github.com/YeHeng/push-webhook/common/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func routerHandler(c *gin.Context) {
	var notification Notification

	key := c.Query("key")
	err := c.BindJSON(&notification)

	bolB, _ := json.Marshal(notification)

	app.Logger.Infof("received alertmanager json: %s, robot key: %s", string(bolB), key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, e := alertManager2Wx(notification, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result.Message, "Code": result.Code})
}

func alertManager2Wx(notification Notification, defaultRobot string) (common.CommonResult, error) {

	markdown, robotURL, err := alertManagerToMarkdown(notification)

	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "marshal json fail " + err.Error(),
			},
			nil
	}

	var qywxRobotURL string

	if robotURL != "" {
		qywxRobotURL = robotURL
	} else {
		qywxRobotURL = defaultRobot
	}

	if len(qywxRobotURL) == 0 {
		return common.CommonResult{
				Code:    404,
				Message: "robot url is nil",
			},
			nil
	}

	req, err := http.NewRequest(
		"POST",
		qywxRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		return common.CommonResult{
				Code:    400,
				Message: "request robot url fail " + err.Error(),
			},
			nil
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return common.CommonResult{
				Code:    404,
				Message: "request wx api url fail " + err.Error(),
			},
			nil
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.Logger.Fatal(err)
	}
	bodyString := string(bodyBytes)
	app.Logger.Debugf("response: %s, header: %s", bodyString, resp.Header)

	return common.CommonResult{
		Code:    resp.StatusCode,
		Message: bodyString,
	}, nil
}
