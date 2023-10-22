package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	ID             int
	DeviceID       string
	DeviceTag      string
	DeviceStatus   string
	DeviceLocation string
	IoTValue       string
	IoTValueType   string
	IssueID        string
}

// Cihaz etiketleri ve eşik değerlerini içeren bir harita
var thresholds = map[string]int{
	"water":    315,
	"gass":     1000,
	"electric": 1060,
}

func main() {
	db, err := sql.Open("mysql", "egitimbu_night:G5CDf9}!Ow?7@tcp(85.95.237.110:3306)/egitimbu_nightwatch?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	for {
		devices := generateDeviceData()
		for _, device := range devices {
			// Device.DeviceStatus değerini ayarlamak için son üç IoTValue'yi kullan
			lastThreeValues := getLastThreeValues(device.DeviceTag, devices)
			setDeviceStatus(&device, lastThreeValues, thresholds[device.DeviceTag])
			insertDeviceData(db, device)
		}
		time.Sleep(5 * time.Second)
	}
}

func generateDeviceData() []Device {
	devices := []Device{}

	// 1-4 arası cihazlar
	for i := 1; i <= 4; i++ {
		device := Device{
			ID:             i,
			DeviceID:       strconv.Itoa(i),
			DeviceTag:      "water",
			DeviceLocation: "water-map",
			IoTValueType:   "m³",
		}
		device.IoTValue = fmt.Sprintf("%d", rand.Intn(46)+280) // 280 ile 325 arasında rastgele bir değer
		// device.DeviceStatus = getStatus(device.IoTValue, 315)
		device.IssueID = "1" // Örnek olarak, tüm "water" cihazlarının Issue ID'si "1" olarak ayarlandı.
		devices = append(devices, device)
	}

	// 5-8 arası cihazlar
	for i := 5; i <= 8; i++ {
		device := Device{
			ID:             i,
			DeviceID:       strconv.Itoa(i),
			DeviceTag:      "gass",
			DeviceLocation: "gass-map",
			IoTValueType:   "m³",
		}
		device.IoTValue = fmt.Sprintf("%d", rand.Intn(56)+960) // 960 ile 1015 arasında rastgele bir değer
		// device.DeviceStatus = getStatus(device.IoTValue, 1000)
		device.IssueID = "2" // Örnek olarak, tüm "gass" cihazlarının Issue ID'si "2" olarak ayarlandı.
		devices = append(devices, device)
	}

	// 9-12 arası cihazlar
	for i := 9; i <= 12; i++ {
		device := Device{
			ID:             i,
			DeviceID:       strconv.Itoa(i),
			DeviceTag:      "electric",
			DeviceLocation: "electric-map",
			IoTValueType:   "kWh",
		}
		device.IoTValue = fmt.Sprintf("%d", rand.Intn(181)+990) // 990 ile 1170 arasında rastgele bir değer
		// device.DeviceStatus = getStatus(device.IoTValue, 1060)
		device.IssueID = "3" // Örnek olarak, tüm "electric" cihazlarının Issue ID'si "3" olarak ayarlandı.
		devices = append(devices, device)
	}

	return devices
}

func getLastThreeValues(deviceTag string, devices []Device) []string {
	values := []string{}
	for _, device := range devices {
		if device.DeviceTag == deviceTag {
			values = append(values, device.IoTValue)
		}
	}
	// Son üç değeri al
	if len(values) > 3 {
		values = values[len(values)-3:]
	}
	return values
}

func setDeviceStatus(device *Device, values []string, threshold int) {
	if len(values) < 3 {
		return // Son üç değer yoksa işlem yapma
	}
	average := calculateAverage(values)
	if int(average) > threshold {
		device.DeviceStatus = "Error"
		setIssue(generateRandomRequestIssue(device.DeviceID))
	} else {
		device.DeviceStatus = "No Error"
	}
}

func calculateAverage(values []string) float64 {
	sum := 0
	for _, value := range values {
		num := parseInt(value)
		sum += num
	}
	return float64(sum) / float64(len(values))
}

func parseInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func insertDeviceData(db *sql.DB, device Device) {
	_, err := db.Exec("INSERT INTO iot_entity (device_id, device_tag, device_statu, device_location, iot_value, iot_value_type, issueId) VALUES (?, ?, ?, ?, ?, ?, ?)",
		device.DeviceID, device.DeviceTag, device.DeviceStatus, device.DeviceLocation, device.IoTValue, device.IoTValueType, device.IssueID)
	if err != nil {
		fmt.Println("Veri eklenirken hata oluştu:", err)
	}
}

type requestIssue struct {
	Deviceid  string `json:"device_id"`
	IssueType string `json:"issue_type"`
	UserId    string `json:"user_id"`
}

func generateRandomRequestIssue(deviceID string) requestIssue {
	// Rastgele olarak 1 veya 2 atanır
	userID := "1"
	if rand.Intn(2) == 1 {
		userID = "2"
	}

	// Rastgele olarak "Acil", "Kesme", "Bakım" veya "Onarım" atanır
	issueTypes := []string{"Acil", "Kesme", "Bakım", "Onarım"}
	issueType := issueTypes[rand.Intn(len(issueTypes))]

	return requestIssue{
		Deviceid:  deviceID,
		UserId:    userID,
		IssueType: issueType,
	}
}

func setIssue(req requestIssue) {
	address := "https://nightwatch-api-production.up.railway.app/api/issue/createIssue"

	jsonVeri, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", address, bytes.NewBuffer(jsonVeri))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Authorization", "Bearer <token>")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}

//Acil Kesme Bakım Onarım
