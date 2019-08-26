int singleNumber(int* nums, int numsSize){
    int a = 0, b = 0;
    for (int i = 0; i < numsSize; ++i) {
        a = (a^nums[i]) & ~b;
        b = (~a & nums[i]) ^ b;
    }
    return a;
}