import http.server
import socketserver
import webbrowser
import os

PORT = 8080
# Ensure we serve from the directory where the script is located
os.chdir(os.path.dirname(os.path.abspath(__file__)))

Handler = http.server.SimpleHTTPRequestHandler

print(f"🚀 AI Command Center starting at http://localhost:{PORT}")
print("Press Ctrl+C to stop.")

# Automatically open the viewer
webbrowser.open(f"http://localhost:{PORT}/viewer.html")

with socketserver.TCPServer(("", PORT), Handler) as httpd:
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\nStopping server...")
        httpd.server_close()