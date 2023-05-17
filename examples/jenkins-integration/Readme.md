### Summary
The Jenkinsfile presents a simple CI/CD scenario where it provisions a VDB, runs an automated test (like Selenium or JUnit) against that application & dataset, and then destroys the VDB. On testing failure, which will happen every time, a bookmark is created. This Jenksinfile leverages the DCT Terrafrom Provider and an API Curl command to show the full breadth of possibilites. All other steps are mocked out.

### Simple Getting Stated 
1) Create a Jenkinsfile Pipeline Job
2) Insert or reference the associated `Jenkinsfile` file.
    - Note: This Jenkinsfile also references the Terraform files in the `../simple-provision` folder. Feel free to fork, update, and modify those.
3) Update the following values:
    - DCT_HOSTNAME - Example: `123.0.0.0`
    - DCT_API_KEY - Example: `2.abc...`
        - [Manage this value through the Jenkins' Credentials plugin](https://docs.cloudbees.com/docs/cloudbees-ci/latest/cloud-secure-guide/injecting-secrets)
        - In a pinch, update it directly.
    - SOURCE_VDB - Example: `Oracle_QA`
4) Run Jenkins Job

Note: I suggest you reduce the sleep timers in early testing scenarios .


### Known Issues
On VDB destroy, the underlying Bookmark's snapshot will be deleted and the Bookmark will become "dangling".
Instead, I recommend using an "Enable/Disable" command instead of "Provision/Destroy" or skip the destroy VDB on failure.