//
// This is only a SKELETON file for the 'Armstrong Numbers' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const isArmstrongNumber = function () {
  const num = [...arguments][0];
  const numOfDigits = num.toString().length;
  const sum = num
    .toString()
    .split("")
    .reduce((acc, curr) => {
      const digit = parseInt(curr, 10);
      return (digit ** numOfDigits) + acc;
    }, 0);
  return sum === num;
};
