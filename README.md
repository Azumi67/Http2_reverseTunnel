**بقیه اموزش فردا نوشته میشود. بعد از این، اسکریپت gre sit ipip را کامل تر میکنم**

![R (2)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/a064577c-9302-4f43-b3bf-3d4f84245a6f)
نام پروژه : تانل Reverse Http2 TCP
---------------------------------------------------------------

![check](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/13de8d36-dcfe-498b-9d99-440049c0cf14)
**امکانات**


- پشتیبانی از TCP
- مناسب برای استفاده شخصی با سرعت خوب
- قابلیت تانل بر روی تک پورت و چندین پورت
- امکان استفاده از ایپی فیلتر شده با ایپی 4
- تنظیم ریست تایمر به دقیقه
- ایجاد سرویس برای تمامی گزینه ها
- امکان حذف تمامی تانل ها و سرویس ها

-------------------
 <div align="right">
  <details>
    <summary><strong>توضیحات</strong></summary>
  
------------------------------------ 

- من این تانل را مثل بقیه تانل ها، تست های زیادی گرفتم و به این نتیجه رسیدم برای استفاده شخصی، سرعت مناسبی را در اختیار شما قرار میدهد.
- این تانل یک مشکل دارد و آن این است که اگر کانکشن فعالی بر روی آن نباشد، کانکشن و کلاینت id دراپ میشود و کانکشن قطع خواهد شد. برای همین یک گزینه به نام reset timer قرار دادم که مواقعی که خواب هستید ریست تایمر را بر روی 1 دقیقه بگذارید و مواقعی که از ان استفاده میکنید، ریست تایمر را بر روی 30 دقیقه بگذارید.( من خودم بر روی یک دقیقه گذاشتم و مشکلش ممکنه قطعی وصلی مقطعی شما به اندازه یک ثانیه باشد)
- برای همین این تانل برای مصرف شخصی، گشت گذار در اینستاگرام و یوتیوب خوب است.
- برای این تانل باید optimizer و وارپ وایرگارد را فعال کنید تا سرعت قابل توجهی داشته باشید.

  </details>
</div>

--------------
  <div align="right">
  <details>
    <summary><strong><img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/71b80a34-9515-42de-8238-9065986104a1" alt="Image"> اموزش نصب go مورد نیاز برای اجرای اسکریپت</strong></summary>
  
------------------------------------ 

- شما میتوانید از طریق اسکریپت [Here](https://github.com/Azumi67/UDP2RAW_FEC#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%85%D9%86) ، این پیش نیاز را نصب کنید یا به صورت دستی نصب نمایید.
- لطفا پس از نصب پیش نیاز ، برای اجرای اسکریپت go برای بار اول، ممکن تا 10 ثانیه طول بکشد اما بعد از آن سریع اجرا میشود.
- یا به صورت دستی :
```
sudo apt update
arm64 : wget https://go.dev/dl/go1.21.5.linux-arm64.tar.gz
arm64 : sudo tar -C /usr/local -xzf go1.21.5.linux-arm64.tar.gz

amd64 : wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
amd64 : sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

nano ~/.bash_profile
paste this into it : export PATH=$PATH:/usr/local/go/bin
save and exit with Ctrl + x , then Y

source ~/.bash_profile
go mod init mymodule
go mod tidy
go get github.com/AlecAivazis/survey/v2
go get github.com/fatih/color

```
- سپس اسکریپت را میتوانید اجرا نمایید.
  </details>
</div>

--------------


![147-1472495_no-requirements-icon-vector-graphics-clipart](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/98d8c2bd-c9d2-4ecf-8db9-246b90e1ef0f)
 **پیش نیازها**

 - لطفا سرور اپدیت شده باشه.
 - فعال کردن وارپ وایرگارد و routing برای سرعت بیشتر و اختلال کمتر در اینستاگرام و یوتیوب
 - میتوانید از اسکریپت اقای [Hwashemi](https://github.com/hawshemi/Linux-Optimizer) و یا [OPIRAN](https://github.com/opiran-club/VPS-Optimizer) هم برای بهینه سازی سرور در صورت تمایل استفاده نمایید. 



----------------------------

  
  ![6348248](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/398f8b07-65be-472e-9821-631f7b70f783)
**آموزش**
-
 <div align="right">
  <details>
    <summary><strong><img src="https://github.com/Azumi67/Rathole_reverseTunnel/assets/119934376/fcbbdc62-2de5-48aa-bbdd-e323e96a62b5" alt="Image"> </strong>ریورس تانل tcp</summary>
  
  
------------------------------------ 


![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران** 

**مسیر : IPV4 TCP > IRAN**




 <p align="right">
  <img src="https://github.com/Azumi67/Http2_reverseTunnel/assets/119934376/03670ca6-38dd-4531-947b-cd5e4d44a672" alt="Image" />
</p>

- سرور ایران را کانفیگ میکنیم
- پورت تانل را 5050 وارد میکنم
- پورت Https را 443 وارد میکنم. شما میتوانید پورت های دیگر Https را وارد کنید.
- پورت Http را 80 وارد میکنم. شما میتوانید پورت های دیگر http را وارد نمایید.
- دقت نمایید که این پورت ها درگیر نباشد.
- ریست تایمر را یک دقیقه میگذارم چون استفاده من از سرور به صورت مداوم نیست و امکان drop connection هست.
- بعدا در menu امکان تغییر ریست تایم هست.

------------------------------------ 

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج**

**مسیر : IPV4 TCP > Kharej**


 <p align="right">
  <img src="https://github.com/Azumi67/Http2_reverseTunnel/assets/119934376/f626edab-69b1-4731-b125-7a5cec0db8d9" alt="Image" />
</p>


- سرور خارج را کانفیگ میکنیم
- ایپی 4 ایران را وارد میکنم و مهم نیست فیلتر هست یا خیر
- تعداد کانفیگ را عدد 1 وارد میکنم چون تنها یک کانفیگ دارم
- پورت تانل را 5050 قرار میدم
- پورت کانفیگ را 8080 قرار میدم
- ریست تایمر را 1 دقیقه میذارم چون سرور ایران هم یک دقیقه گذاشتم و دلیلش هم بالاتر گفتم.
</details>
</div>
 <div align="right">
  <details>
    <summary><strong><img src="https://github.com/Azumi67/Rathole_reverseTunnel/assets/119934376/fcbbdc62-2de5-48aa-bbdd-e323e96a62b5" alt="Image"> </strong>ویرایش ریست تایمر</summary>
  
  
------------------------------------ 

<p align="right">
  <img src="https://github.com/Azumi67/Http2_reverseTunnel/assets/119934376/79314279-5602-4171-aaff-7b7aa1c8b461" alt="Image" />
</p>

- به راحتی زمان جدید را به تانل اضافه کنید.
  </details>
</div>

<div align="right">
  <details>
    <summary><strong><img src="https://github.com/Azumi67/Rathole_reverseTunnel/assets/119934376/fcbbdc62-2de5-48aa-bbdd-e323e96a62b5" alt="Image"> </strong>نصب وارپ وایرگارد در پنل علیرضا</summary>


  <p align="right">
  <img src="https://github.com/Azumi67/Http2_reverseTunnel/assets/119934376/41b5e128-9a4c-4c20-8f27-2d1a500961a5" alt="Image" />
</p>

 <p align="right">
  <img src="https://github.com/Azumi67/Http2_reverseTunnel/assets/119934376/4e079a13-8459-4a64-a436-2c2083ebf0cb" alt="Image" />
</p>

- از قسمت xray setting و warp routing ، وارپ را فعال میکنم. create را بزنید و گزینه add outbound را کلیک و تمامی گزینه ها را فعال کنید.

- سپس داخل تب advanced و تب all میتوانید اینها را اضافه کنید . فقط به جای secrect و سایر موارد؛ مقادیر خود را قرار بدید( مقادیر SECRET KEY و سایر موارد توسط پنل شما 

```
{
  "api": {
    "services": [
      "HandlerService",
      "LoggerService",
      "StatsService"
    ],
    "tag": "api"
  },
	  "routing": {
    "domainStrategy": "AsIs",
    "rules": [
      {
        "inboundTag": [
          "api"
        ],
        "outboundTag": "api",
        "type": "field"
      },
      {
        "ip": [
          "geoip:private"
        ],
        "outboundTag": "blocked",
        "type": "field"
      },
      {
        "outboundTag": "blocked",
        "protocol": [
          "bittorrent"
        ],
        "type": "field"
      },
      {
        "type": "field",
        "outboundTag": "warp",
        "domain": [
          "geosite:openai",
          "geosite:netflix",
          "geosite:spotify",
          "geosite:google",
		  "geosite:microsoft",
          "geosite:youtube",
          "geosite:meta"
        ]
      }
    ]
  },
  "inbounds": [
    {
      "listen": "127.0.0.1",
      "port": 62789,
      "protocol": "dokodemo-door",
      "settings": {
        "address": "127.0.0.1"
      },
      "tag": "api"
    }
  ],
  "log": {
    "loglevel": "warning"
  },
    "dns": {
        "servers": [
            "https://1.1.1.1/dns-query"
        ],
        "queryStrategy": "UseIP"
    },
  "outbounds": [
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    },
    {
      "protocol": "blackhole",
      "settings": {},
      "tag": "blocked"
    },
    {
      "tag": "warp",
      "protocol": "wireguard",
      "settings": {
        "mtu": 1280,
        "DNS": "1.1.1.1 1.0.0.1",
        "secretKey": "YOUR SECRET KEY اینجا",
        "address": [
          "172.16.0.2",
          "YOUR WIREGUARD IPV6 ایپی 6 وایرگارد اینجا"
        ],
        "workers": 2,
        "peers": [
          {
            "publicKey": "پابلیک کی شما",
            "allowedIPs": [
              "0.0.0.0/0",
              "::/0"
            ],
            "endpoint": "engage.cloudflareclient.com:2408",
            "keepAlive": 25
          }
        ],
        "kernelMode": false
      }
    }
  ],
  "policy": {
    "levels": {
      "0": {
        "statsUserDownlink": true,
        "statsUserUplink": true
      }
    },
    "system": {
      "statsInboundDownlink": true,
      "statsInboundUplink": true
    }
  },
  "stats": {}
}
```
- حتما مقادیر خودتان را جایگذاری کنید. سپس میتوانید اینباند های خود را بسازید.
  </details>
</div>
