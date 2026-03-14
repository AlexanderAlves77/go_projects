# GoDevOps

A professional CLI and interactive dashboard tool built with Go for monitoring system resources on Windows, Linux, and macOS.


## Features

- Real-time monitoring of CPU, RAM, Disk, Network, and Processes.
- Interactive dashboard with dynamic colored bars (green/yellow/red).
- Top CPU-consuming processes with live updates.
- Keyboard shortcuts:
  - q → Quit
  - r → Refresh dashboard
- Fully cross-platform CLI tool.
- Can be installed globally as a system command.


## Installation
### Requirements

- Go 1.26+
- Windows, Linux, or macOS


## Compile & Install

# Clone the repository
git clone https://github.com/AlexanderAlves77/go_projects.git
cd go_projects/godevops

# Compile
go build -o godevops.exe main.go

# (Optional) Move to a global path on Windows
move godevops.exe C:\Windows\System32


Or use the provided PowerShell install script:
Set-ExecutionPolicy Bypass -Scope Process -Force
.\install-godevops.ps1


## Usage

### Show CPU usage
godevops cpu

### Show RAM usage
godevops memory

###  Show Disk usage
godevops disk

###  Show network statistics
godevops network

###  Show top processes
godevops process

###  Real-time monitor
godevops monitor

###  Interactive dashboard
godevops dashboard


## Keyboard shortcuts in dashboard:

* q → Quit the dashboard
* r → Refresh manually


## Project Structure
```
godevops/
├── cmd/                 # CLI commands
│   ├── cpu.go
│   ├── memory.go
│   ├── disk.go
│   ├── network.go
│   ├── process.go
│   └── dashboard.go
├── internal/
│   ├── system/          # System monitoring functions
│   └── ui/              # CLI UI utilities (bars, colors)
├── main.go              # Entry point
├── install-godevops.ps1 # Windows installer script
├── go.mod
├── go.sum
└── .gitignore
```

## Dependencies

* github.com/spf13/cobra  → CLI commands
* github.com/shirou/gopsutil  → System metrics
* github.com/rivo/tview  → Interactive terminal UI
* github.com/gdamore/tcell  → Terminal handling


## Screenshots
### Dashboard example:
```
FULLDEVSTACKS SYSTEM METRICS
────────────────────────
CPU: █████████░░░░░░ 45.00%
RAM: ██████░░░░░░░░ 72.00%
DISK: ██████████░░░ 75.29%
PROC: 351

NETWORK
DOWNLOAD ↓ 0.00MB/s
UPLOAD ↑ 0.00MB/s

TOP PROCESSES
PROCESS              CPU
────────────────────────
chrome.exe           12.50%
code.exe             8.30%
explorer.exe         5.60%
```

## License

This project is MIT licensed.


## Author
```
Alexander Alves
Computer Science Student
Developer Java | C++ | Go
QA Automation Engineer 
Game Development Enthusiast
```