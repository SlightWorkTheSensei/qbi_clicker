# qbi_clicker – Reliable Autoclicker in GoLang

**qbi_clicker** is a bare-metal fast and reliable autoclicker built in **GoLang** for tasks where high uptime and consistency are critical — like trading UIs, bot defenses, or any high-frequency GUI automation.

Built specifically because Python clickers like `pyautogui` tend to sleep, freeze, or miss clicks over time.

---

## 🧠 What It Does

This tool continuously clicks on a predefined pixel coordinate (i.e., "the dot") as long as the program is running. It was designed to handle:

- GUI automation for financial apps
- Keeping live UIs responsive (e.g., to keep charts updating)
- Long-running bot tasks without crash risk

---

## ⚡ Why GoLang?

Python clickers were:
- Going idle after minutes
- Getting deprioritized by OS threads
- Inconsistent with millisecond timing

GoLang solves that:
- Compiles to a fast, native binary
- Low CPU overhead
- Very stable under high-duty cycles

---

## 🔧 Setup

### 🔨 Install Dependencies (if using packages)
Uses `github.com/go-vgo/robotgo` (or your preferred mouse automation lib):

```bash
go get github.com/go-vgo/robotgo
