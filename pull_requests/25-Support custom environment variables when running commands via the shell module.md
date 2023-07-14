# Support custom environment variables when running commands via the shell module

**yorinasub17** commented *Aug 12, 2019*

Currently the `shell` module does not support custom environment variables when running commands. This PR adds a new `Env` option to the `ShellOptions` struct which will be converted to environment variables for the `Cmd` struct.
<br />
***


**yorinasub17** commented *Aug 13, 2019*

Merging and releasing! Thanks for review.
***

