```bash
PS C:\Users\hanwei\Documents\GitHub\example\golang\pflag-demo> .\main.exe -h      
Usage of C:\Users\hanwei\Documents\GitHub\example\golang\pflag-demo\main.exe:
  -a, --age int[=25]        Input Your Age (default 22)
  -d, --des.detail string   Input Description
  -g, --gender string       Input Your Gender (default "male")
  -n, --name string         Input Your Name (default "nick")
  -o, --ok                  Input Are You OK
pflag: help requested
PS C:\Users\hanwei\Documents\GitHub\example\golang\pflag-demo> .\main.exe -n abc
name= abc
age= 22
gender= male
ok= false
des=
PS C:\Users\hanwei\Documents\GitHub\example\golang\pflag-demo>
```