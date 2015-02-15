## Integration tests

In the integration folder is an ansible playbook that can be used to run integration tests:

```bash
ansible-playbook -i integration/local integration/test.yml --extra-vars "kay_package=$(PACK_NAME)"
```
# Assumptions

- This has been tested with ansible 1.7.
- The play books assume that the kay_package variable is a zip archive that contains an executable named kay and a bash completion file named kay.bash.
- It installs into an install dir, and a bash dir, and runs the tests against a test dir (see [Install](roles/install/tasks/main.yml) for variable names).
- If you use the test playbook it uses temporary directories for install and runs several integration tests.  It will clean up the temp directories on successful completion of the integration tests.


