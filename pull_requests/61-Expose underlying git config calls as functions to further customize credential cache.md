# Expose underlying git config calls as functions to further customize credential cache

**yorinasub17** commented *Apr 29, 2022*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description

This PR adds two new functions `git.ConfigureCacheCredentialsHelper` and `git.StoreCacheCredentials`. These functions give the CLI more control over the credential caching behavior of git, allowing you to have mixed credentials for accessing different repos.

This is most useful when you are using GitHub Apps for repo access, which can only auth to a single org at a time. In such situation, you need the ability to provide different credentials for different orgs.

<!-- Write a brief description of the changes introduced by this PR -->

## TODOs

Please ensure all of these TODOs are completed before asking for a review.

- [x] Ensure the branch is named correctly with the issue number. e.g: `feature/new-vpc-endpoints-955` or `bug/missing-count-param-434`.
- [x] Update the docs.
- [x] Keep the changes backward compatible where possible.
- [x] Run the pre-commit checks successfully.
- [x] Run the relevant tests successfully.

<br />
***


**yorinasub17** commented *Apr 29, 2022*

Thanks for review! Merging now.
***

