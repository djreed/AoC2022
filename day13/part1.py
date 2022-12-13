# x < y  => negative
# x == 0 => 0
# x > y  => positive
def compare(x, y):
    if type(x) == int:
        if type(y) == int:
            return x - y
        else:
            return compare([x], y)
    else: # List
        if type(y) == int:
            return compare(x, [y])
    
    for x1, y1 in zip(x, y):
        v = compare(x1, y1)
        if v:
            return v
    
    return len(x) - len(y)

# Split data into blocks of two packets each
dataList = list(map(str.splitlines, open(0).read().strip().split("\n\n")))

total = 0

for i, (a, b) in enumerate(inList):
    if compare(eval(a), eval(b)) < 0:
        total += i + 1

print(total)