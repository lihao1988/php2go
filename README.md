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

## 版本要求
Go 1.15 or above.

## PHP 函数
### Array(Slice/Map) Functions
```php
## import array

In()                       [PHP:in_array(),检查数组中是否存在指定的值]
Keys()                 [PHP:array_keys(),返回数组中所有的键名]
Values()             [PHP:array_values(),返回数组中所有的值]
Merge()              [PHP:array_merge(),把一个或多个数组合并为一个数组]
Unique()           [PHP:array_unique(),删除数组中的重复值]
Column()          [PHP:array_column(),返回输入数组中某个单一列的值]
Diff()                   [PHP:array_diff(),比较数组,返回差集(只比较键值)]
Intersect()        [PHP:array_intersect(),比较数组，返回交集(只比较键值)]
```

## LICENSE
PHP2Go source code is licensed under the [MIT](https://github.com/lihao1988/php2go/blob/main/LICENSE) Licence.
