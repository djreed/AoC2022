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

# For this task we just want each line as it comes
dataList = list(map(eval, open(0).read().split()))

# The first packet is at index 1, the second packet is at index 2, and so on.
sep1 = [[2]]
sep1Index = 1 # Index of [[2]] in a list of the separators

sep2 = [[6]]
sep2Index = 2 # Index of [[6]] in a list of the separators

# Don't need to sort the list itself, just need to find where the two
# Separators will exist, so we can just figure out on which side the 
for packet in dataList:
    if compare(packet, sep1) < 0:
        sep1Index += 1
        sep2Index += 1
    elif compare(packet, sep2) < 0:
        sep2Index += 1

print(sep1Index * sep2Index)