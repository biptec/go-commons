# Use the same format for nested commands

**yorinasub17** commented *Nov 30, 2018*

If we have nested commands (e.g `kubergrunt eks configure`), then the help text for the second level is still using the default `urfave` template. This updates it to use the same one as the app level:

Original:

```
NAME:
   kubergrunt eks - Helper commands to configure EKS, including setting up operator machines to authenticate with EKS.

USAGE:
   kubergrunt eks command [command options] [arguments...]

COMMANDS:
     configure  Set up kubectl to be able to authenticate with EKS.
     token      Get token for Kubernetes using AWS IAM credential.

OPTIONS:
   --help, -h  show help

```

Updated:

```
Usage: kubergrunt eks [--help] command [options] [args]

Helper commands to configure EKS, including setting up operator machines to authenticate with EKS.

Commands:

   configure  Set up kubectl to be able to authenticate with EKS.
   token      Get token for Kubernetes using AWS IAM credential.
   help, h    Shows a list of commands or help for one command

```
<br />
***


