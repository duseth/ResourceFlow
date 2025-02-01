import { createCanvas } from 'canvas';
import { writeFileSync } from 'fs';
import { join } from 'path';

const sizes = {
    favicon: [16, 32, 48],
    logo: [192, 512],
    maskable: [512]
};

const generateIcon = (size: number, text: string = 'A'): Canvas => {
    const canvas = createCanvas(size, size);
    const ctx = canvas.getContext('2d');
    
    // Фон
    ctx.fillStyle = '#282c34';
    ctx.fillRect(0, 0, size, size);
    
    // Текст
    ctx.fillStyle = '#61dafb';
    ctx.font = `bold ${size * 0.6}px Arial`;
    ctx.textAlign = 'center';
    ctx.textBaseline = 'middle';
    ctx.fillText(text, size/2, size/2);
    
    return canvas;
};

// Генерация иконок
const outputDir = join(__dirname);

// Favicon
sizes.favicon.forEach(size => {
    const canvas = generateIcon(size);
    writeFileSync(join(outputDir, `favicon-${size}.png`), canvas.toBuffer());
});

// Logo
sizes.logo.forEach(size => {
    const canvas = generateIcon(size);
    writeFileSync(join(outputDir, `logo${size}.png`), canvas.toBuffer());
});

// Maskable
const maskableCanvas = generateIcon(512);
writeFileSync(join(outputDir, 'maskable.png'), maskableCanvas.toBuffer());

export const icons = {
  // Здесь будут определения иконок
}; 