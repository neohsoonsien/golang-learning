# Setup Golang Project
1.  Create a simple working directory with
    |-  greet/
    |-  main/
2.  Initialize the dependecy tracking for each directory with
    ```bash
        cd greet/
        go mod init learning/greet
    ```
    ```bash
        cd main/
        go mod init learning/main
    ```
3.  One can synchronize the modules' dependencies in every path directory with
    ```bash
        go mod tidy
    ```
