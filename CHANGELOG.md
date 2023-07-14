# Changelog


# v0.9.2
_Released Jun 11, 2021_
### Modules affected

-  `lock` [**NEW**]

### Description

- Add new functionality to allow locking of resources when they're being used from multiple places simultaneously.
We expect this to be useful especially when testing AWS resources/services which can only be enabled once for the whole AWS Account - such as SecurityHub, GuardDuty, or any others.

### Related links
https://github.com/gruntwork-io/go-commons/pull/46
https://github.com/gruntwork-io/go-commons/pull/47
<br>

# v0.9.1
_Released Apr 28, 2021_
### Modules affected

-  `awscommons` [**NEW**]

### Description

- Add new functions to interact with the AWS API using the AWS SDK for Go V2.

### Related links

https://github.com/gruntwork-io/go-commons/pull/43
<br>

# v0.9.0
_Released Apr 23, 2021_
### Modules affected

-  `version` [**BACKWARD INCOMPATIBLE**]

### Description

This release renames the `VERSION` variable in the `version` package to `Version` to align with conventional Go style. The `Version()` func has also been renamed to `GetVersion()`.

### Related links

https://github.com/gruntwork-io/go-commons/pull/45

<br>

# v0.8.2
_Released Mar 17, 2021_
### Modules affected

-  `retry` [**NEW**]

### Description

- Add new functions to allow for retrying a specific action.

### Related links
https://github.com/gruntwork-io/go-commons/pull/41
<br>

# v0.8.1
_Released Mar 12, 2021_
### Modules affected

-  `ssh` [**NEW**]

### Description

- Add new functions to allow running a command over SSH.

### Related links
https://github.com/gruntwork-io/go-commons/pull/40
<br>

# v0.8.0
_Released Mar 11, 2021_
**This release is backwards incompatible.**

## Description

This repo has been renamed to to `go-commons`. You will need to update your references from `gruntwork-cli` to `go-commons` to use the updated version.
<br>

# v0.7.2
_Released Jan 22, 2021_
### Modules affected

-  `shell`
### Description

- Add new functions to allow getting the `stdout` from a running command.

### Related links
https://github.com/gruntwork-io/gruntwork-cli/pull/38
<br>

# v0.7.1
_Released Dec 14, 2020_
### Modules affected

-  `version` **[NEW MODULE]**
- `entrypoint`


### Description

- Added a new `version` package that can be used to set the `VERSION` variable from the CLI, and return the version via `version.Version()`. This reduces the boilerplate needed for managing version numbers in CLI apps. See the `version` package and updated README for more details.
- Fixed some minor string formatting in `entrypoint`.

### Related links

https://github.com/gruntwork-io/gruntwork-cli/pull/37
<br>

# v0.7.0
_Released Aug 19, 2020_
### Modules affected

- `logging` [**BACKWARDS INCOMPATIBLE**]

### Description

`logging.GetLogger` now returns a new logger using a custom formatter that prefixes the given name to the log entries.

### Related links

https://github.com/gruntwork-io/gruntwork-cli/pull/32
<br>

# v0.6.1
_Released Feb 6, 2020_
### Modules affected

- `entrypoint`

### Description

- Support using both `WithStackTrace` and `ErrorWithExitCode`, so that you can get both a stack trace and exit with exit code.
- Output the underlying error message instead of `ErrorWithExitCode{ Err = Deployment failed, ExitCode = 1 }`

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/30
<br>

# v0.6.0
_Released Jan 29, 2020_
### Description

This release switches the dependency management from `dep` to go modules. This means that any go libraries and tools that are including this repo will need to switch to go modules from this release onward.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/28
<br>

# v0.5.1
_Released Aug 13, 2019_
### Modules affected

- `random` [**NEW**]

### Description

This release introduces a new module `random`, which exposes utility functions for generating random items.

- `random.RandomString` can be used for generating random strings of specified lengths, pooling from a set of allowed characters passed into the function. For example, to generate a 32 char string only composing of digits, you can do `random.RandomString(32, random.Digits)`. Additionally, you can do `random.RandomString(32, random.LowerLetters + random.SpecialChars)` to generate a 32 char string composing of lowercase letters (`a-z`) and special characters.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/26
<br>

# v0.5.0
_Released Aug 13, 2019_
### Modules affected

- `shell` [**BREAKING CHANGE**]

### Description

- You can now set additional environment variables in the execution context for running commands using the `shell` module. This is handled by setting the `Env` attribute on `ShellOptions`.

### Migration guide

If you were manually constructing `ShellOptions`, you will need to ensure you initialize `Env` to an empty map. For example, if you had:

```
options := shell.ShellOptions{}
```

in your code, you will need to update to initialize `Env`:

```
options := shell.ShellOptions{Env: map[string]string{}}
// OR
options.Env = map[string]string{}
```

Note that if you were using `shell.NewShellOptions` to initialize the `ShellOptions`, then you do not need to change anything.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/25
<br>

# v0.4.2
_Released Dec 19, 2018_
### Modules affected

- `collections`

### Description

- Adds `BatchListIntoGroupsOf` function, which will take a list and return a list of lists that batches the original list into groups of batchSize elements.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/20
<br>

# v0.4.1
_Released Nov 30, 2018_
### Modules affected

- `entrypoint`

### Description

- Help text for nested subcommands now follow the same pattern as the app level help text.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/19
<br>

# v0.4.0
_Released Nov 28, 2018_
### Modules affected

- `shell`
- `entrypoint`

### Description

- There are now new functions to assert required arguments passed into the CLI: `StringFlagRequiredE` and `EnvironmentVarRequiredE`. (This is in `entrypoint` package).
- There is now a `E` variant for `CommandInstalled` in the `entrypoint` package.
- (**Backwards incompatible**) Stack traces in error messages are now suppressed by default, printing only a simple error message. This leads to better user experience at the expense of potential debuggability. The old behavior can still be restored if the command is run with the environment variable `GRUNTWORK_DEBUG` set to something other than empty string.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/15
- https://github.com/gruntwork-io/gruntwork-cli/pull/16
- https://github.com/gruntwork-io/gruntwork-cli/pull/18
<br>

# v0.3.1
_Released Nov 28, 2018_
### Modules affected

- `shell`

### Description

- `RunShellCommand` and its siblings now support the concept of sensitive args to suppress logging the command string if the args contain sensitive data.

### Related links

- https://github.com/gruntwork-io/gruntwork-cli/pull/14
<br>

# v0.3.0
_Released Nov 28, 2018_
https://github.com/gruntwork-io/gruntwork-cli/pull/13: The `url.FormatUrl` method now supports URL fragments (e.g., `foo.com#bar`). This is an additional required parameter in the method, so this is a backwards incompatible change.
<br>

# v0.2.1
_Released Nov 27, 2018_
https://github.com/gruntwork-io/gruntwork-cli/pull/12: Add new `url.FormatUrl` method for building URLs that properly handle escaping, encoding, leading/trailing slashes, etc.
<br>

# v0.2.0
_Released Oct 8, 2018_
https://github.com/gruntwork-io/gruntwork-cli/pull/11: CLI help printer that can wrap text while preserving indentations in tables.

To use, you can manually override the templates and HelpPrinter functions of the cli, or use `entrypoint.NewApp()` to construct the cli app which will take care of applying the modifications. You can also modify the line width by changing `entrypoint.HelpTextLineWidth` (defaults to 80).
<br>

# v0.17.0
_Released Jun 15, 2023_
## What's Changed
* Feature/telemetry forwarder by @ellisonc in https://github.com/gruntwork-io/go-commons/pull/90


**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.16.2...v0.17.0

## Migration Guide
The application no longer needs to provide a mixpanel_token to the telemetry library
<br>

# v0.16.2
_Released Apr 20, 2023_
## What's Changed
* Fix input var type for BatchListIntoGroupsOf by @levkoburburas in https://github.com/gruntwork-io/go-commons/pull/80
* Fix telemetry error by @ellisonc in https://github.com/gruntwork-io/go-commons/pull/88


**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.16.1...v0.16.2
<br>

# v0.16.1
_Released Mar 30, 2023_
## What's Changed
* Bugfix/public event vars by @ellisonc in https://github.com/gruntwork-io/go-commons/pull/87


**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.16.0...v0.16.1
<br>

# v0.16.0
_Released Mar 29, 2023_
## What's Changed
* chore(CORE-640): Centralize mixpanel code by @MoonMoon1919 in https://github.com/gruntwork-io/go-commons/pull/85

## New Contributors
* @eak12913 made their first contribution in https://github.com/gruntwork-io/go-commons/pull/86
* @MoonMoon1919 made their first contribution in https://github.com/gruntwork-io/go-commons/pull/85

**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.15.0...v0.16.0
<br>

# v0.15.0
_Released Mar 10, 2023_
## What's Changed
* Feature/log formatting by @ellisonc in https://github.com/gruntwork-io/go-commons/pull/84

## New Contributors
* @ellisonc made their first contribution in https://github.com/gruntwork-io/go-commons/pull/84

**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.14.0...v0.15.0

## Migration Guide

Previously cli applications were created by calling `entrypoint.NewApp()`, then the application name and version were set after the fact for both the logging and the CLI description.
The new way to do this is a single call to `entrypoint.NewApp(name, version)`

Calling `logging.GetLogger(name, version)` will now return a logrus Entry instead of a logrus Logger. This entry has the name and version preset.

The behavior of the default logger is the same as in previous versions of go-commons.

A new JSON logger is available by calling `logging.SetGLobalLogFormatter("json")`
<br>

# v0.14.0
_Released Nov 11, 2022_
## Description
The minimum supported version of Go has been updated to `1.18`.

## Migration guide

`go-commons` no longer supports any go version under `1.18`. If you need to use `go-commons` going forward, you will need to update to at least go version `1.18` in your project.

## Related links

- https://github.com/gruntwork-io/go-commons/pull/79


**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.13.5...v0.14.0
<br>

# v0.13.5
_Released Oct 25, 2022_
## What's Changed
* Update codeowners by @zackproser in https://github.com/gruntwork-io/go-commons/pull/77
* Make GetPathRelativeTo() resolve symbolic links by @infraredgirl in https://github.com/gruntwork-io/go-commons/pull/78


**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.13.4...v0.13.5
<br>

# v0.13.4
_Released Oct 24, 2022_
## What's Changed
* Add package docs file by @yorinasub17 in https://github.com/gruntwork-io/go-commons/pull/69
* Add utilities for autoscaling. by @rhoboat in https://github.com/gruntwork-io/go-commons/pull/72
* Port a bug fix for GetPathRelativeTo over from Terragrunt by @infraredgirl in https://github.com/gruntwork-io/go-commons/pull/76

## New Contributors
* @rhoboat made their first contribution in https://github.com/gruntwork-io/go-commons/pull/72
* @infraredgirl made their first contribution in https://github.com/gruntwork-io/go-commons/pull/76

**Full Changelog**: https://github.com/gruntwork-io/go-commons/compare/v0.13.3...v0.13.4
<br>

# v0.13.3
_Released Jul 27, 2022_
## Description

- Added a new function `collections.KeyValueStringSliceAsMap` which can be used to work with a slice of `key=value` pairs.

## Related links

- https://github.com/gruntwork-io/go-commons/pull/71
<br>

# v0.13.2
_Released Jul 14, 2022_
## Description

- Fixed bug in `github` package where json unmarshalling was improperly implemented.

## Related links

- https://github.com/gruntwork-io/go-commons/pull/70
<br>

# v0.13.1
_Released Jul 14, 2022_
## Description

- Introduced a new package `github` that will contain helper routines for interacting with GitHub. For the first release, this package exposes a new high level function for getting an installation token for a GitHub App so you can use it to auth to the API.

## Related links

- https://github.com/gruntwork-io/go-commons/pull/68
<br>

# v0.13.0
_Released Jul 13, 2022_
## Description

- Updated underlying dependency versions to latest. As a part of this, the minimum supported version of go has been updated to `1.17`.

## Migration guide

`go-commons` no longer supports any go version under `1.17`. If you need to use go-commons going forward, you will need to update to at least go version `1.17` in your project.

## Related links

- https://github.com/gruntwork-io/go-commons/pull/67
<br>

# v0.12.5
_Released Jun 10, 2022_
### Modules affected
* `collections`

### Description 
* Added `KeyValueStringSlice` and `KeyValueStringSliceWithFormat` functions to convert a map into a string slice using the default or specified format.

### Related links
https://github.com/gruntwork-io/go-commons/pull/65
<br>

# v0.12.4
_Released May 4, 2022_
### Modules affected
* `url`

### Description 
* Added `OpenURL` function to provide a cross-platform means of opening an arbitrary URL in the user's browser from within a Golang program.

### Related links
https://github.com/gruntwork-io/go-commons/pull/64
<br>

# v0.12.3
_Released Apr 29, 2022_
### Modules affected 
* `git`

### Description 
* Added `ConfigureCacheCredentialsHelper` and `StoreCacheCredentials` functions to give CLI tools more control over the credential caching behavior of git, allowing you to have mixed credentials for accessing different repos.

### Related links
https://github.com/gruntwork-io/go-commons/pull/61
<br>

# v0.12.2
_Released Apr 28, 2022_
### Modules affected 
* `lock`

### Description 
* Added `ScanLocks` function to support functionality reporting on all currently held locks. 
<br>

# v0.12.1
_Released Apr 27, 2022_
### Modules affected

- `lock`

### Description

- Added the ability to provide an AWS Session directly to the `lock` functions to customize auth behavior of the SDK.

### Related links

https://github.com/gruntwork-io/go-commons/pull/58
<br>

# v0.12.0
_Released Apr 21, 2022_
### Modules affected

- `git` [**BACKWARD INCOMPATIBLE**]

### Description

- Updated all `git` functions to accept a `logrus.Logger` object. This allows you to customize the logging options of the function (including suppressing the log entries with `io.Discard`).

### Migration guide

The functions in the `git` module has been updated to require a `logrus.Logger` object to be passed in to use as the logger for the git commands. You can pass in `nil` for the new arg to retain the old behavior.

### Related links

https://github.com/gruntwork-io/go-commons/pull/57
<br>

# v0.11.0
_Released Feb 28, 2022_
### Modules affected

- `entrypoint` [**BACKWARD INCOMPATIBLE**]
- `errors`

### Description

- Updated dependency `urfave/cli` to v2 branch. All usage of `urfave/cli` needs to update to v2. Refer to the [urfave/cli v1 to v2 migration guide](https://github.com/urfave/cli/blob/master/docs/migrate-v1-to-v2.md) for information on how to update downstream projects to this version.

### Related links

https://github.com/gruntwork-io/go-commons/pull/33
<br>

# v0.10.2
_Released Feb 17, 2022_
### Modules affected

- `awscommons`

### Description

- Added new function for retrieving all enabled regions in an account.

### Related links

https://github.com/gruntwork-io/go-commons/pull/56
<br>

# v0.10.1
_Released Feb 16, 2022_
### Modules affected

- `git` [**NEW**]

### Description

- Added new functions for interacting with the `git` CLI.

### Related links

https://github.com/gruntwork-io/go-commons/pull/55
<br>

# v0.10.0
_Released Jul 16, 2021_
### Modules affected

- `shell` [**BACKWARD INCOMPATIBLE**]
- `awscommons`


### Description

- Add new function to upload string contents to S3 bucket.
- Refactor shell functions to use a common underlying runtime interface.
- New shell run command functions `RunShellCommandAndGetOutputStruct` and `RunShellCommandAndGetOutputStructAndStreamOutput` which will return a struct that captures stdout, stderr, and merged outputs so you can access all of those outputs.
- Refactor implementation of streaming and capturing outputs. This implementation respects the ordering in stdout and stderr better. The previous implementation always preferred to read stdout over stderr, which resulted in delaying the streaming of stderr if both were simultaneously written to. The updated implementation will properly interleave the contents regardless of timing.

**NOTE**: This release introduces changes that modify the behavior of the captured and streamed outputs for the `*StreamOutput` shell functions. Specifically, the streamed logs are now more interleaved, and the output strings will include terminating newlines if the original stdout/stderr included a terminating newline before EOF.


### Related links

https://github.com/gruntwork-io/go-commons/pull/51
https://github.com/gruntwork-io/go-commons/pull/52
<br>

# v0.1.3
_Released Oct 5, 2018_
https://github.com/gruntwork-io/gruntwork-cli/pull/10: Unit testable shell prompt utilities.
<br>

# v0.1.2
_Released Dec 17, 2017_
https://github.com/gruntwork-io/gruntwork-cli/pull/8: Add a password prompt method.
<br>

# v0.1.1
_Released Nov 27, 2017_
https://github.com/gruntwork-io/gruntwork-cli/pull/5: Add new `Keys` helper for getting the sorted keys from a `map`.
<br>

# v0.1.0
_Released Jun 26, 2017_
[logrus](https://github.com/sirupsen/logrus) changed their GitHub repo to all lowercase and broke all go packages that use their library.

See [here](https://github.com/sirupsen/logrus/issues/553) and [here](https://github.com/sirupsen/logrus)
<br>

# v0.0.5
_Released Jun 7, 2017_
https://github.com/gruntwork-io/gruntwork-cli/pull/3: Add a shell command that can both capture stdout/stderr and stream it to logs.
<br>

# v0.0.4
_Released Jun 7, 2017_
https://github.com/gruntwork-io/gruntwork-cli/commit/6d058a64c7fee3afea2b37cbe6ccbbe632371c54: Fix the logging and returned output for `shell` commands.
<br>

# v0.0.3
_Released May 26, 2017_
https://github.com/gruntwork-io/gruntwork-cli/pull/2: Improve logging and error handling.
<br>

# v0.0.2
_Released Jan 7, 2017_
- Override urfave/cli's `OsExiter` so `os.Exit` isn't called in random places.

<br>

# v0.0.1
_Released Dec 21, 2016_
First release!

<br>

