# Add locking functionality for resources in terragrunt while running tests

**ina-stoyanova** commented *Apr 23, 2021*

**Context:**
`SecurityHub` failing [test](https://app.circleci.com/pipelines/github/gruntwork-io/terraform-aws-cis-service-catalog/798/workflows/26d063c6-9fec-4ba4-ac74-2a3456a5fe2f/jobs/2324/artifacts)
`Guardduty` failing [test](https://circle-production-customer-artifacts.s3.amazonaws.com/picard/5795681ac9e77c00014a6759/607f1a86770ea97ce564c397-0-build/artifacts/tmp/logs/TestGuardDuty.log?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20210423T084112Z&X-Amz-SignedHeaders=host&X-Amz-Expires=60&X-Amz-Credential=AKIAJR3Q6CR467H7Z55A%2F20210423%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=20bb6d15bb3c077b37b7e9ffaef206bda562efe256088a5a0e3572dc0e62127d)
The occurrence of the failures is not very often, but it happens (as far as noticed) when more than 1 test suite runs at the same time (e.g. from multiple branches).

Suggested approach:
Add **locking** to resources as built-in functionality for `go-commons`. We expect that it will be needed going forward, and will help us use the same accounts to test features which can only be enabled once for the whole AWS Account - such as `SecurityHub`, `GuardDuty`.

**Example code for locking functionality:**
https://github.com/gruntwork-io/terragrunt/tree/v0.6.2/locks
https://github.com/gruntwork-io/prototypes/blob/master/release-notes-drafter/aws.go
https://aws.amazon.com/blogs/database/building-distributed-locks-with-the-dynamodb-lock-client/

**Testing**:
- created a new DynamoDB table in Sandbox account - `test-dynamo-lock-eu-jam`
- added a `playground` file to run tests from 
<br />
***


**ina-stoyanova** commented *Apr 23, 2021*

Previously, it was thought that this functionality could live on `terratest` but that was ruled out, as the potential to lock AWS resource could be more useful than just for testing. The [issue](https://github.com/gruntwork-io/terratest/issues/874) on `terratest` is now closed and this is the latest issue to discuss & work on is this one. 
***

**ina-stoyanova** commented *Apr 30, 2021*

Notes from TicketJam:

The issue that happens also in terraform-aws-security/guardduty-multi-region.

There is now a issue in Terratest's repo: [https://github.com/gruntwork-io/terratest/issues/874](https://github.com/gruntwork-io/terratest/issues/874)

Started by using the lock code from Terragrunt as base: [https://github.com/gruntwork-io/terragrunt/tree/v0.6.2/locks](https://github.com/gruntwork-io/terragrunt/tree/v0.6.2/locks)

Workflow:

1. One test (1) tries to enable SecurityHub, it creates an entry on the table of locks. If the table doesnt exist, it is created.
2. Another test (2) tries to enable SecurityHub at the same time. It should see that there is a lock in place and it doesn't fail, but instead it waits.
3. The test (1) runs sucessfully and releases the lock.
4. The test (2) finishes waiting and see that the lock is released, it created another lock and starts the test.

`withLock(createSecurityHub(), NUM_OF_RETRIES, WAIT_TIME, tableName)` {

→ create DynamoDB table if it doesn't exist

→ `**GuardDutyAvocado` already exists → do nothing**

→ `**GuardDutyAvocado` does not exist → create**

→ check for entry in the table if there's a lock - e.g. in `**GuardDutyAvocadoLock**`

→if there is no lock, create in  `**GuardDutyAvocado**` `**GuardDutyAvocadoLock**`

→ else, wait

...

}

Each functionality from each account will have their own table of locks

- guardduty table for PHX-OPS
- SecurityHub for LZ account

Account LandingZone will have a table called **GuardDutyTerratest**, for example.

All the repositories that run a GuardDuty test in whatever AWS account will lock in to the **GuardDutyTerratest** table. 

**We need a pattern on the names of the tables across repos.**
If we create a table `**GuardDutyAvocado**` in the CIS repo, and a **`GuardDutyStrawberries`** in Security repo, and run both tests at the same time **in the same account,** having lock will have no effect.
***

**ina-stoyanova** commented *Apr 30, 2021*

Things to check/do later:
- [x] do we need "SIGINT" / "signalChannel" 
- [x] look at how this code is used for signal checking to guard against ctrl+C
- [x] do we need to create the DynamoDB table used for locking the resources. For now we're assuming it already exists in the interest of simplicity.
- [x] what about testing: 
   - [x] how do we predictably test for this scenario? maybe create some code that enables Sec Hub & try to run it from multiple places?
- [x] which logging library to use? should we pass it in as an argument to the AcquireLock function?
   - [x] Interface - use form go-commons: example -> https://github.com/gruntwork-io/fetch/pull/90/files
   - [x] ideally we want to use the shared go libraries across all our repos, which would mean Terratest could just use go-commons out of the box, instead of passing in the logger it already has. Not sure if this is a bigger task in itself?
- [x] add specific errors instead of cli error types
- [x] we want to add the region in which a resource is created and locked in (add that to the LockID value)
- [x] make it configurable, specifically inputs from the function caller:
   - [x] ProjectAwsRegion
   - [x] ProjectLockTableName
   - [x] ProjectLockRetryTimeout
   - [x] LockTable
   - [x] LockString
   - [x] Timeout
***

**ina-stoyanova** commented *May 7, 2021*

Scenarios we want to cover:
- [x] Security Hub tests when running in parallel 
- [x] Release Notes drafter -> replace the locking mechanism with this new shared mechanism. 

***

