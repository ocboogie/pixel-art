import base64ToArrayBuffer from "./base64ToArrayBuffer";
import rgbToHex from "./rgbToHex";

export default data => {
  const buffer = base64ToArrayBuffer(data);
  const view = new DataView(buffer);
  let offset = 0;

  const width = view.getUint16(offset);
  offset += 2;
  const height = view.getUint16(offset);
  offset += 2;

  const colorAmount = view.getUint8(offset);
  offset += 1;

  const colors = [];
  for (let i = 0; i < colorAmount; i += 1) {
    colors.push(
      rgbToHex(
        view.getUint8(offset), // red
        view.getUint8(offset + 1), // blue
        view.getUint8(offset + 2) // green
      )
    );
    offset += 3;
  }

  const pixels = [];
  for (let i = 0; i < width * height; i += 1) {
    pixels.push(view.getUint8(offset));
    offset += 1;
  }

  return {
    width,
    height,
    colors,
    pixels
  };
};
