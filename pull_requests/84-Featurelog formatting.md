# Feature/log formatting

**ellisonc** commented *Mar 9, 2023*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description

Updated common library with support for json logs and a better interface for logging the app binary and version

## Migration Guide

Previously cli applications were created by calling `entrypoint.NewApp()`, then the application name and version were set after the fact for both the logging and the CLI description.
The new way to do this is a single call to `entrypoint.NewApp(name, version)`

Calling `logging.GetLogger(name, version)` will now return a logrus `Entry` instead of a logrus `Logger`.  This entry has the name and version preset.

The behavior of the default logger is the same as in previous versions of go-commons.

A new JSON logger is available by calling `logging.SetGLobalLogFormatter("json")`

<br />
***


