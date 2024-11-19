class Solution {
private:
    static bool cmpA(const std::vector<int> &a, const std::vector<int> &b) {
        return a[0] > b[0];
    }
    static bool cmpB(const std::vector<int> &a, const std::vector<int> &b) {
        return a[1] > b[1];
    }
public:
    std::vector<int> corpFlightBookings(std::vector<std::vector<int>>& bookings, int n) {
        std::sort(bookings.begin(), bookings.end(), cmpA);
        
        auto sum = 0;
        auto priQue = std::priority_queue<std::vector<int>, std::vector<std::vector<int>>, decltype(&cmpB)>(cmpB);
        
        auto opt = std::vector<int>(n);
        for (int i = 0; i < n; ++i) {
            while (!priQue.empty() && priQue.top()[1] <= i) {
                sum -= priQue.top()[2];
                priQue.pop();
            }
            while (!bookings.empty() && bookings.back()[0] == i+1) {
                sum += bookings.back()[2];
                priQue.push(bookings.back());
                bookings.pop_back();
            }
            opt[i] = sum;
        }
        return opt;
    }
};