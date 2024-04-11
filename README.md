# dnv

Yet another cli tool to load directory specific environemnt variables.

## Installation

### Step 1. Install dnv executable

Select your operating system from the list below to view installation instructions:

<details>
<summary>Windows</summary>

Install `dnv` using the latest executable from the [releases section](https://github.com/sebakri/dnv/releases/latest)

Install `dnv` using any of the following package managers:

| Repository | Instructions                                                                   |
| ---------- | -------------------------------------------------------------------------------|
| [scoop]    | `scoop bucket add sebakri https://github.com/sebakri/scoop`                    |
|            | `scoop install dnv`                                                            |

</details>

### Step 2. Set up your shell to use Starship

Configure your shell to initialize `dnv`. Select yours from the list below:

<details>
<summary>PowerShell</summary>

Add the following to the end of your PowerShell configuration (find it by running `$PROFILE`):

```powershell
Invoke-Expression (& { (dnv init -debug pwsh | Out-String) })
```

</details>
