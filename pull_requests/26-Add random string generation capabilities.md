# Add random string generation capabilities

**yorinasub17** commented *Aug 12, 2019*

In the GCP bootstrap script I am currently developing, I have a need to generate a random string that I can use as a unique identifier for seeding the Project ID when creating a new project. I haven't seen us provide this as a library outside of `terratest`, so I ported a version of `random.UniqueId` here that is more generic, where you can control char sets and digits.

Note that this uses `crypto/rand` instead of `math/rand`, to allow for potential usage with password generation.
<br />
***


**yorinasub17** commented *Aug 13, 2019*

Ok addressed the question about being more clear about the usage. Will merge now and release. Thanks for the review!
***

