# Add functions for interacting with git

**yorinasub17** commented *Feb 15, 2022*

<!--
Have any questions? Check out the contributing docs at https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e,
or ask in this Pull Request and a Gruntwork core maintainer will be happy to help :)
Note: Remember to add '[WIP]' to the beginning of the title if this PR is still a work-in-progress. Remove it when it is ready for review!
-->

## Description

<!-- Write a brief description of the changes introduced by this PR -->
This implements useful functions for interacting with git in a CLI. These are adapted from [our `infrastructure-deploy-script` python script](https://github.com/gruntwork-io/terraform-aws-ci/blob/master/modules/infrastructure-deploy-script/infrastructure_deploy_script/git.py).

I need this for creating a docker entrypoint script for running steampipe checks in `ecs-deploy-runner`, which requires cloning the mods.

## TODOs

Please ensure all of these TODOs are completed before asking for a review.

- [x] Ensure the branch is named correctly with the issue number. e.g: `feature/new-vpc-endpoints-955` or `bug/missing-count-param-434`.
- [x] Update the docs.
- [x] Keep the changes backward compatible where possible.
- [x] Run the pre-commit checks successfully.
- [x] Run the relevant tests successfully.

<br />
***


**zackproser** commented *Feb 16, 2022*

Curious if you've ever played around with https://github.com/go-git/go-git?
***

**yorinasub17** commented *Feb 16, 2022*

> Curious if you've ever played around with https://github.com/go-git/go-git?

I have, and was originally going to use it, but I found it to be difficult to navigate around, especially when it comes to all the rules around checking out the "right" version from a ref (which is simplified to `git checkout` when using the CLI). E.g., see all the logic that `kaniko` implements to ensure you can clone the repo at the right version: https://github.com/GoogleContainerTools/kaniko/blob/main/pkg/buildcontext/git.go#L58-L139

I would have used this function directly, if it didn't hardcode the checkout dir...
***

**yorinasub17** commented *Feb 16, 2022*

Thanks for review! Merging now.
***

