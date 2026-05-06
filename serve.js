const http = require('http');
const fs = require('fs');
const path = require('path');
const { exec } = require('child_process');

const PORT = 8080;
const url = `http://localhost:${PORT}/viewer.html`;
const baseDir = __dirname;
const projectRootDir = path.join(baseDir, '..');

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
    // 1. Serve Server Info
    if (req.url === '/server-info') {
        res.writeHead(200, { 'Content-Type': 'application/json' });
        return res.end(JSON.stringify({ 
            language: 'Node.js', 
            command: 'node serve.js',
            baseDir: baseDir 
        }));
    }

    let urlPath = req.url.split('?')[0];
    if (urlPath === '/' || urlPath === '') urlPath = '/viewer.html';

    // 2. Specialized logic for kanban.json (Check local then parent)
    if (urlPath === '/kanban.json') {
        const localPath = path.join(baseDir, 'kanban.json');
        const rootPath = path.join(projectRootDir, 'kanban.json');
        
        if (fs.existsSync(localPath)) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            return fs.createReadStream(localPath).pipe(res);
        } else if (fs.existsSync(rootPath)) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            return fs.createReadStream(rootPath).pipe(res);
        } else {
            res.writeHead(404);
            return res.end('kanban.json not found in skill folder or project root.');
        }
    }

    // 3. Serve other files from skill folder
    const filePath = path.join(baseDir, urlPath);
    const extname = String(path.extname(filePath)).toLowerCase();
    const contentType = mimeTypes[extname] || 'application/octet-stream';

    fs.readFile(filePath, (error, content) => {
        if (error) {
            res.writeHead(error.code === 'ENOENT' ? 404 : 500);
            res.end(error.code === 'ENOENT' ? 'File not found: ' + urlPath : 'Server error');
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content, 'utf-8');
        }
    });
});

server.listen(PORT, () => {
    console.log(`🚀 AI Command Center starting at ${url}`);
    console.log(`Serving from: ${baseDir}`);
    console.log('Press Ctrl+C to stop.');

    const start = (process.platform === 'darwin' ? 'open' : process.platform === 'win32' ? 'start' : 'xdg-open');
    exec(`${start} ${url}`);
});
