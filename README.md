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

<img src="https://img.shields.io/badge/Version-2.0.0-informational?style=for-the-badge&logo=none" alt="version" />
</p>

# Installation
```bash
go install github.com/mehditeymorian/jwt
```

Commands are as below:
- **config**: view, edit, and set config
- **encode**: create standard JWT token
- **decode**: decode a JWT token
- **gen**: generate different keys such as rsa, hmac, ecdsa
- 
```bash
# run the following command to see the full details.
jwt help
```

# Configuration
The `JWT` use corresponding set of keys when you use different commands. For example, it will use rsa keys if you use command `jwt gen rsa`.
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

> Note: if `interactive` config set to true, the command parameters will be taken from user interactively instead of reading them from command options.

> use `-c` option to pass config file address.

## Set Config
The order of reading config is as follows if config is not specified as option:
1. `.jwt.(yaml|yml)` file in the path where the `JWT` is running.
2. `/etc/jwt/.jwt.(yaml|yml)` file as default configuration. if config file is not present, one will be created.


# Contribution
Any contribution in any form is welcomed. Open an issue to discuss it.

# Contact
- [Email](mailto:mehditeymorian322@gmail.com)

# License
Unless otherwise noted, the JWT source files are distributed under the Apache Version 2.0 license found in the LICENSE file.
