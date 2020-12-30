# linux auto update
![Go](https://github.com/Derbyshire10/linux-auto-update/workflows/Go/badge.svg?branch=main)
## Getting Started

### Install go(lang)
These instructions will get linux auto update up and running. You will need to have go(lang) set up on your linux machine to compile the project to a binary file.

with [apt](http://packages.qa.debian.org/a/apt.html)-get:
```
sudo apt-get install golang
```
[install Golang manually](https://golang.org/doc/install)

### Compile to binary
On your linux machine cd into the cloned folder and build with go.
```
cd linux-auto-update
go build
```
### Run automatically
Enter the crontab editor with
```
crontab -e
```
And run the file when and how often you like. For example daily at 3am
```
0 3 * * * /root/linux-auto-update/linux-auto-update
```
or use [Crontab Generator](https://crontab-generator.org) to create your own.

## License

This project has been released using Unlicense License, this means it is free for all to use and build upon this project, without permission or fee.
