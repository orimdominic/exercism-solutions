//
// This is only a SKELETON file for the 'Square root' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const squareRoot = function () {
  let [n, x, y] = [...arguments];
  x = x || n / 2;
  y = y || 1;
  let checkX = (x + y) / 2;
  if ((checkX ** 2).toFixed(2) === n.toFixed(2)) {
    return parseInt(checkX, 10);
  }
  y = n / checkX;
  return squareRoot(n, checkX, y);
};
