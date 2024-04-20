# Şehir Veri Servisi

Bu basit Go programı, şehir, ilçe ve mahalle verilerini içeren bir JSON dosyasını okur ve bu verilere HTTP API üzerinden erişim sağlar.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun ve Gin paketinin yüklü olduğundan emin olun.
2. **Veri Dosyası**: `data.json` dosyasını projenizin kök dizinine ekleyin ve şehir, ilçe ve mahalle verilerini içeren bir JSON formatında tutun.
3. **Sunucuyu Başlatın**: `go run main.go` komutuyla sunucuyu başlatın.
4. **Verilere Erişin**: Tarayıcıdan veya bir API test aracından `/allData` veya `/cities/{cityName}` endpoint'lerine bir GET isteği yaparak verilere erişin.

## API Dökümantasyonu

### `/allData` Endpoint'i

- **Method**: GET
- **Cevap**: JSON formatında tüm şehir verilerini içeren bir yanıt döndürür.

Örnek istek:
GET http://localhost:8080/allData
![Açıklama](https://github.com/omerfdev/addressAPI/blob/master/goAddress.png)

Örnek cevap:
```json
[
  {
    "name": "İstanbul",
    "alpha_2_code": "IST",
    "towns": [
      {
        "name": "Beşiktaş",
        "districts": [
          {
            "name": "Abbasağa",
            "quarters": [
              {
                "name": "Cihannüma"
              }
            ]
          }
        ]
      }
    ]
  }
]




