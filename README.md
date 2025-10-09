<h1 align="center">asnx</h1>

<h4 align="center">A CLI tool for obtaining information about <a href="https://en.wikipedia.org/wiki/Autonomous_system_(Internet)">ASN</a> hosts using RDAP</h4>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a>
</p>

# Features
![image](docs/img/usage.png)

# Installation

```
Sorry, but now: go build :)
```

# Usage

```
frakenbok@frakenbok$ asnx -h

		▄█████▄  ▄▄█████▄  ██▄████▄  ▀██  ██▀
		▀ ▄▄▄██  ██▄▄▄▄ ▀  ██▀   ██    ████
		▄██▀▀▀██   ▀▀▀▀██▄  ██    ██    ▄██▄
		██▄▄▄███  █▄▄▄▄▄██  ██    ██   ▄█▀▀█▄
		▀▀▀▀ ▀▀   ▀▀▀▀▀▀   ▀▀    ▀▀  ▀▀▀  ▀▀▀

		  You are using asnx version 1.1.0!

	A tool for obtaining information about ASN hosts using RDAP

Usage:
  asnx [flags]

Flags:
  -d, --domains stringArray    Domain names (usage: -d cr4.sh,yandex.ru or --domains domains.txt)
  -h, --help                   help for asnx
  -f, --ip-range-file string   File to save all ip ranges handeled from ASN
  -i, --ips stringArray        IP addresses (usage: -i 127.0.0.1,8.8.8.8 or --ips ips.txt)
      --no-banner              Disables banner
  -o, --output string          File to save output
  -v, --version                Shows asnx version
```

<div align="center">

**asnx** is made by [FrakenboK](https://github.com/FrakenboK) under [MIT License](LICENSE).

</div>