# Setup Golang Project
1.  Create a simple working directory with
    |-  greet/
    |-  main/
2.  Initialize the dependecy tracking only in the parent directory 
    ```bash
        go mod init golang-learning
    ```
3.  One can synchronize the modules' dependencies in every path directory with
    ```bash
        go mod tidy
    ```

# Unit-Testing
1.  Since the unit tests are created for intended package, one needs to run the test command from the project parent diectory as such
    ```bash
        go test -v golang-learning/<PACKAGE>
    ```
    If the test command is run from the package directory 
    ```bash
        cd <PACKAGE>/
        go test -v golang-learning/<PACKAGE>
    ```
