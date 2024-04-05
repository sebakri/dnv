using namespace System;
using namespace System.Management.Automation;

$env:DNV_SHELL = "pwsh";
$env:DNV_DEBUG = "{{.Debug | ternary true false}}";
$env:DNV_SESSION_ID = [System.Diagnostics.Process]::GetCurrentProcess().Id;
$env:DNV_SESSION_FOLDER = Join-Path $([System.IO.Path]::GetTempPath()) "dnv" $env:DNV_SHELL-$env:DNV_SESSION_ID;

$hook = [EventHandler[LocationChangedEventArgs]] {
  param([object] $source, [LocationChangedEventArgs] $eventArgs)
  end {
    $DebugPreference = "{{.Debug | ternary "Continue" "SilentlyContinue"}}"
    $unloadScript = Join-Path $env:DNV_SESSION_FOLDER "unload.ps1";
    $loadScript = Join-Path $env:DNV_SESSION_FOLDER "load.ps1";

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
      Write-Host $_.Exception.Message;
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
