package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func ReceiveHeader(c *fiber.Ctx) error {
	// Nhận dữ liệu từ request
	var data struct {
		Data string `json:"data"`
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Thực hiện công việc xử lý header

	// Bắt đầu một Goroutine để gửi header lên Telegram
	go SendHeadertoTelegram(data.Data)

	// Trả về response 200 OK nhanh nhất có thể
	return c.Status(http.StatusOK).SendString("Header received successfully")
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
		Text:   header,
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
