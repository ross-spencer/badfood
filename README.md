# Filegen

Created for Artefactual Systems Inc. Allows for the generation of files of different sizes for testing Virus scanners. 

* v0.0.1-beta: Simple output, relies entirely on custom virus definition

## Credit

Initial code from @richardlehane: https://gist.github.com/ross-spencer/6776b33de01c2f5afaf8

### TODO

* Pipe in custom strings

## Clamscan and Clamdscan

### Custom Signatures in Clamscan:

[Creating definitions example](http://blog.adamsweet.org/?p=250)

e.g. 

     Trojan.Linux.SPENCERAF:0:*:afbadf00d0

Save using .ndb extension. Can be loaded from '/var/lib/clamav' where it will be loaded following the other definitions.

Relevant settings in Clamscan:

     --max-filesize=#n
          Extract and scan at most #n bytes from each archive. You may pass the value in kilobytes in format xK or xk, or megabytes in format
          xM or xm, where x is a number. This option protects your system against DoS attacks (default: 25 MB, max: <4 GB)

     --max-scansize=#n
          Extract  and  scan  at  most #n bytes from each archive. The size the archive plus the sum of the sizes of all files within archive
          count toward the scan size. For example, a 1M uncompressed archive containing a single 1M inner file counts as 2M toward  max-scan‐
          size.  You  may  pass  the value in kilobytes in format xK or xk, or megabytes in format xM or xm, where x is a number. This option
          protects your system against DoS attacks (default: 100 MB, max: <4 GB)

#### Config:

     /etc/clamav/freshclam.conf

#### Virus Definitions:

     /var/lib/clamav

### Custom definitions and Clamdscan

Clamdscan doesn't support clamscan -d mode:

     $ clamscan -d {customsignature}.ndb {filename}

Instead it has to be placed in the virus definitions folder above. 

     $ service clamav-daemon stop
     $ sudo freshclam
     $ service clamav-daemon start

or possibly, just make sure the file is there and:

     $ service clamav-daemon restart

It will take a few good minutes. The clamdscan restart might also take some time. While it is 
restarting, or if it isn't starting at all you might see the following:

     ERROR: Could not lookup : Servname not supported for ai_socktype

If it all works, you'll see the following:

     github.com/ross-spencer/filegen/badfood: Trojan.Linux.SPENCERAF.UNOFFICIAL FOUND

##### 👍
