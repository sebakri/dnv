using namespace System;
using namespace System.Management.Automation;

$env:DNV_SHELL = "pwsh";
$env:DNV_DEBUG = "{{.Debug | ternary true false}}";
$env:DNV_SESSION_ID = [System.Diagnostics.Process]::GetCurrentProcess().Id;
$env:DNV_SESSION_FOLDER = Join-Path $([System.IO.Path]::GetTempPath()) "dnv";

$sessionFolder = $env:DNV_SESSION_FOLDER;
$unloadScript = Join-Path $sessionFolder $env:DNV_SHELL-$env:DNV_SESSION_ID-"unload.ps1";
$loadScript = Join-Path $sessionFolder $env:DNV_SHELL-$env:DNV_SESSION_ID-"load.ps1";

$hook = [EventHandler[LocationChangedEventArgs]] {
  param([object] $source, [LocationChangedEventArgs] $eventArgs)
  end {
    try {
      if (Test-Path $unloadScript) {
        if ($eventArgs.NewPath -notlike "$env:DNV_ENV_LOADED*") {
          Write-Debug "Unloading environment from $($env:DNV_ENV_LOADED)";
          Invoke-Expression $unloadScript;
          Remove-Item $unloadScript;
        }
      }

      Invoke-Expression "{{ .Command }} generate $($eventArgs.NewPath)";

      if (Test-Path $loadScript) {
        $env:DNV_ENV_LOADED = $eventArgs.NewPath;
        Write-Debug "Loading environment from $($env:DNV_ENV_LOADED)";
        Invoke-Expression $loadScript;
        Remove-Item $loadScript;
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
  removeScript $unloadScript
  removeScript $loadScript
}

function removeScript {
  param([string] $scriptPath)
  try {
    Remove-Item $scriptPath;
  } catch [Exception]{
    Write-Debug $_.Exception.Message;
  }
}

Register-EngineEvent PowerShell.Exiting -Action { cleanupDNV } -SupportEvent;