package BuyandSellStock;

class Solution {
    public int maxProfit(int[] prices) {
        if (prices.length == 0)
            return 0;
        int max = 0, in = prices[0];
        for (int i = 1 ; i < prices.length; i++){
            if (prices[i] - in > 0){
                if (prices[i] - in > max)
                    max = prices[i] - in;
            }
            else
            {
                in = prices[i];
            }
        }
        return max;
    }
}