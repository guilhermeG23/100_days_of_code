### Scopes in SHELL

By default, every variable in SHELL has a global scope, that is, at the moment the variable is created, any part of the script can change it, based on this, there is a way to restrict this action using ``` LOCAL``` in functions.

LOCAL restricts the variable to the scope of the function, even if an existing identification is used, it is not changed in the global scope, follow the example in the shell_script directory.
