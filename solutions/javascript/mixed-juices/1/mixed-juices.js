// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export function timeToMixJuice(name) {
  const juiceTimeMap = {
    'Pure Strawberry Joy' : 0.5 , 
    'Energizer' : 1.5,
    'Green Garden':  1.5 ,
    'Tropical Island': 3, 
    'All or Nothing' :5
  }
  return juiceTimeMap[name] || 2.5
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export function limesToCut(wedgesNeeded, limes) {
  const limeWedgesMap = {
    small: 6,
    medium: 8,
    large: 10
  }
  const wedgesArr = limes.map((l) => limeWedgesMap[l]) 
  let i, wedges = 0;
  for (i = 0; i < wedgesArr.length; i++){
    if(wedges >= wedgesNeeded){
      break;
    }
    wedges += wedgesArr[i]
  }
  return i
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export function remainingOrders(timeLeft, orders) {
  const prepTimes = orders.map(timeToMixJuice)
  let i, timeUsed = 0;
  for (i = 0; i < prepTimes.length; i++){
    if(timeUsed >= timeLeft){
      break;
    }
    timeUsed += prepTimes[i]
  }
  return orders.slice(i)
}
