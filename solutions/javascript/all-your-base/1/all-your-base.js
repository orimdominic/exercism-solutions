//
// This is only a SKELETON file for the 'All Your Base' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const convert = function () {
  const [digitsArr, inBase, outBase] = [...arguments];
  if (inBase <= 1) {
    throw new Error("Wrong input base");
  }
  if (outBase <= 1) {
    throw new Error("Wrong output base");
  }
  if (digitsArr.length === 0 || (digitsArr.length > 1 && digitsArr[0] === 0)) {
    throw new Error("Input has wrong format");
  }
  const invalidDigits = digitsArr.filter(
    (digit) => digit < 0 || digit >= inBase
  );
  if (invalidDigits.length > 0) {
    throw new Error("Input has wrong format");
  }
  let decimalDigits =
    inBase === 10 ? digitsArr : toDecimalDigits(digitsArr, inBase);
  if (outBase === 10) {
    return decimalDigits;
  }
  let decimalNum = parseInt(decimalDigits.toString().split(",").join(""), 10);
  const outDigits = [];
  let remainder = decimalNum,
    result = decimalNum;
  let canDivideFurther = true;

  do {
    result = parseInt(decimalNum / outBase, 10);
    remainder = decimalNum % outBase;
    decimalNum = result;
    outDigits.push(remainder);
    if (result === 0) {
      canDivideFurther = false;
    }
  } while (canDivideFurther);
  return outDigits.reverse();
};

const toDecimalDigits = (digits, base) => {
  return digits
    .reverse()
    .reduce((accum, digit, i) => {
      return accum + digit * Math.pow(base, i);
    })
    .toString()
    .split("")
    .map((val) => parseInt(val, 10));
};
