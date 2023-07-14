# Feature: DynamoDB Lock

**ina-stoyanova** commented *May 14, 2021*

Adding a DynamoDB lock feature based on the code from [prototypes/release-notes-drafter](https://github.com/gruntwork-io/prototypes/blob/65c81b6/release-notes-drafter/aws.go).

Related: https://github.com/gruntwork-io/go-commons/issues/44
<br />
***


**ina-stoyanova** commented *Jun 4, 2021*

I've gone ahead and addressed all of the above feedback. Another set of eyes will be very much appreciated :) 
***

**ina-stoyanova** commented *Jun 7, 2021*

Thanks for the review and approval @bwhaley :) 
***

**ina-stoyanova** commented *Jun 7, 2021*

Thanks, Jim! That's all really great feedback from you and Ben! I'll get to this today 🙂 
***

**ina-stoyanova** commented *Jun 7, 2021*

Next todo: 
- fix build
- add delete dynamoDB table so test can cleanup after itself
- review comments if they make sense
- run test & confirm intermittent errors seen are expected
- use functionality from cis reop branch & see if the test there still passes
- make sure error handling and logging makes sense and does not add extra code for no reason
***

**brikis98** commented *Jun 10, 2021*

Is this ready for re-review?
***

**ina-stoyanova** commented *Jun 10, 2021*

Just coming back to this now! I'll ping you again! I'm following up on this:

- [x] add delete dynamoDB table so test can cleanup after itself
- [x] review comments if they make sense
- [x] run test & confirm intermittent errors seen are expected
- [x] use functionality from cis reop branch & see if the test there still passes
- [x] make sure error handling and logging makes sense and does not add extra code for no reason
***

**marinalimeira** commented *Jun 11, 2021*

Updates on TicketJam:
- remove cleanup, it causes concurrency errors if a test is waiting to lock while another test finished running and deleted the table.
- update the test table to always use the same table, but different lock id strings.
***

**ina-stoyanova** commented *Jun 11, 2021*

> remove cleanup, it causes concurrency errors if a test is waiting to lock while another test finished running and deleted the table.

Ahhh! I didn't see that coming! I actually could see that there was many tables created, which didn't seem right, so wanted a way to clean them up! Anyway - thanks for catching this! 🙇 

***

**ina-stoyanova** commented *Jun 11, 2021*

I believe this is ready for hopefully a final set of eyes @brikis98 if you've got the chance :) 
***

**ina-stoyanova** commented *Jun 11, 2021*

So, I've updated a few more comments actually - hope they make sense now and are useful! Thanks for all the reviews here :) 
***

**ina-stoyanova** commented *Jun 11, 2021*

Thanks for the review, Jim! I'm merging and releasing this! 

***

