# Akamai EdgeGrid Tester

**Note**: This utility is for development and troubleshooting the HMAC authentication logic utilized by Akamai REST APIs. It does not make any web requests.

### macOS
If the utility will not run with the following prompt. (will show as `killed` in terminal)

>Apple could not verify “akamai-edgegrid-tester_darwin_arm64” is free of malware that may harm your Mac or compromise your privacy.

Once the error is received, open the Settings app, then navigate to the Security & Privacy section. There should be a prompt to allow the utility in the Security section.

## Build instructions
1. Install dependencies by running `go get .`
2. Compile the application by running `go build -o akamai-edgegrid-tester` or run the application directly by running `go run main.go`