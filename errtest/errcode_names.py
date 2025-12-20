import sys
import re

def main(infile:str, outfile:str):
    names = read_names(infile)
    write_go(outfile, names)

def read_names(file:str)->list[str]:
    names = []
    with open(file, encoding="utf-8") as fp:
        for line in fp:
            line = re.sub(r"//.*$", "", line.strip())
            if line == "":
                continue
            words = line.split("=")
            if len(words) != 2:
                continue
            names.append(words[0].strip())
    return names

def write_go(file:str, names:list[str]):
    with open(file, "w", encoding="utf-8") as fp:
        print("""package errtest

import "yas80/errcode"

var errcodeNames map[string]string = map[string]string {
""", file=fp)
        for name in names:
            print(f'\terrcode.{name}: "{name}",', file=fp)
        print("}", file=fp)

if __name__ == "__main__":
    if len(sys.argv) != 3:
        exit(1)
    main(*sys.argv[1:])




