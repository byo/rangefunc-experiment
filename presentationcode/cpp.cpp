#include <iostream>
#include <ranges>
 
int main()
{
    auto even = [](int i) { return 0 == i % 2; };
    auto square = [](int i) { return i * i; };

    // the "pipe" syntax of composing the views:
    for (auto i : std::views::iota(0, 10)
                | std::views::filter(even)
                | std::views::transform(square))
        std::cout << i << std::endl;
}
