class Solution {
public:
    std::vector<int> corpFlightBookings(std::vector<std::vector<int>>& bookings, int n) {
        auto opt = std::vector<int>(n);
        for (; !bookings.empty(); bookings.pop_back()) {
            auto v = bookings.back();
            opt[v[0] - 1] += v[2];
            if (v[1] < n) opt[v[1]] -= v[2];
        }
        std::partial_sum(opt.begin(), opt.end(), opt.begin());
        return opt;
    }
};