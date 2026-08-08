# notify-send

A program to send desktop notifications on WSL2.

## Install

```bash
GOOS=windows go install github.com/shibataka000/notify-send@main
```

## Usage

```
a program to send desktop notifications

Usage:
  notify-send <summary> [body] [flags]

Flags:
  -a, --app-name string   Specifies the app name for the notification. (default "notify-send")
  -h, --help              help for notify-send
  -i, --icon string       Specifies an icon filename to display.
```
