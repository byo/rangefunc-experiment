def countdown(n):
    while n > 0:
        yield n
        n -= 1

transformed = (
    x * x
    for x in countdown(10)
    if x % 2 == 0
)

print(type(transformed))

for value in transformed:
    print(value)