x = 1
cycleVals = []

for line in open(0):
    if line == "noop\n":
        # noop takes one cycle to complete. It has no other effect.
        cycleVals.append(x)
    else: 
        # addx V takes two cycles to complete
        # after two cycles, the X register is increased by the value V. (V can be negative.)
        val = int(line.split()[1])
        cycleVals.append(x)
        cycleVals.append(x)
        x += val


print(sum(x * y + y for x, y in list(enumerate(cycleVals))[19::40]))
