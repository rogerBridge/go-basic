package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 对请求的方法做处理
func Analysis(w http.ResponseWriter, r *http.Request) {
	// 判断请求方法是否合法?
	if r.Method == "POST" && r.Header.Get("Authorization") == "fe8db6e4-8f8c-4cc8-8e7d-0f7938a91095" {
		// 对传入的body进行解析
		v, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		// 尽量保持每秒一次的采集精度
		//{
		//	"vcuNo": "VCU1008610000",
		//	"timeRange": ["1573043627', "1573643698"],
		//  "gpsDataList": {
		//  "1573043627": [30.123432, 119.123456, 3.3454],
		//  "1573043632": [30.123433, 119.233432, 5,2345]
		//  }
		//}
		type phoneUploadGpsData struct {
			VcuNo       string               `json:"vcuNo"`
			TimeRange   []string             `json:"timeRange"`
			GpsDataList map[string][]float64 `json:"gpsDataList"` // 分别是纬度, 经度, 速度
		}
		var input phoneUploadGpsData
		err = json.Unmarshal(v, &input)
		if err != nil {
			log.Println(err, "转换[]byte to struct error!")
			http.Error(w, "转换[]byte to struct error!, 请查看上报的数据格式是否正确", 400)
		} else {
			// 开始对传入的数据进行处理咯~
			// 1. 拿到后台存储的ec数据(根据: 时间戳, VCU编号), 将拿到的数据加工成: GpsDataList 类型的数据
			//getEcData()
			// 2. 对比相同时间戳的latitude, longitude, speed, 根据预想的结果进行计算
			fmt.Println(input)

			// 在这里输出结果, 包含: 相同的时间戳下: 手机的数据, VCU的数据; 两者相同的时间戳的点的对比数据, 综合分析数据(方差什么的)

			_, _ = w.Write(v)
		}
	} else {
		http.Error(w, "request is wrong", 400)
	}

}

// searchBgnTime, searchEndTime 将是字符串类型的时间戳
func getEcData(searchBgnTime string, searchEndTime string, vcuNo string) []string{
	reqBody, err := json.Marshal(map[string]interface{}{
		"cmdType":       "0",
		"cmdValue":      "0x0A",
		"messageType":   "",
		"page":          1,
		"pageSize":      1000,
		"searchBgnTime": TimeStampToTimeString(searchBgnTime),
		"searchEndTime": TimeStampToTimeString(searchEndTime),
		"vcuNo":         vcuNo,
	})
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("POST", "http://gateway.global.lierda.com/columbia/dvc/cmdLog/selectByPage", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "eyJhbGciOiJIUzUxMiJ9.eyJvcmdUeXBlIjowLCJzdWIiOiJkb05vdExvZ2luIiwiYXVkaWVuY2UiOiJ3ZWIiLCJyb2xlIjoi6LaF57qn566h55CG5ZGYIiwiY3JlYXRlZCI6MTU3MzA5MDk1ODgzNywidHlwZSI6InN0YWZmQWRtaW4iLCJleHAiOjE1NzU2ODI5NTgsInVzZXJJZCI6NzZ9.CYIl1ZcFcGK4yIVR0NpFgzTy_f80IZfz7de2REtydMKa1Iij3zNVn8-jgvczTVyYJqW39rU4lfPeEW_IRn4wuA")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // 解析出后台返回的byte slice

	type RecordData struct {
		//Id          int         `json:"id"`
		//VcuNo       string      `json:"vcuNo"`
		//CmdType     string      `json:"cmdType"`
		//CmdValue    string      `json:"cmdValue"`
		//CmdTime     int         `json:"cmdTime"`
		CmdDetail   string      `json:"cmdDetail"`
		//MessageType interface{} `json:"messageType"`
	}
	type PageData struct {
		//Offset        int          `json:"offset"`
		//Limit         int          `json:"limit"`
		//Total         int          `json:"total"`
		//Size          int          `json:"size"`
		//Pages         int          `json:"pages"`
		//Current       int          `json:"current"`
		//SearchCount   bool         `json:"searchCount"`
		//OptimizeCount bool         `json:"optimizeCount"`
		//OrderByField  interface{}  `json:"orderByField"`
		Records       []RecordData `json:"records"`
	}
	type OutputData struct {
		CmdTypeMap map[string]string   `json:"cmdTypeMap"`
		Page       PageData `json:"page"`
	}
	type Output struct {
		Code string     `json:"code"`
		Msg  string     `json:"msg"`
		Data OutputData `json:"data"`
	}
	var o Output
	err = json.Unmarshal(body, &o)
	if err != nil {
		log.Println(err)
	}
	ecList := make([]string, len(o.Data.Page.Records))
	for i, v := range o.Data.Page.Records {
		ecList[i] = v.CmdDetail
	}
	//fmt.Println("从mongoDB里面拿到的数据: ", ecList)
	return ecList
}

func TimeStampToTimeString(timestamp string) string {
	s, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Fatalln(err)
	}
	tm := time.Unix(int64(s), 0)
	return tm.Format("2006-01-02 15:04:05")
}

//func processEcList(ecList []string) map[string][]string {
//
//}

func main() {
	//http.HandleFunc("/analysis", Analysis)
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	log.Fatalln("ListenAndServe", err)
	//}
	//getEcData("2019-10-31 10:35:00", "2019-10-31 10:45:00", "V1AB0100A19I240001")
	//fmt.Println(getEcData("1572489300", "1572489900", "V1AB0100A19I240001"))
	//fmt.Println(TimeStampToTimeString("1572489300"))
}
