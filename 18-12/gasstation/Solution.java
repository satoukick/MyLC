class Solution1 {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int stationNum = gas.length;
        int res = -1;

        for (int i = 0 ; i < stationNum; i++) { // start point
            int gasSum = gas[i];
            int j = i + 1;
            if (j >= stationNum) {
                j = 0;
            }
            int cnt = 0; //
            // int cosSum = 0;
            while (cnt <= stationNum) {
                cnt++;
                int prev = j - 1;
                if (prev < 0) {
                    prev = stationNum - 1;
                }
                // if (i == 4) {
                //     System.out.println(cnt + " " + gasSum + " " + cost[prev]);
                // }
                gasSum = gasSum - cost[prev];
               
                if (gasSum < 0) {
                    break;
                }

                if (cnt == stationNum) {
                    if (gasSum >= 0)
                        res = i;
                    break;
                }
                    
                else {
                    gasSum += gas[j];
                    if (++j >= stationNum) {
                        j = 0;
                    }
                }
            }
            if (res != -1)
                break;
        }
        return res;
    }
}

// [5,1,2,3,4]
// [4,4,1,5,1]

class test1 {
    public static void main(String[] args){
        Solution1 solu = new Solution1();
        int[] gas = {5,1,2,3,4};
        int[] cost = {4,4,1,5,1};
        int res = solu.canCompleteCircuit(gas, cost);
        System.out.println(res);
    }
}

