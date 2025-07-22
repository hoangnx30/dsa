/**
 * @param {number[]} nums
 * @return {number}
 */
// var maximumUniqueSubarray = function (nums) {
//   let result = 0;
//   const map = {};
//   let left = 0;
//   const prefixSum = Array(10000 + 1).fill(0);

//   for (let right = 0; right < nums.length; right++) {
//     const num = nums[right];
//     prefixSum[right + 1] = prefixSum[right] + num;

//     if (map[num] !== undefined) {
//       left = Math.max(map[num] + 1, left);
//     }

//     const sum = prefixSum[right + 1] - prefixSum[left];

//     result = Math.max(result, sum);

//     map[num] = right;
//   }

//   return result;
// };

// /**
//  * @param {number[]} nums
//  * @return {number}
//  */
// var maximumUniqueSubarray = function (nums) {
//   let result = 0;
//   const map = {};
//   let left = 0;
//   const prefixSum = Array(10000 + 1).fill(0);

//   for (let right = 0; right < nums.length; right++) {
//     const num = nums[right];
//     prefixSum[right + 1] = prefixSum[right] + num;

//     if (map[num] !== undefined) {
//       left = Math.max(map[num] + 1, left);
//     }

//     const sum = prefixSum[right + 1] - prefixSum[left];

//     result = Math.max(result, sum);

//     map[num] = right;
//   }

//   return result;
// };

var maximumUniqueSubarray = function (nums) {
  let nmap = new Int8Array(10001),
    total = 0,
    best = 0;
  for (let left = 0, right = 0; right < nums.length; right++) {
    nmap[nums[right]]++;
    total += nums[right];
    console.log(nmap);

    while (nmap[nums[right]] > 1) {
      nmap[nums[left]]--;
      total -= nums[left++];
    }
    best = Math.max(best, total);
  }
  return best;
};

console.log(maximumUniqueSubarray([5, 2, 1, 2, 5, 2, 1, 2, 5])); // Output: 17
