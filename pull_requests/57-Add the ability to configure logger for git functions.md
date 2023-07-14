# Add the ability to configure logger for git functions

**yorinasub17** commented *Apr 19, 2022*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description

<!-- Write a brief description of the changes introduced by this PR -->

Sometimes, for security reasons, we may want to suppress the log output of the git commands running. This is not possible with the current functions, as it uses the default logger associated with the `ShellOptions`.

This PR updates the functions in the `git` module to accept the logger object explicitly, and configure it if non-nil.

## Backward compatibility

This PR is backward **incompatible**.

## Migration guide

The functions in the `git` module has been updated to require a `logrus.Logger` object to be passed in to use as the logger for the git commands. You can pass in `nil` for the new arg to retain the old behavior.

## TODOs

Please ensure all of these TODOs are completed before asking for a review.

- [x] Ensure the branch is named correctly with the issue number. e.g: `feature/new-vpc-endpoints-955` or `bug/missing-count-param-434`.
- [x] Update the docs.
- [x] Keep the changes backward compatible where possible.
- [x] Run the pre-commit checks successfully.
- [x] Run the relevant tests successfully.

<br />
***


**yorinasub17** commented *Apr 21, 2022*

Thanks for review!

> Are we aware of all the places we call ConfigureForceHTTPS, Checkout and Clone? Those will need to be updated downstream.

Not really, but it doesn't really matter because go will complain when you bump the version to the new backward incompatible version and will be easy to fix. So as long as we highlight it in the release notes, I think we are ok in terms of taking care of the backward incompatibility.
***

