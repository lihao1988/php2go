# php2go

[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

#### 介绍
PHP 转行到 Golang 过程中，实现部分 PHP 代码中使用频率较高的函数，让 PHPer 蜕变为 Gopher 更简单、更快捷！

#### 安装
```shell
go get github.com/lihao1988/php2go
```
或
```shell
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
UrlEncode()            [PHP:urlencode(),编码 URL 字符串]
UrlDecode()            [PHP:urldecode(),解码已编码的 URL 字符串]
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

## LICENSE
PHP2Go source code is licensed under the [MIT](https://github.com/lihao1988/php2go/blob/main/LICENSE) Licence.
