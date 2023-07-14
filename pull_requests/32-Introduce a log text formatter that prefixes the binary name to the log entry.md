# Introduce a log text formatter that prefixes the binary name to the log entry

**yorinasub17** commented *Aug 19, 2020*

Gives you logs that look like below:

```
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]   ami-builder
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]       build-packer-artifact
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]   docker-image-builder
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]       build-docker-image
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]   terraform-applier
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]       infrastructure-deploy-script
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]       terraform-update-variable
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]   terraform-planner
[infrastructure-deployer] INFO[2020-08-19T11:23:08-05:00]       infrastructure-deploy-script
```

<br />
***


**yorinasub17** commented *Aug 19, 2020*

Thanks for review! Going to merge this in now.
***

