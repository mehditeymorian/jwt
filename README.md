<h1 align="center">
<img alt="Koi logo" src="assets/icon.png" width="500px"/><br/>
JWT CLI
</h1>
<p align="center">Encode & Decode JWT Tokens</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/jwt?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

<img src="https://img.shields.io/badge/Version-1.1.1-informational?style=for-the-badge&logo=none" alt="version" />
</p>

# Installation
```bash
go install github.com/mehditeymorian/jwt
```

# Commands
[![asciicast](https://asciinema.org/a/518698.svg)](https://asciinema.org/a/518698)
```bash
# run to see all commands
jwt
```

# Configuration

Different key will be used for encode or decode base on the token algorithm. The `interactive` field indicates if user input is taken from options or user choose from a set of prompts.
```yaml
interactive: true
rsa:
  public_key: |-
    key
  private_key: |-
    key
hmac:
  key: key
  base64_encoded: false
ecdsa:
  public_key: |-
    key
  private_key: |-
    key
```

## Set Config
1. Use `-c` option to pass config path
2. Put `jwt-config.(yaml|yml)` where you run the jwt cli
3. Default configuration, which is located at `/etc/jwt/config.yaml`


# Contribution
Any contribution in any form is welcomed. Open an issue to discuss about it.

# Contact
- [Email](mailto:mehditeymorian322@gmail.com)

# License
Unless otherwise noted, the JWT source files are distributed under the Apache Version 2.0 license found in the LICENSE file.
