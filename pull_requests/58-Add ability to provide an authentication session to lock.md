# Add ability to provide an authentication session to lock

**yorinasub17** commented *Apr 26, 2022*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description
This PR updates the `lock` library to support using an already configured/authd AWS session. The use case for this is in the `refarch-deployer` utility, where all the authentication is handled internally based on where you are running, and more specifically, the local case where we use `aws-vault` as a library to auth to the different accounts, so the default credential chain is empty.


## TODOs

Please ensure all of these TODOs are completed before asking for a review.

- [x] Ensure the branch is named correctly with the issue number. e.g: `feature/new-vpc-endpoints-955` or `bug/missing-count-param-434`.
- [x] Update the docs.
- [x] Keep the changes backward compatible where possible.
- [x] Run the pre-commit checks successfully.
- [x] Run the relevant tests successfully.

<br />
***


**yorinasub17** commented *Apr 27, 2022*

Thanks for review! Merging now.
***

