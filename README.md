# Description
Simpurl is a tool that reads a list of URLs with multiple parameters from standard input and prints them out with each parameter.

### Usage

```
cat urls.txt | simpurl 
```

### Installation
```
go install https://github.com/channyein1337/simpurl@latest
```


### Example
Before using simpurl
```
https://test.com/index.php?user=admin&password=123456&id=1&email=admin@test.com
```

After using simpurl

```
https://test.com/index.php?user=admin
https://test.com/index.php?password=123456
https://test.com/index.php?id=1
https://test.com/index.php?email=admin@test.com
```
