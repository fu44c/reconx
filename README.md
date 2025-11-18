
# ReconX

![ReconX Logo](./assets/logo.png)  <!-- لو عندك شعار -->

**ReconX** هي أداة متقدمة لأغراض الـ **Bug Bounty & Recon**، تشمل جميع خطوات جمع الدومينات، فحص الثغرات، واستغلالها بطريقة منظمة.  
تم تطوير المشروع بواسطة **Ahmed Ibrahim**.

---

## المميزات

- جمع الدومينات من مصادر متعددة.  
- اختبار البروتوكولات HTTP/HTTPS.  
- فحص الثغرات الشائعة (SQLi, XSS, LFI, …).  
- دعم التشغيل على Bash وGo (نسخة قوية).  
- واجهة سهلة الاستخدام مع عرض ASCII Art جذاب.  

---

## المتطلبات

- Git  
- Go >= 1.20 (للنسخة المحولة لـ Go)  
- Python 3.10+ (للنسخة Bash)  
- أدوات مساعدة: `httpx`, `subfinder`, `nmap`, `ghauri`  

---

## تثبيت الأدوات المساعدة

```bash
# تحديث النظام
sudo apt update && sudo apt upgrade -y

# تثبيت Git و Python
sudo apt install git python3 python3-pip -y

# تثبيت Go
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# تثبيت الأدوات المساعدة
go install github.com/projectdiscovery/httpx/cmd/httpx@latest
go install github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
pip3 install ghauri


كيفية تحميل المشروع وتشغيله
التحميل:
git clone https://github.com/fu44c/reconx.git
cd reconx

التشغيل:
النسخة Bash:
bash reconx.sh

النسخة Go:
go run reconx.go


ستظهر لك واجهة ASCII Art و قائمة بالخيارات، اختر المطلوب بسهولة.


استخدام الأداة


جمع الدومينات:


./reconx --collect-domains -d example.com



فحص البروتوكولات:


./reconx --scan-http -d example.com



فحص الثغرات:


./reconx --vuln-check -d example.com


الأداة تدعم التشغيل الآلي لكل الخطوات دفعة واحدة.


هيكل المشروع (Architecture)
ReconX/
├── reconx.sh        # نسخة Bash
├── reconx.go        # نسخة Go
├── modules/
│   ├── domain_collector.go
│   ├── http_scanner.go
│   └── vuln_checker.go
├── assets/
│   └── logo.png
├── README.md
└── LICENSE

الموديولات:


domain_collector: جمع الدومينات من مصادر مختلفة.


http_scanner: فحص HTTP/HTTPS وفلترة المواقع الحية.


vuln_checker: فحص SQLi, XSS, LFI، وغيرها.



المساهمة


عمل Fork للمشروع.


إنشاء فرع جديد:


git checkout -b feature-name



عمل Commit و Push للفرع.


عمل Pull Request.



الترخيص
تم اعتماد ترخيص MIT License. يمكنك استخدام المشروع للأغراض التعليمية و Bug Bounty مع الإشارة للمطور.

الدعم
للاستفسار عن المشروع أو طلب تطوير جديد، تواصل مع Ahmed Ibrahim عبر GitHub.

---

إذا تحب، أقدر أسوي لك **نسخة جاهزة بصيغة Markdown** مع **ASCII Art كبير وجذاب** في الأعلى ويكون جاهز مباشرة للـ GitHub.  
هل أبغاك أسويها لك؟
