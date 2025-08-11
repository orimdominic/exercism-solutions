// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  let total = 0

  for (let i = 0; i < birdsPerDay.length; i++) {
    total = total + birdsPerDay[i]
  }

  return total
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  const weekStartIndex = 7 * (week - 1)
  let total = 0

  for (let i = weekStartIndex; i < weekStartIndex + 7; i++) {
    total = total + birdsPerDay[i]
  }

  return total
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let index = 0; index < birdsPerDay.length; index = index + 2) {
    if(index % 2 === 0){
      birdsPerDay[index] = birdsPerDay[index] + 1
    }
  }
  return birdsPerDay
}
