monkeys = []

# Monkey 0:
#   Starting items: 56, 56, 92, 65, 71, 61, 79
#   Operation: new = old * 7
#   Test: divisible by 3
#     If true: throw to monkey 3
#     If false: throw to monkey 7
for group in open(0).read().strip().split("\n\n"):
    lines = group.splitlines()
    monkey = []
    monkey.append(list(map(int, lines[1].split(": ")[1].split(", "))))
    monkey.append(eval("lambda old:" + lines[2].split("=")[1]))
    for l in lines[3:]:
        monkey.append(int(l.split()[-1]))
    monkeys.append(monkey)

counts = [0] * len(monkeys)

for _ in range(20): # 20 Rounds
    for idx, monkey in enumerate(monkeys):
        testDivisible = monkey[2]
        trueIdx = monkey[3]
        falseIdx = monkey[4]

        for item in monkey[0]:
            newItem = monkey[1](item)
            newItem //= 3
            if newItem % testDivisible == 0:
                monkeys[trueIdx][0].append(newItem)
            else:
                monkeys[falseIdx][0].append(newItem)
        counts[idx] += len(monkey[0])
        monkey[0] = []

print(counts)
counts.sort()
print(counts[-1] * counts[-2])