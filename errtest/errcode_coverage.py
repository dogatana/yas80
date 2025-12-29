import re
import glob
import os.path

def main():
    count = {}
    code_names = get_codenames()
    for name in code_names:
        count[name] = 0
    
    for dir in ["errtest", "evaluator", "parser"]:
        read_testfiles(dir, count)

    uncovered = [k for k, v in count.items() if v == 0]
    print(f"coverage {len(count) - len(uncovered)}/{len(count)}")
    for name in uncovered:
        print(name, end=" ")

def get_codenames()->list[str]:
    path = os.path.join(os.path.dirname(__file__), "..", "errcode", "errcode.go")
    names = []
    with open(path, encoding="utf-8") as fp:
        for line in fp:
            line = re.sub(r"//.*$", "", line.strip())
            if line == "":
                continue
            words = line.split("=")
            if len(words) != 2:
                continue
            names.append(words[0].strip())
    return names

def read_testfiles(dir: str, count: map):
    path = os.path.join(os.path.dirname(__file__), "..", dir, "*_test.go")
    for file in glob.glob(path):
        names = read_testfile(file)
        for name in names:
            count[name] += 1

def read_testfile(file):
    names = []
    with open(file, encoding="utf-8") as fp:
        for line in fp:
            line = re.sub(r"//.*$", "", line.strip())
            m = re.search(r"errcode\.([a-zA-Z0-9_]+)", line)
            if m is None:
                continue
            names.append(m.group(1))
    return names

if __name__ == "__main__":
    main()