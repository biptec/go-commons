# Improve ErrorWithExitCode

**yorinasub17** commented *Jan 31, 2020*

I noticed that the user experience for returning `ErrorWithExitCode` isn't quite great, so I made a few improvements here:

- Support using both `WithStackTrace` and `ErrorWithExitCode`, so that you can get both a stack trace and exit with exit code.
- Output the underlying error message instead of `ErrorWithExitCode{ Err = Deployment failed, ExitCode = 1 }`
- Refactor `checkForErrorsAndExit` so we can test the exit code functionality without the test exiting due to the call to `os.Exit`.
<br />
***


**yorinasub17** commented *Feb 6, 2020*

I'm blocked on this in my other PR so going to go ahead and merge+release this, but feedback is still welcome.
***

