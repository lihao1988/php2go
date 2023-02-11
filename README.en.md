# php2go

[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

#### Introduce
In the process of PHP to Golang，realize functions that are used frequently in some PHP code, making it easier and faster for PHPer to transform into Gopher!

#### README
[README](https://github.com/lihao1988/php2go/blob/main/README.md)  
[README-en](https://github.com/lihao1988/php2go/blob/main/README.en.md)

#### Install
```shell
// github
go get github.com/lihao1988/php2go

// gitee
go get gitee.com/lihao1988/php2go
```

## Requirements
Go 1.15 or above.

## PHP Function
### Array(Slice/Map) Functions
```php
## import array

In()                   [PHP:in_array(),Checks if a value exists in an array]
Keys()                 [PHP:array_keys(),Return all the keys or a subset of the keys of an array]
Values()               [PHP:array_values(),Return all the values of an array]
Merge()                [PHP:array_merge(),Merge one or more arrays]
Unique()               [PHP:array_unique(),Removes duplicate values from an array]
Column()               [PHP:array_column(),Return the values from a single column in the input array]
Diff()                 [PHP:array_diff(),Computes the difference of arrays]
Intersect()            [PHP:array_intersect(),Computes the intersection of arrays]
```
### Crypto Functions
```php
## import crypto

MD5()                  [PHP:md5(),Calculate the md5 hash of a string]
Sha1()                 [PHP:sha1(),Calculate the sha1 hash of a string]
Base64Encode()         [PHP:base64_encode(),Encodes data with MIME base64]
Base64Decode()         [PHP:base64_decode(),Decodes data encoded with MIME base64]
```
### Date/Time Functions
```php
## import datetime

Time()                 [PHP:time(),Return current Unix timestamp]
StrToTime()            [PHP:strtotime(),Parse about any English textual datetime description into a Unix timestamp]
Date()                 [PHP:date(),Format a Unix timestamp]
CheckDate()            [PHP:checkdate(),Validate a Gregorian date]
Sleep()                [PHP:sleep(),Delay execution]
Usleep()               [PHP:usleep(),Delay execution in microseconds]
```
### Filesystem Functions
```php
## import file

Basename()             [PHP:basename(),Returns trailing name component of path]
Dirname()              [PHP:dirname(),Returns a parent directory's path]
Filesize()             [PHP:filesize(),Gets file size]
PathInfo()             [PHP:pathinfo(),Returns information about a file path]
FileExists()           [PHP:file_exists(),Checks whether a file or directory exists]
IsDir()                [PHP:is_dir(),Tells whether the filename is a directory]
IsFile()               [PHP:is_file(),Tells whether the filename is a regular file]
FileGetContents()      [PHP:file_get_contents(),Reads entire file into a string]
FilePutContents()      [PHP:file_put_contents(),Write data to a file]
Chmod()                [PHP:chmod(),Changes file mode]
Chown()                [PHP:chown(),Changes file owner]
```
### Math Functions
```php
## import math

Abs()                  [PHP:abs(),Absolute value]
Round()                [PHP:round(),Rounds a float]
Floor()                [PHP:floor(),Round fractions down]
Ceil()                 [PHP:ceil(),Round fractions up]
Max()                  [PHP:max(),Find highest value]
Min()                  [PHP:min(),Find lowest value]
DecBin()               [PHP:decbin(),Decimal to binary]
DecHex()               [PHP:dechex(),Decimal to hexadecimal]
```
### String Functions
```php
## import string

StrLen()               [PHP:strlen(),Get string length]
MbStrLen()             [PHP:mb_strlen(),Get string length by UTF-8]
SubstrCount()          [PHP:substr_count(),Count the number of substring occurrences]
Substr()               [PHP:substr(),Return part of a string]
MbSubstr()             [PHP:mb_substr(),Get part of string by UTF-8]
StrPos()               [PHP:strpos(),Find the position of the first occurrence of a substring in a string]
StrRPos()              [PHP:strrpos(),Find the position of the last occurrence of a substring in a string]
StrSplit()             [PHP:str_split(),Convert a string to an array]
UCFirst()              [PHP:ucfirst(),Make a string's first character uppercase]
```
### URL Functions
```php
## import url

ParseUrl()             [PHP:parse_url(),Parse a URL and return its components]
UrlEncode()            [PHP:urlencode(),URL-encodes string]
UrlDecode()            [PHP:urldecode(),Decodes URL-encoded string]
RawUrlEncode()         [PHP:rawurlencode(),URL-encode according to RFC 3986]
RawUrlDecode()         [PHP:rawurldecode(),Decode URL-encoded strings]
HttpBuildQuery()       [PHP:http_build_query(),Generate URL-encoded query string]
```

## LICENSE
PHP2Go source code is licensed under the [MIT](https://github.com/lihao1988/php2go/blob/main/LICENSE) Licence.
