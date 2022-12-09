package main

import (
    "fmt"
    "github.com/getlantern/systray"
    "os/exec"
    "widetools/icon"
)

var vpnConnected = false

func main() {
    onExit := func() {}
    systray.Run(onReady, onExit)
}

func onReady() {
    systray.SetIcon(icon.Data)
    systray.SetTitle("WideTools")
    systray.SetTooltip("WideTools")

    vpnStart := systray.AddMenuItem("VPN Start", "")
    vpnStop := systray.AddMenuItem("VPN Stop", "")

    // We can manipulate the systray in other goroutines
    go func() {
        for {
            select {
            case <-vpnStart.ClickedCh:
                cmd := exec.Command("D:\\bin\\vpn.bat", "connect")
                // cmd.Run()
                b, _ := cmd.CombinedOutput()
                // 移除 -ldflags "-H=windowsgui"
                fmt.Println(string(b), cmd.String())
                vpnConnected = true
            case <-vpnStop.ClickedCh:
                cmd := exec.Command("D:\\bin\\vpn.bat", "disconnect")
                b, _ := cmd.CombinedOutput()
                fmt.Println(string(b), cmd.String())
                vpnConnected = false
            }
        }
    }()

    systray.AddSeparator()
    quitApp := systray.AddMenuItem("退出", "")
    go func() {
        <-quitApp.ClickedCh
        if vpnConnected {
            cmd := exec.Command("D:\\bin\\vpn.bat", "disconnect")
            b, _ := cmd.CombinedOutput()
            fmt.Println(string(b), cmd.String())
        }
        systray.Quit()
    }()
}
