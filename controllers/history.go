package controllers

import (
	"backend/configs"
	"backend/serializers"
	"backend/services"
	"github.com/gin-gonic/gin"
)

//func GetRealtimeData(ctx *gin.Context)  {
//	result, err := services.Service.GetRealtimeData()
//	if err != nil {
//		returnMsg(ctx, 500, "", err.Error())
//		return
//	}
//
//	res := make([]map[string]interface{}, 0)
//	for _, v := range result {
//		valueMap := make(map[string]interface{})
//
//		valueMap["code"] = v.Tags["code"]
//		for i, key := range v.Columns {
//			valueMap[key] = v.Values[0][i]
//		}
//		res = append(res, valueMap)
//	}
//
//	returnMsg(ctx, 200, res, "success")
//}

func GetRealtimeData(ctx *gin.Context)  {
	res := make([]map[string]interface{}, 0)
	for _, v := range configs.RealtimeData {
		res = append(res, v)
	}

	returnMsg(ctx, 200, res, "success")
}

func GetRealtimeTemperature(ctx *gin.Context)  {
	res := make([]interface{}, 0)
	n := 0
	for _, v := range configs.RealtimeTemperature {
		n += 1
		if n > 20 {
			break
		}
		res = append(res, v)
	}

	returnMsg(ctx, 200, res, "success")
}

func GetHistoryData(ctx *gin.Context) {
	var filter serializers.HistoryRequest
	if err := ctx.ShouldBindJSON(&filter); err != nil {
		returnMsg(ctx, 500, "", err.Error())
		return
	}

	// 获取数据
	result, err := services.Service.GetHistoryData(filter)
	if err != nil {
		returnMsg(ctx, 500, "", err.Error())
		return
	}
	returnMsg(ctx, 200, result, "success")
}