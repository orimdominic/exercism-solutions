//
// This is only a SKELETON file for the 'Darts' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const score = function () {
  const [dartX, dartY] = [...arguments];
  const dartRadius = Math.sqrt(dartX ** 2 + dartY ** 2);
  const dartArea = parseFloat((Math.PI * dartRadius ** 2).toFixed(2));
  const innC = {
    minArea: 0,
    maxArea: parseFloat(Math.PI),
    point: 10,
  };
  const midC = {
    minArea: parseFloat((Math.PI).toFixed(2)),
    maxArea: parseFloat((Math.PI * 5 ** 2).toFixed(2)) ,
    point: 5,
  };
  const outC = {
    minArea: parseFloat((Math.PI * 5 ** 2) .toFixed(2)),
    maxArea: parseFloat((Math.PI * 10 ** 2).toFixed(2)) ,
    point: 1,
  };
  if (dartArea > outC.maxArea) {
    return 0;
  } else if (dartArea > outC.minArea && dartArea <= outC.maxArea) {
    return outC.point;
  } else if (dartArea > midC.minArea && dartArea <= midC.maxArea) {
    return midC.point;
  } else if (dartArea >= innC.minArea && dartArea < innC.maxArea) {
    return innC.point;
  }
};
