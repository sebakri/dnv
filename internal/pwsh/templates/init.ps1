using namespace System;
using namespace System.Management.Automation;

$env:DNV_SHELL = "pwsh";
$env:DNV_DEBUG = "{{.Debug | ternary true false}}";
$env:DNV_SESSION_ID = [System.Diagnostics.Process]::GetCurrentProcess().Id;
$env:DNV_SESSION_FOLDER = Join-Path $([System.IO.Path]::GetTempPath()) "dnv";

$hook = [EventHandler[LocationChangedEventArgs]] {
  param([object] $source, [LocationChangedEventArgs] $eventArgs)
  end {
    try {
      $unloadCmd = $({{ .Command }} unload | Out-String)
      if($unloadCmd -ne "") {
        Invoke-Expression $unloadCmd;
      }

      $loadCmd = $({{ .Command }} load | Out-String)
      if($loadCmd -ne "") {
        Invoke-Expression $loadCmd;
      }
    }
    catch {
      Write-Debug $_.Exception.Message;
    }
  }
};

$currentAction = $ExecutionContext.SessionState.InvokeCommand.LocationChangedAction;
if ($currentAction) {
  $ExecutionContext.SessionState.InvokeCommand.LocationChangedAction = [Delegate]::Combine($currentAction, $hook);
}
else {
  $ExecutionContext.SessionState.InvokeCommand.LocationChangedAction = $hook;
};

function cleanupDNV {
  Write-Debug "Cleaning up DNV environment";
  Invoke-Expression "{{ .Command }} clean";
}

Register-EngineEvent PowerShell.Exiting -Action { cleanupDNV } -SupportEvent;