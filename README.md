beacon
=========
Filters and Parses IRC on #HDBits and downloads the announcements based on the filtering requirements for Video extraction. *Audio Track not supported*

Requirements
------------
* Go 1.4.x or greater

## Installation
    # Use go install to install checksum package
    go install github.com/souleiman/beacon

Upon installation you will need to make some configuration to begin using beacon. The following set assumes you have your Go environment setup properly specifically your gopaths, if not, here https://golang.org/doc/code.html. 

Inside the sample/, you will find 3 files and you will only need to configure two of these files. You will need to choose on whether you want beacon to connect to ZNC or act as an IRC Bot. Based on what you choose, copy the file over to your home directory and once completed, you will need to modify them. Both files should follow JSON formatting.

    > cp $GOPATH/src/github.com/souleiman/beacon/sample/.beacon.rc.sample.xxx ~/.beaconrc

Open up your favorite editor and make the modifications suitable for your needs. Below is a table explaining each necessary parameters. 

## Basic Beacon Configuration
For both znc and irc configurations, you must fill out these fields.

Key | Type | Comment |
-----|:---------:|:------|
username|string| The username used on hdbits
passkey|string| Can be obtained under profile page.
stalk|[]string| List of users to watch out for, this is most likely midgards.
*channels|[]string| Only specify the list of channels that you want beacon to keep an eye on. Make sure to include #
output|string| Specify the directory in which you want to save the torrent file. This is most likely your watch folder.

### IRC Configuration
If you opted to choose this configuration, you will need to make modification to the irc values.

Key | Type | Comment |
-----|:---------:|:------|
host|string| The ip address of the server you want to connect to. i.e. irc.freenode.net
port|int| The port in which you want to connect to the server.
nick|string| The name in which you want beacon to be referred to when connected to the host.

### ZNC Integration
If you have ZNC setup and are on the recent version (I don't recall which version, but if your bouncer supports connecting by specifying the PASS then you should be good.) you may choose to set up beacon to connect on a preconfigured account.

Key | Type | Comment |
-----|:---------:|:------|
host|string| The address of your bouncer.
port|int| The port in which you want to connect to the host's bouncer.
password|string| You must specify the password to connect to on of your users configured on your bouncer to connect to the bouncer. For example, Username/NetworkName:Password
ssl|bool| If you have ssl configured on your bouncer, you will need to set this to true, otherwise false.
insecure_ssl_verify_skip|bool| If you receive an error x509 or you have an unknown certificate, set this to true, otherwise false.

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
