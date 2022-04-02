<link rel="stylesheet" href="styles.css">

# The Mortuary QR Code Scanner that takes in JSON tags
This scanner takes in json strings and if the user is logged in with a name registered on the scanner, the application sends a post request to the server. This is how the server receives the information and keeps the information of casualty updated. As different scanners scan the same QR Code (even at the same time), the application itself handles all its concurrency using Go's mutexes found in the sync package


<h2>Introduction</h2>

<p>This project can easily be repurposed for something else. You will find types in the entities folders and your SSL certs should be placed into the certs. They're omitted for obvious reasons. Required libraries and JS dependencies have already been included as a local dependency. Feel free to upload them to your edge CDN servers if you like.</p>

<h2>Development Environment</h2> 
<p>For development purposes in your own environment, use the following command once you've in the working directory</p>

```go run .```

<p>You can access the server on <i>localhost</i> or <i>localhost:80</i> if the former doesn't work</p>

<h2>Building binary</h2>
<p>Since the server is written in go, you can build binaries using the go build command. The following command will build an binary name called mortuaryApp tailored for the unbuntu environment</p>
<p>GOOS=linux GOARCH=amd64 go build -o MortuaryApp .</p>

<p>Remember to use ```chmod +x MortuaryApp``` if the binary cannot be exected</p>

<h2>Dependencies</h2>
<p>While workspace files have to be included as part of the deployment, external dependencies don't have to be referenced from elsewhere. All the vendor files have already been included as part of the project. If not found or accidentally deleted, please run the ```go mod vendor``` command</p>

<h2>Executing for Production</h2>
<p>After your files have been deployed to an actual web server such as AWS EC2's web server, you can run the binary using the following command</p>
<p>```go run . --productionFlg --domain www.exampledomain.com```</p>

<p>When production flag is activated, requests routed on port:80 will be automatically routed to port:443</p>
<p>If you open up chrome's dev tools, you'll see HTTP/1.1 and then see a HTTP/2 almost immedidately</p>
