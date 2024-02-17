def countdown(n):
    while n > 0:
        yield n     # <--- here's the `magic`
        n -= 1

for value in countdown(5):
    print(value)