package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"
)

func main() {
	user, err := user.Current()
	if err != nil {
		println(err)
	}

	fmt.Println("merd Edge Cleaner")

	anadizin := user.HomeDir
	c := exec.Command("cmd", "/C", "Taskkill /IM msedge.exe /F")
	c.Wait()
	if err := c.Run(); err != nil {
		fmt.Println("Error: EDGE ZATEN KAPALI", err)
	}
	fmt.Println("msedge kapatıldı.")
	fmt.Println("anadizin", anadizin)
	fmt.Println("1 saniye bekleniyor")
	time.Sleep(1000 * time.Millisecond)
	os.RemoveAll(fmt.Sprintf("%s\\AppData\\Local\\Microsoft\\Edge\\User Data", anadizin))

	fmt.Println("msedge userdata silindi.")
	time.Sleep(100 * time.Millisecond)
	os.RemoveAll(fmt.Sprintf("%s\\AppData\\Local\\Temp", anadizin))
	fmt.Println("temp silindi.")

	ce := exec.Command("cmd", "/C ", "rmdir", "/s", "/q", "C:\\Windows\\Prefetch")
	if err := ce.Run(); err != nil {
		fmt.Println("Errore: ", err)
	}
	fmt.Println("prefetch silindi.")

	fmt.Println("EDGE TEMİZLENDİ BİLGİSAYARI YENİDEN BAŞLATMAK İÇİN BİR TUŞA BASIN!")
	fmt.Scanln()
	if err := exec.Command("cmd", "/C", "shutdown", "/r", "t", "00").Run(); err != nil {
		fmt.Println("Yeniden Başlatılamadı:", err)
	}

}
