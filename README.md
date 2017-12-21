# badfood

Created for Artefactual Systems Inc. Allows for the generation of files of different sizes for testing Virus scanners. 

* v0.0.1-beta: Simple output, relies entirely on custom virus definition
* v0.0.2-beta: Incorporates Nukem virus definition
* v0.0.2-beta: Incorporates EICAR

The tool will create a file of -size (n)MB. It will place a pattern that will match against a ClamAV definition part-way (half-way) through the file. So for example, in the description below, we will inject the following sequence:

      afbadf00d0

And providing our custom definition is available to Clamscan and Clamdscan then we will find a match. This allows us to test Clam's ability to spot the virus across a wide range of files of varying size. 

### ClamAV use

You have to run a command extending the filesize and/or scansize if the file is bigger than 25MB

     clamscan --max-scansize=32M --max-filesize=32M {filename}

clamdscan you need to set the config elsewhere:

     #/etc/clamav/clamd.conf settings...
     MaxScanSize 200M
     MaxFileSize 200M

### Acknowledgements

Initial code from @richardlehane: https://gist.github.com/ross-spencer/6776b33de01c2f5afaf8

#### TODO

* Pipe-in custom strings

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

#### Decompiling the signature database to see byte sequences

     sigtool --unpack=/var/lib/clamav/main.cvd

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

## Eicar

Any anti-virus product that supports the EICAR test file should detect it in any 
file providing that the file starts with the following 68 characters, and is 
exactly 68 bytes long:

     X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*

The first 68 characters is the known string. It may be optionally appended by 
any combination of whitespace characters with the total file length not 
exceeding 128 characters. The only whitespace characters allowed are the space 
character, tab, LF, CR, CTRL-Z. To keep things simple the file uses only upper 
case letters, digits and punctuation marks, and does not include spaces. The 
only thing to watch out for when typing in the test file is that the third 
character is the capital letter "O", not the digit zero.

EICAR will always be at the beginning of the file so in a lot of cases you're not
double checking a virus scanner's ability to seek through the file and correctly
identify a sequence there. 

