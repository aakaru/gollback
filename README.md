# <div align="center">🔄 Gollback</div>

<div align="center">

### **Version Control for n8n Workflows**

<img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=22&duration=3000&pause=1000&color=00ADD8&center=true&vCenter=true&width=600&lines=Git-like+Version+Control+for+n8n;Backup+%7C+Compare+%7C+Restore+Workflows;Built+with+Go+for+Performance;Production-Ready+CLI+Tool" alt="Typing SVG" />

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge)](http://makeapullrequest.com)
[![Made with Love](https://img.shields.io/badge/Made%20with-❤️-red.svg?style=for-the-badge)](https://github.com/yourusername/gollback)

</div>

---

<div align="center">

## 🎯 Problem Statement

</div>

<table>
<tr>
<td width="50%">

### ❌ The Problem

```diff
- No version control for workflows
- Risk of data loss
- Can't track changes
- No automated backups
- Difficult to compare versions
```

</td>
<td width="50%">

### ✅ The Solution

```diff
+ Git-like version control
+ Automatic backups
+ Smart diff comparison
+ One-click restore
+ Beautiful CLI interface
```

</td>
</tr>
</table>

---

<div align="center">

## ✨ Features That Make You Go "WOW"

</div>

<details open>
<summary><b>🔐 Secure Configuration</b></summary>
<br>

- 🔑 Multiple config sources (file, env, defaults)
- 🛡 API keys stored securely in `~/.gollbackrc`
- 🚫 Never commit secrets to Git
- ⚙️ Interactive setup wizard

</details>

<details open>
<summary><b>💾 Intelligent Backup</b></summary>
<br>

- ⏰ Automatic timestamped backups
- 📦 Complete workflow JSON preservation
- 🗂 Organized directory structure
- 🔄 Idempotent operations

</details>

<details open>
<summary><b>🔍 Smart Comparison</b></summary>
<br>

- 🎨 Git-style colored diffs
- 📊 Node count & status tracking
- ⚡️ Lightning-fast comparison
- 🎯 Pinpoint exact changes

</details>

<details open>
<summary><b>♻️ Easy Restore</b></summary>
<br>

- 🚀 One-command restoration
- 🔒 Safe restore (no overwriting)
- 📋 Preserves all connections
- ✅ Validation checks

</details>

---

<div align="center">

## 🚀 Quick Start

</div>

### 📋 Prerequisites

<table>
<tr>
<td align="center" width="33%">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="60" height="60" /><br>
<b>Go 1.21+</b>
</td>
<td align="center" width="33%">
<img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="60" height="60" /><br>
<b>Docker</b>
</td>
<td align="center" width="33%">
<img src="https://n8n.io/favicon.ico" width="60" height="60" /><br>
<b>n8n Instance</b>
</td>
</tr>
</table>

### ⚡️ Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/gollback.git
cd gollback

# Install dependencies
go mod download

# Build the binary
go build -o gollback

# (Optional) Move to PATH
sudo mv gollback /usr/local/bin/
```

### 🐳 Setup n8n with Docker

```bash
# Start n8n
docker compose up -d

# Access n8n at http://localhost:5678
# Create your owner account
# Go to Settings → n8n API → Create API Key
```

### 🔧 Configure Gollback

```bash
# Interactive setup (Recommended)
./gollback init

# Or set environment variable
export GOLLBACK_API_KEY="your-n8n-api-key"
```

---

<div align="center">

## 📖 Usage Examples

</div>

<table>
<tr>
<td width="50%">

### 💾 Backup Workflows

```bash
gollback backup
```

**Output:**
```
🚀 n8n Backup Tool Starting...
📡 Fetching workflows from n8n...
✅ Found 3 workflow(s)
💾 Starting backup...
   Backing up: Email Campaign...
💾 Saved: Email_Campaign_2025-11-03.json
✅ Backup Complete! 3/3 workflows saved
```

</td>
<td width="50%">

### 📋 List All Backups

```bash
gollback list
```

**Output:**
```
📋 Listing all backups...
Found 5 backup file(s):
1. Email_Campaign_2025-11-03.json
   Size: 12.34 KB
   Path: ./backups/Email_Campaign.json
```

</td>
</tr>
<tr>
<td width="50%">

### 🔍 Compare Versions

```bash
gollback diff "Email_Campaign"
```

**Shows colored diff with:**
- ✅ Green for additions
- ❌ Red for deletions
- ⚠️ Yellow for modifications

</td>
<td width="50%">

### ♻️ Restore Workflow

```bash
gollback restore backups/file.json
```

**Output:**
```
🔄 Restoring workflow...
📋 Workflow: Email Campaign
📤 Uploading to n8n...
✅ Workflow restored successfully!
```

</td>
</tr>
</table>

---

<div align="center">

## 🏗 Architecture

</div>

```
gollback/
├── 📂 cmd/              # CLI commands (Cobra)
│   ├── backup.go        # Backup workflows
│   ├── compare.go       # Compare two files
│   ├── diff.go          # Quick compare latest versions
│   ├── init.go          # Configuration setup
│   ├── list.go          # List backups
│   ├── restore.go       # Restore workflows
│   └── root.go          # Root command
├── 📂 api/              # n8n API client
│   └── client.go        # HTTP client for n8n REST API
├── 📂 backup/           # Backup management
│   └── manager.go       # File operations & backup logic
├── 📂 config/           # Configuration
│   ├── config.go        # Config struct
│   └── loader.go        # Viper-based config loader
└── 📄 main.go           # Entry point
```

### 🧠 Key Design Decisions

<table>
<tr>
<td width="50%">

#### Why Go? 🚀

- ⚡️ Fast compilation and execution
- 📦 Single binary distribution
- 🛠 Excellent standard library
- 💪 Strong CLI framework ecosystem

</td>
<td width="50%">

#### Why n8n? 🔗

- 🌐 Open-source workflow automation
- 🏠 Self-hosted with full API access
- 📈 Growing popularity in 2025
- ❌ No built-in version control

</td>
</tr>
</table>

---

<div align="center">

## 🛠 Configuration

</div>

Gollback supports **multiple configuration methods** with intelligent hierarchy:

<table>
<tr>
<th width="33%">1️⃣ Config File</th>
<th width="33%">2️⃣ Environment Variables</th>
<th width="33%">3️⃣ Command Flags</th>
</tr>
<tr>
<td>

```bash
gollback init
```

Creates `~/.gollbackrc`

```yaml
api_key: n8n_api_123
n8n_url: http://localhost:5678
backup_dir: ./backups
```

</td>
<td>

```bash
export GOLLBACK_API_KEY="..."
export GOLLBACK_N8N_URL="..."
export GOLLBACK_BACKUP_DIR="..."
```

Overrides config file

</td>
<td>

```bash
gollback backup \
  --api-key="..." \
  --url="..."
```

Highest priority

</td>
</tr>
</table>

---

<div align="center">

## 🧪 Development

</div>

### Running Tests

```bash
go test ./...
```

### Building

```bash
# Development build
go build -o gollback

# Production build with optimizations
go build -ldflags="-s -w" -o gollback

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o gollback-linux
GOOS=darwin GOARCH=amd64 go build -o gollback-mac
GOOS=windows GOARCH=amd64 go build -o gollback.exe
```

---

<div align="center">

## 📊 Project Stats

![GitHub stars](https://img.shields.io/github/stars/yourusername/gollback?style=social)
![GitHub forks](https://img.shields.io/github/forks/yourusername/gollback?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/yourusername/gollback?style=social)

</div>

---

<div align="center">

## 🤝 Contributing

</div>

Contributions are **welcome**! Please feel free to submit a Pull Request.

<table>
<tr>
<td align="center" width="20%">
<b>1️⃣</b><br>
Fork the repo
</td>
<td align="center" width="20%">
<b>2️⃣</b><br>
Create branch
</td>
<td align="center" width="20%">
<b>3️⃣</b><br>
Make changes
</td>
<td align="center" width="20%">
<b>4️⃣</b><br>
Push changes
</td>
<td align="center" width="20%">
<b>5️⃣</b><br>
Open PR
</td>
</tr>
</table>

---

<div align="center">

## 🔮 Future Enhancements

</div>

<table>
<tr>
<td width="50%">

### Phase 1 🚀
- [ ] Scheduled automatic backups
- [ ] Git integration
- [ ] Backup encryption
- [ ] Cloud storage (S3, GCS)

</td>
<td width="50%">

### Phase 2 🌟
- [ ] Webhook notifications
- [ ] Web UI dashboard
- [ ] Workflow templates
- [ ] Multi-tenant support

</td>
</tr>
</table>

---

<div align="center">

## 🙏 Acknowledgments

</div>

<table>
<tr>
<td align="center" width="20%">
<a href="https://n8n.io/"><img src="https://n8n.io/favicon.ico" width="40" height="40" /></a><br>
<b>n8n</b>
</td>
<td align="center" width="20%">
<a href="https://github.com/spf13/cobra"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" height="40" /></a><br>
<b>Cobra</b>
</td>
<td align="center" width="20%">
<a href="https://github.com/spf13/viper"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" height="40" /></a><br>
<b>Viper</b>
</td>
<td align="center" width="20%">
<a href="https://github.com/nsf/jsondiff"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" height="40" /></a><br>
<b>jsondiff</b>
</td>
<td align="center" width="20%">
<a href="https://github.com/fatih/color"><img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="40" height="40" /></a><br>
<b>color</b>
</td>
</tr>
</table>

---

<div align="center">

## 📧 Contact

**Follow me on:** • [Twitter](https://twitter.com/RageCreates) • [LinkedIn](https://linkedin.com/in/aakriti-kaushik/)

**Project Link:** [https://github.com/aakaru/gollback](https://github.com/aakaru/gollback)

---

### Show some ❤️ by starring ⭐️ this repository!

<img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=16&pause=1000&color=00ADD8&center=true&vCenter=true&width=435&lines=Built+with+%E2%9D%A4%EF%B8%8F+using+Go;Star+%E2%AD%90+if+you+find+it+useful!" alt="Footer" />
<img src="https://capsule-render.vercel.app/api?type=waving&color=gradient&customColorList=6,11,20&height=100&section=footer&fontSize=40" width="100%"/>

</div>

---

<div align="center">



</div>
