bool increasingTriplet(int* nums, int numsSize){
    if (numsSize < 3) return false;
    
    int dp[3];
    memset(dp, 0x7F, sizeof dp);
    
    for (int j = 0; j < numsSize; ++j) {
        int num = nums[j];
        for (int i = 0; i < 3; ++i) {
            if (dp[i] < num) continue;
            if (i == 2) return true;
            dp[i] = num;
            break;
        }
    }
    
    return false;
}