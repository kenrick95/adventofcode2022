//@ts-check
import fs from 'node:fs';

const RESULT = {
  CORRECT: -1,
  INCORRECT: 1,
  INDETERMINATE: 0,
};

async function main() {
  const input = await fs.promises.readFile('day13.in', 'utf-8');

  let packetListsPart1 = [];
  let packetListsPart2 = [];
  let packetCount = 0;
  let ansPart1 = 0;
  for (const rawLine of input.split('\n')) {
    const line = rawLine.trim();
    if (line === '' && packetListsPart1.length === 2) {
      packetCount += 1;
      const res = getOrder(packetListsPart1[0], packetListsPart1[1]);
      //   console.log(`Pair ${packetCount}: ${res}`)
      if (res === RESULT.CORRECT) {
        ansPart1 += packetCount;
      }

      packetListsPart1 = [];
    } else if (line === '') {
      continue;
    } else {
      const parsedLine = JSON.parse(line);
      packetListsPart1.push(parsedLine);
      packetListsPart2.push(parsedLine);
    }
  }
  packetListsPart2.push(JSON.parse('[[2]]'));
  packetListsPart2.push(JSON.parse('[[6]]'));

  console.log('ansPart1', ansPart1);

  packetListsPart2.sort(getOrder);
  // Find index of "[[2]]" && "[[6]]"

  let ansPart2 = 1;
  for (let i = 0; i < packetListsPart2.length; i++) {
    const str = JSON.stringify(packetListsPart2[i]);
    if (str === '[[2]]' || str === '[[6]]') {
      ansPart2 = ansPart2 * (i + 1);
    }
  }
  console.log('ansPart2', ansPart2);
}
await main();

function getOrder(packetA, packetB) {
  //   console.log('Compare', packetA, packetB);
  if (typeof packetA === 'number' && typeof packetB === 'number') {
    if (packetA < packetB) {
      return RESULT.CORRECT;
    } else if (packetA > packetB) {
      return RESULT.INCORRECT;
    }
    return RESULT.INDETERMINATE;
  } else if (Array.isArray(packetA) && Array.isArray(packetB)) {
    const lenPacketA = packetA.length;
    const lenPacketB = packetB.length;
    for (let i = 0; i < Math.max(lenPacketA, lenPacketB); i++) {
      if (i >= lenPacketA) {
        return RESULT.CORRECT;
      } else if (i >= lenPacketB) {
        return RESULT.INCORRECT;
      }
      const res = getOrder(packetA[i], packetB[i]);
      if (res !== RESULT.INDETERMINATE) {
        return res;
      }
    }
    return RESULT.INDETERMINATE;
  } else {
    if (Array.isArray(packetA) && typeof packetB === 'number') {
      return getOrder(packetA, [packetB]);
    } else if (typeof packetA === 'number' && Array.isArray(packetB)) {
      return getOrder([packetA], packetB);
    }
  }

  return RESULT.INDETERMINATE;
}
