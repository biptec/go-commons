# Must have struct for json to work, not nil

**yorinasub17** commented *Jul 14, 2022*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description

<!-- Write a brief description of the changes introduced by this PR -->
Fixes a bug in the new `github` package to properly handle the json unmarshal. It turns out using a nil uninitialized pointer is not supported (leads to the following runtime error):

```
Error loading configured Github App Config JSON: json: Unmarshal(nil *github.GithubAppConfig)
```

## TODOs

Please ensure all of these TODOs are completed before asking for a review.

- [x] Ensure the branch is named correctly with the issue number. e.g: `feature/new-vpc-endpoints-955` or `bug/missing-count-param-434`.
- [x] Update the docs.
- [x] Keep the changes backward compatible where possible.
- [x] Run the pre-commit checks successfully.
- [x] Run the relevant tests successfully.
- [x] Ensure any 3rd party code adheres with our [license policy](https://www.notion.so/gruntwork/Gruntwork-licenses-and-open-source-usage-policy-f7dece1f780341c7b69c1763f22b1378) or delete this line if its not applicable.
<br />
***


**yorinasub17** commented *Jul 14, 2022*

Thanks for review! Merging now.
***

