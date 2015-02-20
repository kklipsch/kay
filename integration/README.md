## Integration tests

In the integration folder are ansible playbooks for integration tests and installation. 

# Assumptions

- This has been tested with ansible 1.7.
- The play books assume that the kay_package variable is a zip archive that contains an executable named kay and a bash completion file named kay.bash.
- It installs into an install dir, and a bash dir, and runs the tests against a test dir (see [Install](roles/install/tasks/main.yml) for variable names).
- If you use the test playbook it uses temporary directories for install and runs several integration tests.  It will clean up the temp directories on successful completion of the integration tests.

```bash
ansible-playbook -i integration/local integration/test.yml --extra-vars "kay_package=$(PACK_NAME)"
```

# Installation

```bash
ansible-playbook -i integration/local integration/install.yml --extra-vars '{"kay_package": "$(PACK_NAME)", "install_directory": {"stdout": "$(INSTALL_DIR)"}, "bash_sources": {"stdout": "$(BASH_SOURCES)"}}'
```

This installs the app in the provided app diretory and the bash autocomplete in the directory labelled BASH_SOURCES. 
