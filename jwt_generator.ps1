# generate_jwt_secret.ps1

# Generate a secure random key
$bytes = New-Object Byte[] 64
[Security.Cryptography.RNGCryptoServiceProvider]::Create().GetBytes($bytes)
$secret = [Convert]::ToBase64String($bytes)

# Path to .env file (create if it doesn't exist)
$envFilePath = ".\.env"
if (-not (Test-Path $envFilePath)) {
    Copy-Item ".\example.env" $envFilePath
}

# Read the current .env file
$content = Get-Content $envFilePath -Raw

# Replace JWT_SECRET line
$pattern = '(?m)^JWT_SECRET=.*$'
$replacement = "JWT_SECRET=$secret"

if ($content -match $pattern) {
    # Update existing line
    $content = $content -replace $pattern, $replacement
} else {
    # Add new line if not found
    $content += "`nJWT_SECRET=$secret"
}

# Write back to the file
Set-Content -Path $envFilePath -Value $content

Write-Host "JWT secret updated in .env file"