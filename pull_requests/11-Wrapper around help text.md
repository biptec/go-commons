# Wrapper around help text

**yorinasub17** commented *Oct 7, 2018*

This adds the wrapped help text printer to `gruntwork-cli` and a function to create a new `cli.App` that takes advantage of it.

## Sample output:

### Main help
Original
```
NAME:
   houston - A CLI tool for interacting with Gruntwork Houston that you can use to authenticate to AWS on the CLI and to SSH to your EC2 Instances.

USAGE:
   houston [global options] command [command options] [arguments...]

VERSION:
   snapshot

COMMANDS:
     exec       Execute a command with temporary AWS credentials obtained by logging into Gruntwork Houston
     ssh        Connect to an EC2 instance via SSH with your public key.
     configure  Configure houston CLI options.
     profiles   Show a status of all the defined profiles.
     start      Start the background HTTP server process so it can cache your AWS credentials. Note: the exec command starts the HTTP server process automatically.
     stop       Stop the background HTTP server process.
     status     Show the status of the background HTTP server process.
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

New:
```
Usage: houston [--help] [--version] command [options] [args]

A CLI tool for interacting with Gruntwork Houston that you can use to
authenticate to AWS on the CLI and to SSH to your EC2 Instances.

Commands:

   exec       Execute a command with temporary AWS credentials obtained by
              logging into Gruntwork Houston
   ssh        Connect to an EC2 instance via SSH with your public key.
   configure  Configure houston CLI options.
   profiles   Show a status of all the defined profiles.
   start      Start the background HTTP server process so it can cache your
              AWS credentials. Note: the exec command starts the HTTP server
              process automatically.
   stop       Stop the background HTTP server process.
   status     Show the status of the background HTTP server process.
   help, h    Shows a list of commands or help for one command
```

### Command help
Original:
```
NAME:
   houston exec - Execute a command with temporary AWS credentials obtained by logging into Gruntwork Houston

USAGE:
   houston exec [OPTIONS] <profile> -- <command>

DESCRIPTION:
   The exec command makes it easier to use CLI tools that need AWS credentials, such as aws, terraform, and packer. Here's how it works:

   1. The first time you run this command for a <profile>, it will open your web browser and have you login to your Identity Provider (i.e., Google, ADFS, Okta).
   2. After login, the Identity Provider will redirect you to the Gruntwork Houston web console, where you will pick the AWS IAM Role you want to use.
   3. Gruntwork Houston will fetch temporary AWS credentials for this IAM Role and POST them back to houston CLI running on your computer.
   4. The houston CLI will set those credentials as the appropriate environment variables and execute <command>.
   5. The houston CLI will cache those credentials in memory, so all subsequent commands will execute without going through the login flow (until those credentials expire).

EXAMPLES:

   houston exec dev -- aws s3 ls
   houston exec prod -- terraform apply
   houston exec stage -- packer build server.json

OPTIONS:
   --config value, -c value  The configuration file for houston (default: "/Users/yoriy/.houston/houston.yml")
   --port value              The TCP port the http server is running on (default: 41170)
```

New:
```
Usage: houston exec [options] <profile> -- <command>

The exec command makes it easier to use CLI tools that need AWS credentials,
such as aws, terraform, and packer. Here's how it works:

  1. The first time you run this command for a <profile>, it will open your web
  browser and have you login to your Identity Provider (i.e., Google, ADFS,
  Okta).
  2. After login, the Identity Provider will redirect you to the Gruntwork
  Houston web console, where you will pick the AWS IAM Role you want to use.
  3. Gruntwork Houston will fetch temporary AWS credentials for this IAM Role
  and POST them back to houston CLI running on your computer.
  4. The houston CLI will set those credentials as the appropriate environment
  variables and execute <command>.
  5. The houston CLI will cache those credentials in memory, so all subsequent
  commands will execute without going through the login flow (until those
  credentials expire).

Examples:

   houston exec dev -- aws s3 ls
   houston exec prod -- terraform apply
   houston exec stage -- packer build server.json

Options:

   --config value, -c value  The configuration file for houston (default:
                             "/Users/yoriy/.houston/houston.yml")
   --port value              The TCP port the http server is running on (default:
                             41170)
```
<br />
***


**yorinasub17** commented *Oct 7, 2018*

UPDATE:

- Comments describing difference of these templates vs the default ones
- Add a text that came out of `urfave/cli` into unit test
- No more return type declared variables
- `TabAwareWrapText` => `IndentAwareWrapText` (the latter is a more accurate name)
- Comment with example of what `IndentAwareWrapText` is doing
- Make linewidth a non const so that it can be overridden.
- `SplitKeepDelimiter` => `RegexpSplitAfter`

TODO:

- [x] Unit test for the actual printer (the main func) checking for
    - [x] Using the new templates as opposed to the default
    - [x] Overriding linewidth
    - [x] Real world example
***

**yorinasub17** commented *Oct 8, 2018*

Ok added some unit tests with real world examples for the overall printers. @brikis98 Ready for final review before merge!
***

**yorinasub17** commented *Oct 8, 2018*

Ok merging and doing a release! Thanks!
***

