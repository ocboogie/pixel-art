export function random(size) {
  const cells = [];

  for (let y = 0; y < size; y += 1) {
    for (let x = 0; x < Math.ceil(size / 2); x += 1) {
      const pos = y * size + x;
      const mirroredPos = y * size + (size - 1 - x);
      const cell = Math.random() >= 0.5;
      cells[pos] = cell;
      if (mirroredPos !== pos) {
        cells[mirroredPos] = cell;
      }
    }
  }

  return cells;
}

export function serialize(cells, color) {
  let data = cells.map((cell) => (cell ? "1" : "0")).join("");
  data += color;

  return data;
}

export function deserialize(data) {
  const [cellsString, colorString] = data.split("#");
  const size = Math.ceil(Math.sqrt(cellsString.length));
  const cells = [...cellsString].map((cell) => cell === "1");
  const color = `#${colorString}`;

  return [size, cells, color];
}
