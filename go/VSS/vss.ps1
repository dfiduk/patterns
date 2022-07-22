Param(
    [Parameter(Position=0,
                Mandatory=$true)]
    [string] $VolumeGUID,
    [Parameter(Position=1, 
               Mandatory=$false)]
    [string] $ComputerName = "localhost"    
  )

(gwmi -List Win32_ShadowCopy -ComputerName 'localhost').Create($VolumeGUID + '\', 'ClientAccessible').ShadowId|Write-Host
