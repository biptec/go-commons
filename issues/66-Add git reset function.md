# Add git reset function

**rhoboat** commented *Jul 5, 2022*

Probably makes the most sense to put it in `go-commons`. I think the use case for reseting in testing is much more limited than for go CLIs, since we don't want to accidentally undo all the changes while you are iterating in testing (e.g., running `git reset` on the current repo instead of the copy we are using for upgrade testing).

_Originally posted by @yorinasub17 in https://github.com/gruntwork-io/terraform-aws-ci/pull/466#discussion_r914044640_
<br />
***


