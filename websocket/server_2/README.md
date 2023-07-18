## Testing the app
1.  Launch the app with
    ```bash
        go run server.go
    ```
2.  Open the app from browser, and send the message through the text box.
    ```bash
        http://localhost:8080
    ```
3.  This will trigger the the web-socket in the background which is hosted on `localhost:8080/echo`