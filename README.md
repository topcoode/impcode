# impcode

https://medium.com/@ansujain/mastering-middleware-in-go-tips-tricks-and-real-world-use-cases-79215e72b4a8 
-----------------> middleware in golang using gin framework
https://gobyexample.com/command-line
https://medium.com/@kdnotes/command-line-flags-with-go-lang-
-----------------> command line flags



update the go version 
Step 1: Remove the existing golang
sudo apt-get purge golang*



Step 2: Download the latest version from the official site. Click Here


Step 3: Extract it in /usr/local using the following command. I am using Go 1.11.4 here. You may need to replace the filename with the actual filename based on the version you have downloaded.
tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz


Step 4: Create .go directory in home. (It is easy to install the necessary packages without admin privilege)
mkdir ~/.go


Step 5: Set up the following environment variables

GOROOT=/usr/local/go
GOPATH=~/.go
PATH=$PATH:$GOROOT/bin:$GOPATH/bin
Check this link on how to set environment variables permanently.

Step 6: Update the go command

sudo update-alternatives --install "/usr/bin/go" "go" "/usr/local/go/bin/go" 0
sudo update-alternatives --set go /usr/local/go/bin/go
Step 7: Test the golang version

go version
