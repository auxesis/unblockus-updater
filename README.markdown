# Unblock-Us Updater

Ping the (unofficial) Unblock-Us API to associate your current IP address with your account.

Implemented in Go.

## Usage

``` bash
./updater <username> <password>
```

Example:

``` bash
./updater me@domain.example s3cr3t
```

## Building

Ensure you have Go installed, then run:

``` bash
git clone https://github.com/auxesis/unblockus-updater.git
cd unblockus-updater
go build updater.go
```

To develop:

``` bash
go run updater.go me@domain.example s3cr3t
```
