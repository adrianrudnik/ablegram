export const colorMap: { [key: number]: string } = {
  // Taken from Ableton Track colorization
  // Column 1
  0: '#FF94A6',
  14: '#FF3636',
  28: '#E2675A',
  42: '#C6928B',
  56: '#AF3333',

  // Column 2
  1: '#FFA529',
  15: '#F66C03',
  29: '#FFA374',
  43: '#B78256',
  57: '#A95131',

  // Column 3
  2: '#CC9927',
  16: '#99724B',
  30: '#D3AD71',
  44: '#99836A',
  58: '#724F41',

  // Column 4
  3: '#F7F47C',
  17: '#FFF034',
  31: '#EDFFAE',
  45: '#BFBA69',
  59: '#DBC300',

  // Column 5
  4: '#BEFB00',
  18: '#86FF67',
  32: '#D1E397',
  46: '#A5BC00',
  60: '#849524',

  // Column 6
  5: '#1EFF32',
  19: '#3FC200',
  33: '#B9CF73',
  47: '#7CAE4E',
  61: '#549D34',

  // Column 7
  6: '#27FFA6',
  20: '#00BEAE',
  34: '#9AC38C',
  48: '#87C1B9',
  62: '#0F9B8D',

  // Column 8
  7: '#5BFFE7',
  21: '#19E8FF',
  35: '#D3FDE0',
  49: '#99B2C3',
  63: '#276383',

  // Column 9
  8: '#89C4FF',
  22: '#11A3EE',
  36: '#CCF1F8',
  50: '#84A4C1',
  64: '#1F3395',

  // Column 10
  9: '#547FE3',
  23: '#007CBF',
  37: '#B7C0E2',
  51: '#8292CB',
  65: '#3253A1',

  // Column 11
  10: '#90A6FF',
  24: '#876CE3',
  38: '#CCBAE3',
  52: '#A494B4',
  66: '#624DAC',

  // Column 12
  11: '#D76CE3',
  25: '#B576C5',
  39: '#AC97E4',
  53: '#BE9EBD',
  67: '#A24DAC',

  // Column 13
  12: '#E5549F',
  26: '#FF3CD3',
  40: '#E4DBE0',
  54: '#BB7195',
  68: '#CB326E',

  // Column 14
  13: '#FFFFFF',
  27: '#CFCFCF',
  41: '#A8A8A8',
  55: '#7A7A7A',
  69: '#3E3E3E'
}

export function resolveColorByIndex(index: number | undefined): string | null {
  if (index === undefined) return null
  return colorMap[index] ?? null
}
