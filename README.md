dossh [![Gobuild Download](http://gobuild.io/badge/github.com/9uuso/dossh/download.png)](http://gobuild.io/github.com/9uuso/dossh)
=======

dossh is a tool for which creates `ssh` hostname rules from your DigitalOcean servers for easier connection. For example, if you have a droplet in DigitalOcean with hostname `foo.com`, you will be able to connect to it with command `ssh foo.com`. All you need is an API key with read access from `https://cloud.digitalocean.com/settings/applications`. If you don't want to memorize the API key every time you can also create a `.profile` (or `.bashrc` depending on your platform) function as so:

    function dossh { PATH_TO_DOSSH YOUR_API_KEY; }
    export -f dossh

For example, mine looks like this

    function dossh { ~/Github/dossh/dossh e1aa1...; }
    export -f dossh

After opening a new command line window you should then be able to call `dossh` globally. Each time you call the program it will refresh your SSH hostname rules to match your current servers. You can then connect to your droplet with its hostname. For example, we can now login to our `foo.com` server with simply `ssh foo.com`.

##Installation

Here you have two choices. The first one is to download a ready-made binary from [Gobuild](http://gobuild.io/download/github.com/9uuso/dossh). Alternatively you can build the source by yourself.

To build the source you need to `git clone https://github.com/9uuso/dossh` and then `go build`. For this to work you have to have [Go](http://golang.org/doc/install) installed on your system.

##What operating systems does it support?

MacOSX, at least. And all other operating systems which have ssh config file located at `~/.ssh` should work as well.

##Why you made this?

I personally manage a bunch of DO servers and all of them are located behind a proxy server. With this awesome tool of mine, I'm able to connect to the servers with ease.

##License

MIT