const http = require('http');
const fs = require('fs');
const path = require('path');
const { exec } = require('child_process');

const PORT = 8080;
const url = `http://localhost:${PORT}/viewer.html`;
const baseDir = __dirname; // Always serve relative to script location

const mimeTypes = {
    '.html': 'text/html',
    '.js': 'text/javascript',
    '.css': 'text/css',
    '.json': 'application/json',
    '.png': 'image/png',
    '.jpg': 'image/jpg',
    '.gif': 'image/gif',
};

const server = http.createServer((req, res) => {
    let urlPath = req.url === '/' ? '/viewer.html' : req.url;
    let filePath = path.join(baseDir, urlPath);

    // Security: Prevent directory traversal outside of baseDir or its parent
    // (Parent allowed for kanban.json lookup)
    const relative = path.relative(path.join(baseDir, '..'), filePath);
    if (relative.startsWith('..')) {
        res.writeHead(403);
        return res.end('Forbidden');
    }

    const extname = String(path.extname(filePath)).toLowerCase();
    const contentType = mimeTypes[extname] || 'application/octet-stream';

    fs.readFile(filePath, (error, content) => {
        if (error) {
            if (error.code === 'ENOENT') {
                res.writeHead(404);
                res.end('File not found: ' + urlPath);
            } else {
                res.writeHead(500);
                res.end('Server error: ' + error.code);
            }
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content, 'utf-8');
        }
    });
});

server.listen(PORT, () => {
    console.log(`🚀 AI Command Center starting at ${url}`);
    console.log(`Serving files from: ${baseDir}`);
    console.log('Press Ctrl+C to stop.');

    const start = (process.platform === 'darwin' ? 'open' : process.platform === 'win32' ? 'start' : 'xdg-open');
    exec(`${start} ${url}`);
});
