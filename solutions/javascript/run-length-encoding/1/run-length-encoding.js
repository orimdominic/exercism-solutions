//
// This is only a SKELETON file for the 'Run Length Encoding' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const encode = function () {
  const unencoded = [...arguments][0];
  if (!unencoded.length || unencoded.length === 1) {
    return unencoded;
  }
  let encoded = "",
    count = 0;
  for (let i = 0; i <= unencoded.length; i++) {
    const prevChar = unencoded[i - 1];
    const currChar = unencoded[i];
    if (prevChar === undefined || prevChar === currChar) {
      count++;
      continue;
    }
    encoded = encoded + `${count === 1 ? "" : count}` + prevChar;
    count = 1;
  }
  return encoded;
};

export const decode = function () {
  let encoded = [...arguments][0];
  if (!encoded.length || encoded.length === 1) {
    return encoded;
  }
  let tracingIntStr = "",
    currChar = "",
    decoded = "";
  for (let i = 0; i < encoded.length; i++) {
    currChar = encoded[i];
    if (Number.isInteger(parseInt(currChar, 10))) {
      tracingIntStr += currChar;
    } else {
      const times =
        parseInt(tracingIntStr, 10) < 2 || tracingIntStr.length === 0
          ? 1
          : parseInt(tracingIntStr, 10);
      decoded = decoded + currChar.repeat(times);
      tracingIntStr = "";
    }
  }
  return decoded;
};
