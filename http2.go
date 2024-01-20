//Author:github.com/Azumi67
//This script is for educational use and for my own learning, but I'd be happy if you find it useful too.
//This script simplifies the configuration of Go http2 reverse tunnel.
//You can send me feedback so I can use it to learn more.
//This script comes without any warranty
//Thank you.
package main
import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509/pkix"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
    "github.com/fatih/color"
	"log"
	"github.com/AlecAivazis/survey/v2"
	"runtime"
)

func displayProgress(total, current int) {
	width := 40
	percentage := current * 100 / total
	completed := width * current / total
	remaining := width - completed

	fmt.Printf("\r[%s>%s] %d%%", strings.Repeat("=", completed), strings.Repeat(" ", remaining), percentage)
}

func displayError(message string) {
	fmt.Printf("\u2718 Error: %s\n", message)
}

func displayNotification(message string) {
	fmt.Printf("\033[93m%s\033[0m\n", message)
}

func displayCheckmark(message string) {
	fmt.Printf("\033[92m\u2714 \033[0m%s\n", message)
}

func displayLoading() {
    frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
    delay := 100 * time.Millisecond
    duration := 5 * time.Second

    endTime := time.Now().Add(duration)

    for time.Now().Before(endTime) {
        for _, frame := range frames {
            fmt.Printf("\r[%s] Loading... ", frame)
            time.Sleep(delay)
        }
    }
    fmt.Println()
}
func displayLogo2() error {
	cmd := exec.Command("bash", "-c", "/etc/./logo.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
func displayLogo() {
	logo := `
   ______    _______    __      _______          __       _____  ___  
  /    " \  |   __ "\  |" \    /"      \        /""\      (\"   \|" \ 
 // ____  \ (. |__) :) ||  |  |:        |      /    \     |.\\   \   |
/  /    ) :)|:  ____/  |:  |  |_____/   )     /' /\  \    |: \.   \\ |
(: (____/ // (|  /     |.  |   //       /    //  __'  \   |.  \    \ |
\        // |__/ \     /\  |\  |:  __   \   /   /  \\  \  |    \    \|
 \"_____ / (_______)  (__\_|_) |__|  \___) (___/    \___) \___|\____\)
`
	
    cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
    blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	green := color.New(color.FgHiGreen, color.Bold).SprintFunc()      
    yellow := color.New(color.FgHiYellow, color.Bold).SprintFunc()   
    red := color.New(color.FgHiRed, color.Bold).SprintFunc()        


	

	    logo = cyan("  ______   ") + blue(" _______  ") + green("  __    ") + yellow("   ______   ") + red("     __      ") + cyan("  _____  ___  \n") +
		cyan(" /     \" \\  ") + blue("|   __ \" ") + green(" |\" \\  ") + yellow("   /\"     \\   ") + red("   /\"\"\\     ") + cyan(" (\\\"   \\|\"  \\ \n") +
		cyan("//  ____  \\ ")  + blue("(. |__) :)") + green("||  |  ") + yellow(" |:       |  ") + red("  /    \\   ") + cyan("  |.\\\\   \\   |\n") +
		cyan("/  /    ) :)") + blue("|:  ____/ ") + green("|:  |  ") + yellow(" |_____/  )  ") + red(" /' /\\  \\   ") + cyan(" |: \\.   \\\\ |\n") +
		cyan("(: (____/ / ") + blue("(|  /     ") + green("|.  | ") + yellow("  //      /  ") + red("//   __'  \\  ") + cyan(" |.  \\    \\ |\n") +
		cyan("\\        / ") + blue("/|__/ \\   ") + green(" /\\  |\\ ") + yellow(" |:  __  \\ ") + red(" /   /  \\\\   ") + cyan ("  |    \\    \\|\n") +
		cyan(" \"_____ / ") + blue("(_______)") + green("  (__\\_|_)") + yellow(" (__) \\___)") + red("(___/    \\___)") + cyan(" \\___|\\____\\)\n")


	fmt.Println(logo)
}
func main() {
	if os.Geteuid() != 0 {
		fmt.Println("\033[91mThis script must be run as root. Please use sudo -i.\033[0m")
		os.Exit(1)
	}

	mainMenu()
}
func readInput() {
	fmt.Print("Press Enter to continue..")
	fmt.Scanln()
	mainMenu()
}
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func mainMenu() {
	for {
		err := displayLogo2()
		if err != nil {
			log.Fatalf("failed to display logo: %v", err)
		}
		displayLogo()
		border := "\033[93m+" + strings.Repeat("=", 70) + "+\033[0m"
		content := "\033[93m║            ▌║█║▌│║▌│║▌║▌█║ \033[92mMain Menu\033[93m  ▌│║▌║▌│║║▌█║▌                  ║"
		footer := " \033[92m            Join Opiran Telegram \033[34m@https://t.me/OPIranClub\033[0m "

		borderLength := len(border) - 2
		centeredContent := fmt.Sprintf("%[1]*s", -borderLength, content)

		fmt.Println(border)
		fmt.Println(centeredContent)
		fmt.Println(border)

		fmt.Println(border)
		fmt.Println(footer)
		fmt.Println(border)
		prompt := &survey.Select{
			Message: "Enter your choice Please:",
			Options: []string{"0. \033[91mSTATUS Menu\033[0m", "1. \033[92mStop | \033[96mRestart \033[93mService\033[0m", "2. \033[93mIPV4 \033[92mTCP \033[0m", "3. \033[96mEdit \033[92mReset Timer \033[0m", "4. \033[91mUninstall\033[0m",  "q. Exit"},
		
		}
		fmt.Println("\033[93m╰─────────────────────────────────────────────────────────────────────╯\033[0m")

		var choice string
		err = survey.AskOne(prompt, &choice)
		if err != nil {
			log.Fatalf("\033[91muser input is wrong:\033[0m %v", err)
		}
		switch choice {
		case "0. \033[91mSTATUS Menu\033[0m":
			status()
		case "1. \033[92mStop | \033[961mRestart \033[93mService\033[0m":
			startMain()
		case "2. \033[93mIPV4 \033[92mTCP \033[0m":
			tcp4Menu()
		case "3. \033[96mEdit \033[92mReset Timer \033[0m":
			cronMenu()
		case "4. \033[91mUninstall\033[0m":
			UniMenu()
		case "q. Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}

		
		readInput()
	}
}
func tcp4Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mTCP \033[96mIPV4 \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[92m\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		iranno4()
	case "2. \033[93mKHAREJ\033[92m\033[0m":
		kharejno4()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharejno4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	if _, err := os.Stat("/root/tunnel"); os.IsNotExist(err) {
		godl()
	}

	reader := bufio.NewReader(os.Stdin)

	err := os.Chdir("/root/tunnel")
	if err != nil {
		fmt.Println("\033[91mCouldn't change to server dir\033[0m")
		return
	}

	err = certKey()
	if err != nil {
		fmt.Println("\033[91mGenerating SSL cert failed gloriously\033[0m", err)
		return
	}

	var iranIP4, tunnelPort string
	var numConfigs int
	var kharejConfigPorts []string
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	for {
		fmt.Print("\033[93mEnter \033[92mIRAN IPV4\033[93m:\033[0m ")
		iranIP4, _ = reader.ReadString('\n')
		iranIP4 = strings.TrimSpace(iranIP4)
		if iranIP4 != "" {
			break
		}
		fmt.Println("\033[91minvalid input. Plz try again\033[0m")
	}

	for {
		fmt.Print("\033[93mEnter the \033[92mtunnel port\033[93m:\033[0m ")
		tunnelPort, _ = reader.ReadString('\n')
		tunnelPort = strings.TrimSpace(tunnelPort)
		if tunnelPort != "" {
			break
		}
		fmt.Println("\033[91minvalid input. Plz try again\033[0m")
	}

	for {
		fmt.Print("\033[93mEnter the \033[92mnumber of configs\033[93m: ")
		numConfigsStr, _ := reader.ReadString('\n')
		numConfigsStr = strings.TrimSpace(numConfigsStr)
		numConfigs, err = strconv.Atoi(numConfigsStr)
		if err != nil || numConfigs < 0 {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
		} else {
			break
		}
	}

	for i := 0; i < numConfigs; i++ {
		for {
			fmt.Printf("\033[93mEnter \033[92mConfig\033[96m port\033[93m %d:\033[0m ", i+1)
			kharejConfigPort, _ := reader.ReadString('\n')
			kharejConfigPort = strings.TrimSpace(kharejConfigPort)
			if kharejConfigPort != "" {
				kharejConfigPorts = append(kharejConfigPorts, kharejConfigPort)
				break
			}
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
		}
	}

	config := fmt.Sprintf(`server_addr: "%s:%s"
tls_crt: /root/tunnel/client.crt
tls_key: /root/tunnel/client.key
root_ca: ""

backoff:
  interval: 500ms
  multiplier: 1.5
  max_interval: 1m0s
  max_time: 15m0s

tunnels:`, iranIP4, tunnelPort)

	for i, port := range kharejConfigPorts {
		tunnel := fmt.Sprintf(`
  tcp_tunnel%d:
    proto: tcp
    addr: "localhost:%s"
    remote_addr: 0.0.0.0:%s`, i+1, port, port)
		config += tunnel
	}

	conf, err := os.Create("/root/tunnel/config.yaml")
	if err != nil {
		fmt.Println("\033[91mCreating config has failed!:\033[0m", err)
		return
	}
	defer conf.Close()

	_, err = conf.WriteString(config)
	if err != nil {
		fmt.Println("\033[91mCouldn't copy contents into da Config:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Go-Http2-Tunnel
After=network.target
After=network-online.target

[Service]
ExecStart=/root/tunnel/./tunnel -config /root/tunnel/config.yaml start-all
TimeoutSec=30
Restart=on-failure
RestartSec=5
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target`

	serviceFile, err := os.Create("/etc/systemd/system/gotunnel-kharej.service")
	if err != nil {
		fmt.Println("\033[91mCreating the service has failed!:\033[0m", err)
		return
	}
	defer serviceFile.Close()

	_, err = serviceFile.WriteString(service)
	if err != nil {
		fmt.Println("\033[91mCouldn't copy the contents into service:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mFailed to reload:\033[0m", err)
		return
	}

	cmd = exec.Command("chmod", "u+x", "/etc/systemd/system/gotunnel-kharej.service")
	err = cmd.Run()
	if err!= nil {
		fmt.Println("\033[91mCouldn't set permission for the service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "gotunnel-kharej.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mEnabling the service has failed:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "gotunnel-kharej.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mStarting the service has failed:\033[0m", err)
		return
	}
    resKharej()
	displayCheckmark("\033[92mService configured successfully\033[0m")
}

func certKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("\033[91mGenerating private key has failed: %s\033[0m", err)
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("\033[91mCreating cert failed: %s\033[0m", err)
	}

	privatePath, err := os.Create("/root/tunnel/client.key")
	if err != nil {
		return fmt.Errorf("\033[91mCreating private key failed: %s\033[0m", err)
	}
	defer privatePath.Close()

	azumiKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privatePath, azumiKey)
	if err != nil {
		return fmt.Errorf("\033[91mWritng privatekey content failed: %s\033[0m", err)
	}

	certAzumi, err := os.Create("/root/tunnel/client.crt")
	if err != nil {
		return fmt.Errorf("\033[91mCreating client cert failed: %s\033[0m", err)
	}
	defer certAzumi.Close()

	err = pem.Encode(certAzumi, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return fmt.Errorf("\033[91mWriting cert content failed: %s\033[0m", err)
	}

	return nil
}
func iranno4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mTCP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	if _, err := os.Stat("/root/tunnel"); os.IsNotExist(err) {
		godl()
	}

	if _, err := os.Stat("/etc/http.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/http.sh")
	}
	var tunnelPort, httpsPort, httpPort int
	var err error
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	for {
		fmt.Print("\033[93mEnter the \033[92mTunnel\033[93m port:\033[0m ")
		_, err = fmt.Scan(&tunnelPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}

		fmt.Print("\033[93mEnter \033[92mHTTPS\033[93m port:\033[0m ")
		_, err = fmt.Scan(&httpsPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}

		fmt.Print("\033[93mEnter \033[92mHTTP\033[93m port:\033[0m ")
		_, err = fmt.Scan(&httpPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}

		break
	}

	if err != nil {
		fmt.Println("\033[91mCouldn't read da input\033[0m")
		return
	}

	err = os.Chdir("/root/tunnel")
	if err != nil {
		fmt.Println("\033[91mCouldn't change to server dir\033[0m")
		return
	}

	err = certKey2()
	if err != nil {
		fmt.Println("\033[91mGenerating SSL cert failed gloriously\033[0m", err)
		return
	}

	service := fmt.Sprintf(`[Unit]
Description=Go-Http2-Tunnel
After=network.target
After=network-online.target

[Service]
ExecStart=/root/tunnel/./tunneld -tlsCrt /root/tunnel/server.crt -tlsKey /root/tunnel/server.key -tunnelAddr 0.0.0.0:%d -httpsAddr :%d -httpAddr :%d
TimeoutSec=30
Restart=on-failure
RestartSec=5
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target`, tunnelPort, httpsPort, httpPort)

	serv, err := os.Create("/etc/systemd/system/gotunnel-iran.service")
	if err != nil {
		fmt.Println("\033[91mCouldn't create da service\033[0m")
		return
	}
	defer serv.Close()

	_, err = serv.WriteString(service)
	if err != nil {
		fmt.Println("\033[91mCouldn't put config inside the service\033[0m")
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mFailed to reload\033[0m")
		return
	}

	cmd = exec.Command("chmod", "u+x", "/etc/systemd/system/gotunnel-iran.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't set permission for the service\033[0m")
		return
	}

	cmd = exec.Command("systemctl", "enable", "gotunnel-iran")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't enable service\033[0m")
		return
	}

	cmd = exec.Command("systemctl", "restart", "gotunnel-iran")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't restart service\033[0m")
		return
	}
    resIran()
	displayCheckmark("\033[92mService configured successfully\033[0m")
}
func certKey2() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("\033[91mGenerating private key has failed: %s\033[0m", err)
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("\033[91mCreating cert failed: %s\033[0m", err)
	}

	privatePath, err := os.Create("/root/tunnel/server.key")
	if err != nil {
		return fmt.Errorf("\033[91mCreating private key failed: %s\033[0m", err)
	}
	defer privatePath.Close()

	azumiKey := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privatePath, azumiKey)
	if err != nil {
		return fmt.Errorf("\033[91mWritng privatekey content failed: %s\033[0m", err)
	}

	certAzumi, err := os.Create("/root/tunnel/server.crt")
	if err != nil {
		return fmt.Errorf("\033[91mCreating server cert failed: %s\033[0m", err)
	}
	defer certAzumi.Close()

	err = pem.Encode(certAzumi, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		return fmt.Errorf("\033[91mWriting cert content failed: %s\033[0m", err)
	}

	return nil
}
func cronMenu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reset \033[92mTimer \033[93mMenu\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIRAN\033[0m", "2. \033[93mKHAREJ\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIRAN\033[0m":
		resIran()
	case "2. \033[93mKHAREJ\033[0m":
		resKharej()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func installICMP() {
	displayNotification("\033[93mInstalling \033[92mIcmptunnel\033[93m ...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayLoading()

	ipv4ForwardStatus, _ := exec.Command("sysctl", "-n", "net.ipv4.ip_forward").Output()
	if string(ipv4ForwardStatus) != "1\n" {
		exec.Command("sysctl", "net.ipv4.ip_forward=1").Run()
	}

	if _, err := os.Stat("/root/icmptunnel"); err == nil {
		os.RemoveAll("/root/icmptunnel")
	}

	cloneCommand := "git clone https://github.com/jamesbarlow/icmptunnel.git icmptunnel"
	cloneCmd := exec.Command("bash", "-c", cloneCommand)
	cloneOutput, cloneErr := cloneCmd.StdoutPipe()
	if cloneErr != nil {
		fmt.Println("git clone failed!")
		return
	}

	cloneCmd.Start()

	go func() {
		scanner := bufio.NewScanner(cloneOutput)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	cloneCmd.Wait()

	if _, err := os.Stat("/root/icmptunnel"); err == nil {
		os.Chdir("/root/icmptunnel")

		exec.Command("sudo", "apt", "install", "-y", "net-tools").Run()
		exec.Command("sudo", "apt", "install", "-y", "make").Run()
		exec.Command("sudo", "apt-get", "install", "-y", "libssl-dev").Run()
		exec.Command("sudo", "apt", "install", "-y", "g++").Run()

		exec.Command("make").Run()

		os.Chdir("..")
	} else {
		displayError("\033[91micmptunnel folder not found !\033[0m")
	}
}
func disableICMPEcho() {
	cmd := exec.Command("bash", "-c", "echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all")
	err := cmd.Run()

	if err != nil {
		displayError(fmt.Sprintf("\033[91mError occurred disabling echo:\033[0m %s", err.Error()))
	} else {
		displayCheckmark("\033[92mecho disabled..\033[0m")
	}
}
func forward() {
	ipForward("net.ipv4.ip_forward", "0", "1")
	ipForward("net.ipv6.conf.all.forwarding", "0", "1")
}

func ipForward(sysctlKey, expectedValue, newValue string) {
	cmd := exec.Command("sysctl", sysctlKey)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\033[91mSysctl command for %s failed: %s\n\033[0m", sysctlKey, err)
		return
	}

	if !strings.Contains(string(output), sysctlKey+" = "+expectedValue) {
		fmt.Printf("\033[93m%s is already enabled!\033[0m\n", sysctlKey)
		return
	}

	enableCmd := exec.Command("sudo", "sysctl", "-w", sysctlKey+"="+newValue)
	enableOutput, err := enableCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\033[91mCouldn't enable %s: %s\n\033[0m", sysctlKey, err)
		return
	}

	fmt.Printf("\033[92m%s enabled successfully!\nOutput: %s\033[0m\n", sysctlKey, enableOutput)
}
func startICKharej() {
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mConfiguring ICMPtunnel \033[92mKharej\033[93m ...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	disableICMPEcho()

	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	}

	if _, err := os.Stat("/etc/icmp.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp.sh")
	}

	file, err := os.Create("/etc/icmp.sh")
	if err != nil {
		fmt.Println("\033[91mError creating icmp.sh:", err, "\033[0m")
		return
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("/root/icmptunnel/icmptunnel -s -d\n")
	file.WriteString("/sbin/ifconfig tun0 70.0.0.1 netmask 255.255.255.0\n")

	os.Chmod("/etc/icmp.sh", 0700)

	cmd := exec.Command("/bin/bash", "/etc/icmp.sh")
	cmd.Run()

	cronJobCommand := "@reboot /bin/bash /etc/icmp.sh\n"
	cronFile, err := os.Create("/etc/cron.d/icmp-kharej")
	if err != nil {
		fmt.Println("\033[91mError creating cron:", err, "\033[0m")
		return
	}
	defer cronFile.Close()

	cronFile.WriteString(cronJobCommand)

	cronCmd := exec.Command("crontab", "-u", "root", "/etc/cron.d/icmp-kharej")
	cronCmd.Run()

	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}
func runICMP() {
	cmd := exec.Command("/bin/bash", "/etc/icmp-iran.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't run da script:", err, "\033[0m")
		return
	}
	fmt.Println("\033[92mScript ran successfully!\033[0m")
}
func startICIran() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mConfiguring ICMPtunnel \033[92mIRAN \033[93m...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	disableICMPEcho()

	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	}
	if _, err := os.Stat("/etc/icmp.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp.sh")
	}
	err := os.Chdir("/root/icmptunnel")
	if err != nil {
		fmt.Println("\033[91mError using CD:", err, "\033[0m")
		return
	}

	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mKharej\033[93m IPv4 address:\033[0m ")
	serverIPv4, _ := reader.ReadString('\n')
	serverIPv4 = strings.TrimSuffix(serverIPv4, "\n") 

	if _, err := os.Stat("/etc/icmp-iran.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp-iran.sh")
	}

	file, err := os.Create("/etc/icmp-iran.sh")
	if err != nil {
		fmt.Println("\033[91mError creating icmp-iran.sh:", err, "\033[0m")
		return
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString(fmt.Sprintf("/root/icmptunnel/icmptunnel %s -d\n", serverIPv4))
	file.WriteString("/sbin/ifconfig tun0 70.0.0.2 netmask 255.255.255.0\n")

	os.Chmod("/etc/icmp-iran.sh", 0700)

	cmd := exec.Command("/bin/bash", "/etc/icmp-iran.sh")
	cmd.Run()

	cronJobCommand := "@reboot /bin/bash /etc/icmp-iran.sh\n"
	cronFile, err := os.Create("/etc/cron.d/icmp-iran")
	if err != nil {
		fmt.Println("\033[91mError creating cron:", err, "\033[0m")
		return
	}
	defer cronFile.Close()

	cronFile.WriteString(cronJobCommand)

	cronCmd := exec.Command("crontab", "-u", "root", "/etc/cron.d/icmp-iran")
	cronCmd.Run()
    runICMP()
	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}

func kharejicmp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mICMP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring KHAREJ")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	if _, err := os.Stat("/root/tunnel"); os.IsNotExist(err) {
		godl()
	}

    startICKharej()

	reader := bufio.NewReader(os.Stdin)

	err := os.Chdir("/root/tunnel")
	if err != nil {
		fmt.Println("\033[91mCouldn't change to server dir\033[0m")
		return
	}

	err = certKey()
	if err != nil {
		fmt.Println("\033[91mGenerating SSL cert failed gloriously\033[0m", err)
		return
	}

	var tunnelPort string
	var numConfigs int
	var kharejConfigPorts []string
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	for {
		fmt.Print("\033[93mEnter the \033[92mTunnel port\033[93m: ")
		tunnelPort, _ = reader.ReadString('\n')
		tunnelPort = strings.TrimSpace(tunnelPort)
		if tunnelPort != "" {
			break
		}
		fmt.Println("\033[91minvalid input. Plz try again\033[0m")
	}

	for {
		fmt.Print("\033[93mEnter the \033[92mnumber of Configs\033[93m:\033[0m ")
		numConfigsStr, _ := reader.ReadString('\n')
		numConfigsStr = strings.TrimSpace(numConfigsStr)
		numConfigs, err = strconv.Atoi(numConfigsStr)
		if err != nil || numConfigs < 0 {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
		} else {
			break
		}
	}

	for i := 0; i < numConfigs; i++ {
		for {
			fmt.Printf("\033[93mEnter \033[92mConfig\033[96m port\033[93m %d:\033[0m ", i+1)
			kharejConfigPort, _ := reader.ReadString('\n')
			kharejConfigPort = strings.TrimSpace(kharejConfigPort)
			if kharejConfigPort != "" {
				kharejConfigPorts = append(kharejConfigPorts, kharejConfigPort)
				break
			}
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
		}
	}

	config := fmt.Sprintf(`server_addr: "70.0.0.2:%s"
tls_crt: /root/tunnel/client.crt
tls_key: /root/tunnel/client.key
root_ca: ""

backoff:
  interval: 500ms
  multiplier: 1.5
  max_interval: 1m0s
  max_time: 15m0s

tunnels:`, tunnelPort)

	for i, port := range kharejConfigPorts {
		tunnel := fmt.Sprintf(`
  tcp_tunnel%d:
    proto: tcp
    addr: "localhost:%s"
    remote_addr: 0.0.0.0:%s`, i+1, port, port)
		config += tunnel
	}

	conf, err := os.Create("/root/tunnel/config.yaml")
	if err != nil {
		fmt.Println("\033[91mCreating config has failed!:\033[0m", err)
		return
	}
	defer conf.Close()

	_, err = conf.WriteString(config)
	if err != nil {
		fmt.Println("\033[91mCouldn't copy contents into da Config:\033[0m", err)
		return
	}

	service := `[Unit]
Description=Go-Http2-Tunnel
After=network.target
After=network-online.target

[Service]
ExecStart=/root/tunnel/./tunnel -config /root/tunnel/config.yaml start-all
TimeoutSec=30
Restart=on-failure
RestartSec=5
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target`

	serv, err := os.Create("/etc/systemd/system/gotunnel-kharej.service")
	if err != nil {
		fmt.Println("\033[91mCreating the service has failed!:\033[0m", err)
		return
	}
	defer serv.Close()

	_, err = serv.WriteString(service)
	if err != nil {
		fmt.Println("\033[91mCouldn't copy the contents into service:\033[0m", err)
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mFailed to reload:\033[0m", err)
		return
	}

	cmd = exec.Command("chmod", "u+x", "/etc/systemd/system/gotunnel-kharej.service")
	err = cmd.Run()
	if err!= nil {
		fmt.Println("\033[91mCouldn't set permission for the service:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "enable", "gotunnel-kharej.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mEnabling the service has failed:\033[0m", err)
		return
	}

	cmd = exec.Command("systemctl", "restart", "gotunnel-kharej.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mStarting the service has failed:\033[0m", err)
		return
	}
    resKharej()
	displayCheckmark("\033[92mService configured successfully\033[0m")
}

func iranicmp4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Reverse \033[92mIPV4 \033[96mICMP\033[0m ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	if _, err := os.Stat("/root/tunnel"); os.IsNotExist(err) {
		godl()
	}
	startICIran()

	var tunnelPort, httpsPort, httpPort int
	var err error
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	for {
		fmt.Print("\033[93mEnter the \033[92mTunnel\033[93m port:\033[0m ")
		_, err = fmt.Scan(&tunnelPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}
	
		fmt.Print("\033[93mEnter \033[92mHTTPS\033[93m port:\033[0m ")
		_, err = fmt.Scan(&httpsPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}

		fmt.Print("\033[93mEnter \033[92mHTTP\033[93m port:\033[0m ")
		_, err = fmt.Scan(&httpPort)
		if err != nil {
			fmt.Println("\033[91minvalid input. Plz try again\033[0m")
			break
		}

		break
	}

	if err != nil {
		fmt.Println("\033[91mCouldn't read da input\033[0m")
		return
	}

	err = os.Chdir("/root/tunnel")
	if err != nil {
		fmt.Println("\033[91mCouldn't change to server dir\033[0m")
		return
	}

	err = certKey2()
	if err != nil {
		fmt.Println("\033[91mGenerating SSL cert failed gloriously\033[0m", err)
		return
	}

	service := fmt.Sprintf(`[Unit]
Description=Go-Http2-Tunnel
After=network.target
After=network-online.target

[Service]
ExecStart=/root/tunnel/./tunneld -tlsCrt /root/tunnel/server.crt -tlsKey /root/tunnel/server.key -tunnelAddr 0.0.0.0:%d -httpsAddr :%d -httpAddr :%d
TimeoutSec=30
Restart=on-failure
RestartSec=5
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target`, tunnelPort, httpsPort, httpPort)

	serv, err := os.Create("/etc/systemd/system/gotunnel-iran.service")
	if err != nil {
		fmt.Println("\033[91mCouldn't create da service\033[0m")
		return
	}
	defer serv.Close()

	_, err = serv.WriteString(service)
	if err != nil {
		fmt.Println("\033[91mCouldn't put config inside the service\033[0m")
		return
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mFailed to reload\033[0m")
		return
	}

	cmd = exec.Command("chmod", "u+x", "/etc/systemd/system/gotunnel-iran.service")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't set permission for the service\033[0m")
		return
	}

	cmd = exec.Command("systemctl", "enable", "gotunnel-iran")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't enable service\033[0m")
		return
	}

	cmd = exec.Command("systemctl", "restart", "gotunnel-iran")
	err = cmd.Run()
	if err != nil {
		fmt.Println("\033[91mCouldn't restart service\033[0m")
		return
	}
    resIran()
	displayCheckmark("\033[92mService configured successfully\033[0m")
}
func rmv() error {
	file := "/etc/http.sh"
	if _, err := os.Stat(file); err == nil {
		err := os.Remove(file)
		if err != nil {
			return fmt.Errorf("\033[91mbash file doesn't exists:\033[0m %v", err)
		}
		fmt.Println("\033[91mbash file removed successfully!\033[0m")
	}
	return nil
}
func deleteCron() {
	entriesToDelete := []string{
		"*/1 * * * * /etc/http.sh",
		"*/2 * * * * /etc/http.sh",
		"*/3 * * * * /etc/http.sh",
		"*/4 * * * * /etc/http.sh",
		"*/5 * * * * /etc/http.sh",
		"*/6 * * * * /etc/http.sh",
		"*/7 * * * * /etc/http.sh",
		"*/8 * * * * /etc/http.sh",
		"*/9 * * * * /etc/http.sh",
		"*/10 * * * * /etc/http.sh",
		"*/11 * * * * /etc/http.sh",
		"*/12 * * * * /etc/http.sh",
		"*/13 * * * * /etc/http.sh",
		"*/14 * * * * /etc/http.sh",
		"*/15 * * * * /etc/http.sh",
		"*/16 * * * * /etc/http.sh",
		"*/17 * * * * /etc/http.sh",
		"*/18 * * * * /etc/http.sh",
		"*/19 * * * * /etc/http.sh",
		"*/20 * * * * /etc/http.sh",
		"*/21 * * * * /etc/http.sh",
		"*/22 * * * * /etc/http.sh",
		"*/23 * * * * /etc/http.sh",
		"*/24 * * * * /etc/http.sh",
		"*/25 * * * * /etc/http.sh",
		"*/26 * * * * /etc/http.sh",
		"*/27 * * * * /etc/http.sh",
	    "*/28 * * * * /etc/http.sh",
		"*/29 * * * * /etc/http.sh",
		"*/30 * * * * /etc/http.sh",
		"*/31 * * * * /etc/http.sh",
		"*/32 * * * * /etc/http.sh",
		"*/33 * * * * /etc/http.sh",
		"*/34 * * * * /etc/http.sh",
		"*/35 * * * * /etc/http.sh",
		"*/36 * * * * /etc/http.sh",
		"*/37 * * * * /etc/http.sh",
		"*/38 * * * * /etc/http.sh",
		"*/39 * * * * /etc/http.sh",
		"*/40 * * * * /etc/http.sh",
		"*/41 * * * * /etc/http.sh",
		"*/42 * * * * /etc/http.sh",
		"*/43 * * * * /etc/http.sh",
		"*/44 * * * * /etc/http.sh",
		"*/45 * * * * /etc/http.sh",
		"*/46 * * * * /etc/http.sh",
		"*/47 * * * * /etc/http.sh",
		"*/48 * * * * /etc/http.sh",
		"*/49 * * * * /etc/http.sh",
		"*/50 * * * * /etc/http.sh",
		"*/51 * * * * /etc/http.sh",
		"*/52 * * * * /etc/http.sh",
		"*/53 * * * * /etc/http.sh",
		"*/54 * * * * /etc/http.sh",
		"*/55 * * * * /etc/http.sh",
		"*/56 * * * * /etc/http.sh",
		"*/57 * * * * /etc/http.sh",
		"*/58 * * * * /etc/http.sh",
		"*/59 * * * * /etc/http.sh",
		
	}

	existingCrontab, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		fmt.Println("\033[91mNo existing cron found!\033[0m")
		return
	}

	newCrontab := string(existingCrontab)
	for _, entry := range entriesToDelete {
		if strings.Contains(newCrontab, entry) {
			newCrontab = strings.Replace(newCrontab, entry, "", -1)
		}
	}

	if newCrontab != string(existingCrontab) {
		cmd := exec.Command("crontab")
		cmd.Stdin = strings.NewReader(newCrontab)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		displayNotification("\033[92mDeleting Previous Crons..\033[0m")
	} else {
		fmt.Println("\033[91mCron doesn't exist, moving on..!\033[0m")
	}
}

func startMain() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Service \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mRestart\033[0m", "2. \033[93mStop \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mRestart\033[0m":
		start()
	case "2. \033[93mStop \033[0m":
		stop()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func start() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Restart \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP\033[0m":
		restarttcp()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func restarttcp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mRestarting Reverse Tunnel \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	cmd = exec.Command("systemctl", "restart", "gotunnel-kharej")
	cmd.Run()
	time.Sleep(1 * time.Second)

	cmd = exec.Command("systemctl", "restart", "gotunnel-iran")
	cmd.Run()
	time.Sleep(1 * time.Second)

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mRestart completed!\033[0m")
}
func stop() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Stop \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP\033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP\033[0m":
		stoptcp()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func stoptcp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mStopping Reverse Tunnel \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	cmd = exec.Command("systemctl", "stop", "gotunnel-kharej")
	cmd.Run()
	time.Sleep(1 * time.Second)

	cmd = exec.Command("systemctl", "stop", "gotunnel-iran")
	cmd.Run()
	time.Sleep(1 * time.Second)

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mService Stopped!\033[0m")
}
func status() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Status \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP\033[0m":
		tcpStatus()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func tcpStatus() {
	services := []string{"gotunnel-iran", "gotunnel-kharej"}

	fmt.Println("\033[93m            ╔════════════════════════════════════════════╗\033[0m")
	fmt.Println("\033[93m            ║               \033[92mReverse Status\033[93m               ║\033[0m")
	fmt.Println("\033[93m            ╠════════════════════════════════════════════╣\033[0m")

	for _, service := range services {
		cmd := exec.Command("systemctl", "is-active", "--quiet", service)
		err := cmd.Run()
		if err != nil {
			continue
		}

		status := "\033[92m✓ Active      \033[0m"
		displayName := ""
		switch service {
		case "gotunnel-iran":
			displayName = "\033[93mIRAN Server   \033[0m"
		case "gotunnel-kharej":
			displayName = "\033[93mKharej Server \033[0m"
		default:
			displayName = service
		}

		fmt.Printf("           \033[93m ║\033[0m    %s   |    %s\033[93m    ║\033[0m\n", displayName, status)
	}

	fmt.Println("\033[93m            ╚════════════════════════════════════════════╝\033[0m")
}
func UniMenu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Uninstallation \033[96mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mTCP\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mTCP\033[0m":
		removews()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func removews() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving Config ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	deleteCron()
	rmv()

	if _, err := os.Stat("/root/tunnel"); err == nil {
		err := os.RemoveAll("/root/tunnel")
		if err != nil {
			fmt.Printf("Error removing /root/tunnel: %v\n", err)
		}
	}

	azumiServices := []string{
		"gotunnel-iran", "gotunnel-kharej",
	}

	for _, serviceName := range azumiServices {
		hideCmd("systemctl", "disable", serviceName+".service")
		hideCmd("systemctl", "stop", serviceName+".service")
		hideCmd("rm", "/etc/systemd/system/"+serviceName+".service")
	}

	runCmd("systemctl", "daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92m Uninstallation completed!\033[0m")
}
func runCmd(cmd string, args ...string) {
	command := exec.Command(cmd, args...)
	err := command.Run()
	if err != nil {
		fmt.Printf("\033[91mCouldn't run cmd: %s, %v\n\033[0m", cmd, err)
	}
}
func hideCmd(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	nullDevice, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	command.Stdout = nullDevice
	command.Stderr = nullDevice

	return command.Run()
}
var dnldURLs = map[string]string{
	"amd64": "https://github.com/mmatczuk/go-http-tunnel/releases/download/2.1/tunnel_linux_amd64.tar.gz",
	"arm64": "https://github.com/mmatczuk/go-http-tunnel/releases/download/2.1/tunnel_linux_arm.tar.gz",
}

func dwnl(url, filePath string) error {
	cmd := exec.Command("wget", "--quiet", "--show-progress", "-O", filePath, url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func tar(tarGzPath, extractDir string) error {
	cmd := exec.Command("tar", "-xzf", tarGzPath, "-C", extractDir)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func godl() {
	arch := runtime.GOARCH

	dnldURL, ok := dnldURLs[arch]
	if !ok {
		fmt.Printf("\033[91mUnsupported CPU: %s\n\033[0m", arch)
		return
	}

	filePath := "tunnel.tar.gz"
	extractDir := "tunnel"

	err := os.MkdirAll(extractDir, 0755)
	if err != nil {
		fmt.Printf("\033[91mCouldn't create dir: %s\n\033[0m", err.Error())
		return
	}

	err = dwnl(dnldURL, filePath)
	if err != nil {
		fmt.Printf("\033[91mCouldn't download the binary: %s\n\033[0m", err.Error())
		return
	}

	err = tar(filePath, extractDir)
	if err != nil {
		fmt.Printf("\033[91mCouldn't extract the file: %s\n\033[0m", err.Error())
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("\033[91mCouldn't remove the file: %s\n\033[0m", err.Error())
		return
	}

	displayCheckmark("\033[92mDownloaded and Extracted!\033[0m")
}
const crontabFilePath = "/var/spool/cron/crontabs/root"

func resKharej() {
	deleteCron()
	if _, err := os.Stat("/etc/http.sh"); err == nil {
		os.Remove("/etc/http.sh")
	}

	file, err := os.Create("/etc/http.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("sudo systemctl daemon-reload\n")
	file.WriteString("sudo systemctl restart gotunnel-kharej\n")
	file.WriteString("sudo journalctl --vacuum-size=1M\n")

	cmd := exec.Command("chmod", "+x", "/etc/http.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mReset timer\033[93m (minutes):\033[0m ")
	minutesStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("\033[91mError reading input: %v\033[0m", err)
	}
	minutesStr = strings.TrimSpace(minutesStr)
	fmt.Println("╰──────────────────────────────────────╯")

	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		log.Fatalf("\033[91mInvalid input for reset timer:\033[0m %v", err)
	}

    var cronEntry string
    cronEntry = fmt.Sprintf("*/%d * * * * /etc/http.sh", minutes)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh... Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}

	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
func resIran() {
	deleteCron()
	if _, err := os.Stat("/etc/http.sh"); err == nil {
		os.Remove("/etc/http.sh")
	}

	file, err := os.Create("/etc/http.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("sudo systemctl daemon-reload\n")
	file.WriteString("sudo systemctl restart gotunnel-iran\n")
	file.WriteString("sudo journalctl --vacuum-size=1M\n")

	cmd := exec.Command("chmod", "+x", "/etc/http.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mReset timer\033[93m (minutes):\033[0m ")
	minutesStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("\033[91mError reading input: %v\033[0m", err)
	}
	minutesStr = strings.TrimSpace(minutesStr)
	fmt.Println("╰──────────────────────────────────────╯")

	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		log.Fatalf("\033[91mInvalid input for reset timer:\033[0m %v", err)
	}

    var cronEntry string
    cronEntry = fmt.Sprintf("*/%d * * * * /etc/http.sh", minutes)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh... Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}

	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
