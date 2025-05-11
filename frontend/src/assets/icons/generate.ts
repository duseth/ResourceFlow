import { Canvas, createCanvas } from 'canvas';
import * as fs from 'fs';
import * as path from 'path';

const generateIcon = (size: number, text = 'A'): Canvas => {
    const canvas = createCanvas(size, size);
    const ctx = canvas.getContext('2d');
    
    ctx.fillStyle = '#282c34';
    ctx.fillRect(0, 0, size, size);
    
    ctx.fillStyle = '#61dafb';
    ctx.font = `bold ${size * 0.6}px Arial`;
    ctx.textAlign = 'center';
    ctx.textBaseline = 'middle';
    ctx.fillText(text, size/2, size/2);
    
    return canvas;
};

const sizes = {
    favicon: [16, 32, 48],
    logo: [192, 512],
    maskable: [512]
};

const outputDir = path.join(__dirname);

try {
    // Favicon
    sizes.favicon.forEach(size => {
        const canvas = generateIcon(size);
        fs.writeFileSync(path.join(outputDir, `favicon-${size}.png`), canvas.toBuffer());
    });

    // Logo
    sizes.logo.forEach(size => {
        const canvas = generateIcon(size);
        fs.writeFileSync(path.join(outputDir, `logo${size}.png`), canvas.toBuffer());
    });

    // Maskable
    const maskableCanvas = generateIcon(512);
    fs.writeFileSync(path.join(outputDir, 'maskable.png'), maskableCanvas.toBuffer());

    console.log('Icons generated successfully');
} catch (error) {
    console.error('Error generating icons:', error);
}

export const icons = {
    favicon: '/favicon.ico',
    logo192: '/logo192.png',
    logo512: '/logo512.png',
    maskable: '/maskable.png'
}; 