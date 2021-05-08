Write-Host "Downloading hydra..."

$url = "https://github.com/Shravan-1908/hydra/releases/latest/download/hydra-windows-amd64.exe"

$dir = $env:USERPROFILE + "/.hydra"
$filepath = $env:USERPROFILE + "/.hydra/hydra.exe"

[System.IO.Directory]::CreateDirectory($dir)
(Invoke-WebRequest -Uri $url -OutFile $filepath)

Write-Host "Adding hydra to PATH..."
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";"+$dir,
    [EnvironmentVariableTarget]::Machine)

Write-Host 'hydra installation is successfull!'
Write-Host "You need to restart your shell to use hydra."
