# Add function to convert key=value slice pairs to map

**yorinasub17** commented *Jul 26, 2022*

This adds a new function that helps with converting a key=value slice of strings to a map. This is useful for working with the return value of `os.Environ()`.

## TODOs

Read the [Gruntwork contribution guidelines](https://gruntwork.notion.site/Gruntwork-Coding-Methodology-02fdcd6e4b004e818553684760bf691e).

- [x] Update the docs.
- [x] Run the relevant tests successfully, including pre-commit checks.
- [x] Ensure any 3rd party code adheres with our [license policy](https://www.notion.so/gruntwork/Gruntwork-licenses-and-open-source-usage-policy-f7dece1f780341c7b69c1763f22b1378) or delete this line if its not applicable.
- [x] Include release notes. If this PR is backward incompatible, include a migration guide.

## Release Notes (draft)

- Added a new function `collections.KeyValueStringSliceAsMap` which can be used to work with a slice of `key=value` pairs.
<br />
***


**yorinasub17** commented *Jul 27, 2022*

Thanks for review! Merging in now.
***

