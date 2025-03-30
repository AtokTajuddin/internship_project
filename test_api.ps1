#
# test-api.ps1

# 0. Registrasi User Baru
try {
    $registerPayload = @{
        full_name = "Beta Project"
        email     = "betaproject@example.com"
        phone     = "08123456789"
        password  = "Th1s1s_secure_p4assw0rd"
    } | ConvertTo-Json

    Write-Host "Mencoba registrasi user..."
    
    $registerResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/register" `
        -Method POST `
        -ContentType "application/json" `
        -Body $registerPayload `
        -UseBasicParsing `
        -ErrorAction Stop

    Write-Host "Registrasi Berhasil! Response:"
    $registerResponse.Content | ConvertFrom-Json | Format-List
}
catch {
    Write-Host "Gagal registrasi! Error: $($_.Exception.Message)"
    
    # Handle error khusus konflik (user sudah ada)
    if ($_.Exception.Response.StatusCode -eq 409) {
        Write-Host "User sudah terdaftar, melanjutkan ke login..."
    }
    else {
        exit
    }
}

$loginResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/login" `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"email":"betaproject@example.com","password":"Th1s1s_secure_p4assw0rd"}' `
    -UseBasicParsing

# Convert response dari JSON
$loginData = $loginResponse.Content | ConvertFrom-Json
$token = $loginData.token

Write-Host "Token: $token"

# 2. Buat produk dengan token
$productResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/products" `
    -Method POST `
    -ContentType "application/json" `
    -Headers @{"Authorization"="Bearer $token"} `
    -Body '{"name":"ProdukPakaianLebaran ","price":200,"stock":100,"category_id":2,"store_id":2}' `
    -UseBasicParsing

# Tampilkan respons pembuatan produk
Write-Host "Product Creation Response:"
Write-Host $productResponse.Content

# 3. Ambil daftar produk
$productsResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/products" `
    -Method GET `
    -UseBasicParsing

# Tampilkan daftar produk
Write-Host "Products List:"
Write-Host $productsResponse.Content




# Metode register user baru 

# 0. Registrasi User Baru
# try {
#     $registerPayload = @{
#         full_name = "Beta Project"
#         email     = "betaproject@example.com"
#         phone     = "08123456789"
#         password  = "Th1s1s_secure_p4assw0rd"
#     } | ConvertTo-Json

#     Write-Host "Mencoba registrasi user..."
    
#     $registerResponse = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/auth/register" `
#         -Method POST `
#         -ContentType "application/json" `
#         -Body $registerPayload `
#         -UseBasicParsing `
#         -ErrorAction Stop

#     Write-Host "Registrasi Berhasil! Response:"
#     $registerResponse.Content | ConvertFrom-Json | Format-List
# }
# catch {
#     Write-Host "Gagal registrasi! Error: $($_.Exception.Message)"
    
#     # Handle error khusus konflik (user sudah ada)
#     if ($_.Exception.Response.StatusCode -eq 409) {
#         Write-Host "User sudah terdaftar, melanjutkan ke login..."
#     }
#     else {
#         exit
#     }
# }
