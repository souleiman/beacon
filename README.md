beacon
=========
Filters and Parses IRC on #HDBits and downloads the announcements based on the filtering requirements for Video extraction. *Audio Track not supported*

Requirements
------------
* Go 1.4.x or greater

## Installation
    # Use go install to install beacon package
    go install github.com/souleiman/beacon

Upon installation you will need to make some configuration to begin using beacon. The following set assumes you have your Go environment setup properly specifically your gopaths, if not, here https://golang.org/doc/code.html. 

Inside the sample/, you will find 2 files and you will only need to configure. You will need to choose on whether you want beacon to connect to ZNC or act as an IRC Bot. Based on what you choose, you will need to configure it to support your needs. Both files should follow JSON formatting.

    > cp $GOPATH/src/github.com/souleiman/beacon/sample/.beacon.rc.sample ~/.beaconrc

Open up your favorite editor and make the modifications suitable for your needs. Below is a table explaining each necessary parameters. 

## Basic Beacon Configuration
For both znc and irc configurations, you must fill out these fields.

Key | Type | Comment |
-----|:---------:|:------|
username|string| The username used on hdbits
passkey|string| Can be obtained under profile page.
output|string| Specify the directory in which you want to save the torrent file. This is most likely your watch folder.

### IRC Configuration
If you opted to choose this configuration, you will need to make modification to the irc values.

Key | Type | Comment |
-----|:---------:|:------|
host|string| The ip address of the server you want to connect to. i.e. irc.freenode.net
port|int| The port in which you want to connect to the server.
nick|string| The name in which you want beacon to be referred to when connected to the host.
password|string| You must specify the password to connect to the server if needed. For znc users, follow the guide for your version, most recent version can be as followed "Nick/Server Name:Password"
ssl|bool| If you want to connect to the server with ssl
insecure_ssl_verify_skip|bool| If you receive an error x509 or you have an unknown certificate, set this to true, otherwise false.
stalk|[]string| List of users to watch out for, this is most likely midgards.
*channels|[]string| Only specify the list of channels that you want beacon to keep an eye on. Make sure to include #
commands|[][]string| For each commands, will consists of a command that you will send to a user. So, for example, first parameters will contain the command, "msg" for example. Followed by the target who to "msg" and followed by the message you want to send. This is useful for cases when you want to join a channel that requires invitation.


### ZNC Integration
If you have ZNC setup and are on the recent version (I don't recall which version, but if your bouncer supports connecting by specifying the PASS then you should be good.) you may choose to set up beacon to connect on a preconfigured account.



\* Will connect you to the channel if not joined.

## Filtering
Now the last step is to finalize the filtering options.

    > cp $GOPATH/src/github.com/souleiman/beacon/sample/.filter.beacon.sample ~/.filter.beacon

Similar to the previous configuration, however for these settings you are only removing what you do not want. It should be noted that, the filters behave very specific.

    always_dl_internal — if set to true and any torrent that's considered internal will automatically be snatched.
    maximum_seeder — An integer value which indicates the maximum number seeders in which to ignore the snatch. Setting this value to anything below 1 inclusively will snatch anything.
    
For the rest, you may modify the array to refine your snatches, keep in mind and make sure that there's at least one option in each parameter exists or it will always fail.

Usage
-------

Running beacon is very simple. Just call it and it should be connected to IRC, if everything is setup properly.

    > beacon
