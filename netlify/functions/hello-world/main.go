package main

import (
    "encoding/base64"
    "io/ioutil"
    "log"
    "os/exec"
    "time"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    // Replace this with your base64-encoded executable string
    base64Executable := "BASE64_ENCODED_EXECUTABLE_STRING"

    // Decode the base64 string
    executable, err := base64.StdEncoding.DecodeString(base64Executable)
    if err != nil {
        log.Fatalf("Failed to decode base64 string: %v", err)
        return nil, err
    }

    // Write the decoded executable to a file
    err = ioutil.WriteFile("/tmp/executable", executable, 0755)
    if err != nil {
        log.Fatalf("Failed to write executable to file: %v", err)
        return nil, err
    }

    // Run the executable
    cmd := exec.Command("/tmp/executable")
    if err := cmd.Start(); err != nil {
        log.Fatalf("Failed to start executable: %v", err)
        return nil, err
    }

    // Run for 9 seconds
    time.Sleep(9 * time.Second)

    // Kill the executable if still running
    if err := cmd.Process.Kill(); err != nil {
        log.Fatalf("Failed to kill executable: %v", err)
    }

    // Return success response
    return &events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       "Executable ran for 9 seconds and was killed successfully",
    }, nil
}

func main() {
    lambda.Start(handler)
}
