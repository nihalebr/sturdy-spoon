/**
 *
 * A function which returns the length of longest substring.
 * Which takes a String as input.
 *
 * @param  {string} inputString - Input String
 * @returns Length of the Longest Substring
 */
function lengthOfLongestSubstring(inputString: string): number {
  if (inputString.length === 0) {
    return 0;
  }
  let subString = new Set<string>();
  let start = 0,
    end = 0,
    maxLength = 0;
  while (start < inputString.length && end < inputString.length) {
    if (!subString.has(inputString[end])) {
      subString.add(inputString[end]);
      end++;
      maxLength = Math.max(maxLength, end - start);
    } else {
      subString.delete(inputString[start]);
      start++;
    }
  }
  return maxLength;
}

console.log(lengthOfLongestSubstring("abcabcbb"));
console.log(lengthOfLongestSubstring("bbbbb"));
console.log(lengthOfLongestSubstring("pwwkew"));
console.log(lengthOfLongestSubstring("javaconceptoftheday"));
