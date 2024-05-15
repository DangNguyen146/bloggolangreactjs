package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ReceiveLink(c *fiber.Ctx) error {
	// Nhận dữ liệu từ Parameters
	data, err := strconv.Atoi(c.Query("data", ""))
	if err != nil {
		return err
	}

	// Thực hiện công việc xử lý dữ liệu

	// Gửi dữ liệu lên Telegram
	go SendHeadertoTelegram(strconv.Itoa(data))

	// Trả về response 200 OK
	return c.SendStatus(fiber.StatusOK)
}

// Struct đại diện cho thông tin cần gửi lên nhóm Telegram
type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendHeadertoTelegram(header string) {
	// Thông tin API token và chat ID của nhóm Telegram
	apiToken := os.Getenv("YOUR_TELEGRAM_API_TOKEN")
	chatID := os.Getenv("YOUR_GROUP_CHAT_ID")

	// Tạo cấu trúc tin nhắn
	message := TelegramMessage{
		ChatID: chatID,
		Text:   "Có người vừa truy cặp link:" + header,
	}

	// Chuyển đổi cấu trúc tin nhắn thành JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Tạo request POST đến API của Telegram
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken)
	req, err := http.NewRequest("POST", telegramURL, bytes.NewBuffer(messageBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Thực hiện request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Kiểm tra phản hồi từ Telegram
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status code:", resp.StatusCode)
		return
	}

	fmt.Println("Header sent to Telegram group successfully.")
}
