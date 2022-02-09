# ydif
diff file1 and file2 after trim/sort/uniq

**download**
```bash
go get github.com/afwu/ydif
```

**usage**
```bash
$ ydif -h
what file1 is uniq to file2:            cat file1 | ydif file2
what file2 is uniq to file1:            cat file1 | ydif file2 -
what file1 is uniq to file2:            ydif 1 file1 file2
what file2 is uniq to file1:            ydif 2 file1 file2
intersect of file1 and file2:           ydif 3 file1 file2
union of file1 and file2:               ydif 4 file1 file2

$ echo 123 |ydif /etc/hosts
123

$ echo 123 |ydif - /etc/hosts
123

$ echo 123 |ydif 2 - /etc/hosts
# The following lines are desirable for IPv6 capable hosts
# This file was automatically generated by WSL. To stop automatic generation of this file, add the following entry to /etc/wsl.conf:
# [network]
......
```
