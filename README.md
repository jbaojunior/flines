# flines

flines is a little binary write in go to facilitate the extraction of specified range from a file. I always use the conjunte "head | tail" but sometimes is a little boring.

The options is very simples:
```
-f  - File to extract the information. If not specified will be read frm STDIN.
-s  - Start line to extract. If not specified will be start from line 1.
-e  - End line to extract. If not specified will be until the end of file.
-h  - A little help
```

Example:
``` 
echo -e "First Line\nSecond Line\nLast Line" | ./flines -s 2
```