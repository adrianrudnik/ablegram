export const colorMap: { [key: number]: string } = {
  // Taken from Ableton Track colorization
  // Column 2
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
  59: '#DBC300'
}

export function resolveColorByIndex(index: number | undefined): string | null {
  if (index === undefined) return null
  return colorMap[index] ?? null
}
