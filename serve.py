import http.server
import socketserver
import webbrowser
import os
import json

PORT = 8080
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
PROJECT_ROOT = os.path.dirname(BASE_DIR)

class CommandCenterHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        # 1. Serve Server Info
        if self.path == '/server-info':
            self.send_response(200)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            info = {
                "language": "Python",
                "command": "python serve.py",
                "baseDir": BASE_DIR
            }
            self.wfile.write(json.dumps(info).encode())
            return

        # 2. Specialized logic for kanban.json
        clean_path = self.path.split('?')[0]
        if clean_path == '/kanban.json':
            local_path = os.path.join(BASE_DIR, 'kanban.json')
            root_path = os.path.join(PROJECT_ROOT, 'kanban.json')
            
            target = None
            if os.path.exists(local_path):
                target = local_path
            elif os.path.exists(root_path):
                target = root_path
            
            if target:
                self.send_response(200)
                self.send_header('Content-type', 'application/json')
                self.end_headers()
                with open(target, 'rb') as f:
                    self.wfile.write(f.read())
                return
            else:
                self.send_error(404, "kanban.json not found")
                return

        # 3. Serve UI files from skill folder
        # SimpleHTTPRequestHandler serves from CWD, so we need to ensure we serve from BASE_DIR
        os.chdir(BASE_DIR)
        return super().do_GET()

print(f"🚀 AI Command Center starting at http://localhost:{PORT}")
print(f"Serving from: {BASE_DIR}")
webbrowser.open(f"http://localhost:{PORT}/viewer.html")

with socketserver.TCPServer(("", PORT), CommandCenterHandler) as httpd:
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        httpd.server_close()