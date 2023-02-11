# php2go

[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

#### 介绍
PHP 转行到 Golang 过程中，实现部分 PHP 代码中使用频率较高的函数，让 PHPer 蜕变为 Gopher 更简单、更快捷！

#### README
[README](https://github.com/lihao1988/php2go/blob/main/README.md)  
[README-en](https://github.com/lihao1988/php2go/blob/main/README.en.md)

#### 安装
```shell
// github
go get github.com/lihao1988/php2go

// gitee
go get gitee.com/lihao1988/php2go
```

## 版本要求
Go 1.15 or above.

## PHP 函数
### Array(Slice/Map) Functions
```php
## import array

In()                   [PHP:in_array(),检查数组中是否存在指定的值]
Keys()                 [PHP:array_keys(),返回数组中所有的键名]
Values()               [PHP:array_values(),返回数组中所有的值]
Merge()                [PHP:array_merge(),把一个或多个数组合并为一个数组]
Unique()               [PHP:array_unique(),删除数组中的重复值]
Column()               [PHP:array_column(),返回输入数组中某个单一列的值]
Diff()                 [PHP:array_diff(),比较数组,返回差集(只比较键值)]
Intersect()            [PHP:array_intersect(),比较数组，返回交集(只比较键值)]
```
### Crypto Functions
```php
## import crypto

MD5()                  [PHP:md5(),计算字符串的 MD5 散列]
Sha1()                 [PHP:sha1(),计算字符串的 SHA-1 散列]
Base64Encode()         [PHP:base64_encode(),使用 base64 对数据进行编码]
Base64Decode()         [PHP:base64_decode(),对使用 base64 编码的数据进行解码]
```
### Date/Time Functions
```php
## import datetime

Time()                 [PHP:time(),返回当前时间的 Unix 时间戳]
StrToTime()            [PHP:strtotime(),将任何英文文本的日期或时间描述解析为 Unix 时间戳]
Date()                 [PHP:date(),格式化本地日期和时间]
CheckDate()            [PHP:checkdate(),验证日期格式的合法性]
Sleep()                [PHP:sleep(),延迟代码执行若干秒]
Usleep()               [PHP:usleep(),延迟代码执行若干微秒]
```
### Filesystem Functions
```php
## import file

Basename()             [PHP:basename(),返回路径中的文件名部分]
Dirname()              [PHP:dirname(),返回路径中的目录名称部分]
Filesize()             [PHP:filesize(),返回文件大小]
PathInfo()             [PHP:pathinfo(),返回关于文件路径的信息]
FileExists()           [PHP:file_exists(),检查文件或目录是否存在]
IsDir()                [PHP:is_dir(),判断文件是否是一个目录]
IsFile()               [PHP:is_file(),判断文件是否是常规的文件]
FileGetContents()      [PHP:file_get_contents(),把文件读入字符串]
FilePutContents()      [PHP:file_put_contents(),把字符串写入文件]
Chmod()                [PHP:chmod(),改变文件模式]
Chown()                [PHP:chown(),改变文件所有者]
```
### Math Functions
```php
## import math

Abs()                  [PHP:abs(),绝对值]
Round()                [PHP:round(),对浮点数进行四舍五入]
Floor()                [PHP:floor(),向下舍入为最接近的整数]
Ceil()                 [PHP:ceil(),向上舍入为最接近的整数]
Max()                  [PHP:max(),返回最大值]
Min()                  [PHP:min(),返回最小值]
DecBin()               [PHP:decbin(),把十进制转换为二进制]
DecHex()               [PHP:dechex(),把十进制转换为十六进制]
```
### String Functions
```php
## import string

StrLen()               [PHP:strlen(),获取字符串长度]
MbStrLen()             [PHP:mb_strlen(),按照 UTF-8 获取字符串的长度]
SubstrCount()          [PHP:substr_count(),计算字串出现的次数]
Substr()               [PHP:substr(),返回字符串的子串]
MbSubstr()             [PHP:mb_substr(),获取部分字符串]
StrPos()               [PHP:strpos(),查找字符串首次出现的位置]
StrRPos()              [PHP:strrpos(),计算指定字符串在目标字符串中最后一次出现的位置]
StrSplit()             [PHP:str_split(),将字符串转换为数组]
UCFirst()              [PHP:ucfirst(),将字符串的首字母转换为大写]
```
### URL Functions
```php
## import url

ParseUrl()             [PHP:parse_url(),解析 URL，返回其组成部分]
UrlEncode()            [PHP:urlencode(),编码 URL 字符串]
UrlDecode()            [PHP:urldecode(),解码已编码的 URL 字符串]
RawUrlEncode()         [PHP:rawurlencode(),按照 RFC 3986 对 URL 进行编码]
RawUrlDecode()         [PHP:rawurldecode(),对已编码的 URL 字符串进行解码]
HttpBuildQuery()       [PHP:http_build_query(),生成 URL-encode 之后的请求字符串]
```

## LICENSE
PHP2Go source code is licensed under the [MIT](https://github.com/lihao1988/php2go/blob/main/LICENSE) Licence.
