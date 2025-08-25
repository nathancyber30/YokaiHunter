YokaiHunter adalah tool kombinasi `Gospider`, `Katana`, dan `Dalfox` untuk melakukan crawling URL + parameter discovery + XSS auto scan. Dibuat dengan 💖 oleh bug hunter 👨‍💻 Gen-Z.

---

✨ Fitur
- 🔍 Crawling cepat pakai Gospider dan Katana
- 📎 Filter otomatis URL yang mengandung parameter (`?`)
- ⚔️ Scan XSS dengan Dalfox
- 📊 Progress bar interaktif
- 🟩 Output warna-warni (biar nggak bosen)
- 📁 Simpan output ke file `.txt`

---

🚀 Cara Pakai

 1. Clone atau upload file
```bash
git clone https://github.com/nathancyber30/yokaihunter.git
cd yokaihunter
ls
# Pastikan jalanin ini, dan kalo error mungkin perlu: chmod +x install.sh && sudo ./install.sh
chmod +x install.sh
bash install.sh

atau kalo butuh sudo
sudo bash install.sh 

# lalu jalankan hunting-nya:
./yokaihunter -u http://testphp.vulnweb.com -tool all


⚠️ Requirements

	Go minimal versi 1.18+

	Dalfox harus sudah ter-install:
	👉  https://github.com/hahwul/dalfox.git
