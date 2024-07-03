package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 設置CORS中間件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3456","http://52.194.190.91",} // 允許的前端地址，根據實際情況修改
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// 設置路由處理函數
	r.GET("", func(c *gin.Context) {
		// 讀取 CSV 檔案並處理資料
		bankData, err := readBankData("data.csv")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "無法讀取銀行資料",
			})
			return
		}

		// 提取所有銀行的唯一bank_code和bank_name
		bankCodes := make(map[string]string)
		for _, data := range bankData {
			bankName := extractBankName(data["bank_name"])
			bankCodes[data["bank_code"]] = bankName
		}

		// 將map轉換成切片
		var banks []map[string]string
		for code, name := range bankCodes {
			banks = append(banks, map[string]string{
				"code": code,
				"name": name,
			})
		}

		// 返回 JSON 格式的資料
		c.JSON(http.StatusOK, banks)
	})

	// 設置處理 /:bank_code/branches 的路由
	r.GET("/:bank_code/branches", func(c *gin.Context) {
		bankCode := c.Param("bank_code")

		// 讀取 CSV 檔案並處理資料
		bankData, err := readBankData("data.csv")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "無法讀取銀行資料",
			})
			return
		}

		// 過濾符合 bank_code 的分行資料
		var branches []map[string]string
		for _, data := range bankData {
			if data["bank_code"] == bankCode {
				branch := map[string]string{
					"bank_code":   data["bank_code"],
					"branch_code": data["branch_code"],
					"address":     data["address"],
					"phone":       data["phone"],
				}
				branches = append(branches, branch)
			}
		}

		// 返回 JSON 格式的分行資料
		c.JSON(http.StatusOK, branches)
	})

	// 啟動 Gin 服務
	r.Run(":8080")
}

// readBankData 從 CSV 檔案中讀取銀行資料並返回一個 map 切片
func readBankData(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// 如果你要忽略首行（標頭），可以先讀取一次
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// 讀取剩餘的 CSV 記錄
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var bankData []map[string]string

	// 遍歷每一行記錄並存入 map 切片
	for _, record := range records {
		data := map[string]string{
			"bank_code":   record[0],
			"branch_code": record[1],
			"bank_name":     record[2],
			"address":       record[3],
			"phone":   record[4],
		}
		bankData = append(bankData, data)
	}

	return bankData, nil
}

// extractBankName 從銀行名稱中提取不包含「銀行」後面部分的名稱
func extractBankName(fullName string) string {
	re := regexp.MustCompile(`^(.*銀行)`)
	match := re.FindStringSubmatch(fullName)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return strings.TrimSpace(fullName)
}
