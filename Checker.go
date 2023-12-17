package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "time"
)

const (
    Vanity       =   ""
    Client       =   ""
    GuildID      =   ""
    Millisecond  =    100
)

func sendDiscordRequest() {
    for {
        url := fmt.Sprintf("https://discord.com/api/v9/invites/%s?with_counts=true&with_expiration=true", Vanity)
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Println("[Request Hatası:]", err)
            time.Sleep(Millisecond)
            continue
        }

        req.Header.Set("accept", "*/*")
        req.Header.Set("accept-language", "tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7")
        req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"")
        req.Header.Set("sec-ch-ua-mobile", "?0")
        req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
        req.Header.Set("sec-fetch-dest", "empty")
        req.Header.Set("sec-fetch-mode", "cors")
        req.Header.Set("sec-fetch-site", "same-origin")
        req.Header.Set("x-debug-options", "bugReporterEnabled")
        req.Header.Set("x-discord-locale", "tr")
        req.Header.Set("x-fingerprint", "1106586137847410698.uJEFGnNq74j9WcA25zdQPaKQybE")
        req.Header.Set("referrer", fmt.Sprintf("https://discord.com/invite/%s", Vanity))
        req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
        req.Header.Set("credentials", "include")

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            fmt.Println("[Request Hatası:]", err)
            time.Sleep(Millisecond)
            continue
        }

        var data map[string]interface{}
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
            fmt.Println("[Response Hatası:]", err)
        }

        fmt.Println(data)


        if code, ok := data["code"].(float64); ok && data["message"] == "Bilinmeyen Davet" && code == 10006 {
            change()
            fmt.Println("[URL:] Boşa Düştü")
        }

        resp.Body.Close()
        time.Sleep(Millisecond)
    }
}

func change() {
    guildID := GuildID
    newURL := Vanity

    url := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/vanity-url", guildID)
    payload := fmt.Sprintf(`{"code": "%s"}`, newURL)

    req, err := http.NewRequest("PATCH", url, strings.NewReader(payload))
    if err != nil {
        fmt.Println("[Request Hatası:]", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", Client)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("[Request Hatası:]", err)
        return
    }

    var result map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        fmt.Println("[Response Dekod Hatası:]", err)
    }

    fmt.Println(result)

    if result != nil && result["code"] == newURL {
        fmt.Println("[URL:] Alındı")
    } else {
        fmt.Println("[URL:] Alınamadı!")
    }

    resp.Body.Close()
}

func main() {
    sendDiscordRequest()
}
