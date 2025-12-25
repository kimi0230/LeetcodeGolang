import os
import sys


def encrypt_file(file_path, password):
    print(f"🔒 Encrypting {file_path}...")
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()

        # 建立一個簡單但有效的加密保護頁面
        # 內容被放在隱藏區塊中，輸入正確密碼後才會顯示
        protected_html = f"""
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>受保護的內容</title>
    <style>
        body {{ font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f5f5f5; }}
        .box {{ background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); text-align: center; }}
        input {{ padding: 10px; width: 200px; border: 1px solid #ccc; border-radius: 4px; margin-bottom: 10px; }}
        button {{ padding: 10px 20px; background: #3f51b5; color: white; border: none; border-radius: 4px; cursor: pointer; }}
        button:hover {{ background: #303f9f; }}
        #content {{ display: none; }}
    </style>
</head>
<body>
    <div class="box" id="login-box">
        <h2>🔒 受保護的解法</h2>
        <p>請輸入密碼以查看解題思路</p>
        <input type="password" id="pass" placeholder="輸入密碼..." onpython="if(event.key==='Enter') verify()">
        <br>
        <button onclick="verify()">確認</button>
    </div>
    <div id="content">{content}</div>

    <script>
        function verify() {{
            const p = document.getElementById('pass').value;
            if (p === "{password}") {{
                document.getElementById('login-box').style.display = 'none';
                const content = document.getElementById('content');
                content.style.display = 'block';
                // 重新載入 Mermaid 或其他腳本（如果有的話）
                if (window.mermaid) mermaid.init();
            }} else {{
                alert('密碼錯誤！');
            }}
        }}
        // 支援 Enter 鍵
        document.getElementById('pass').addEventListener('keypress', function (e) {{
            if (e.key === 'Enter') verify();
        }});
    </script>
</body>
</html>
"""
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(protected_html)
        return True
    except Exception as e:
        print(f"❌ Error: {e}")
        return False


if __name__ == "__main__":
    site_dir = "site"
    # 優先從環境變數讀取密碼，本地測試若沒設定則預設為 "111"
    password = os.environ.get("SOLUTION_PASSWORD", "111")

    count = 0
    for root, dirs, files in os.walk(site_dir):
        if "README_solution" in root and "index.html" in files:
            file_path = os.path.join(root, "index.html")
            if encrypt_file(file_path, password):
                count += 1

    print(f"✅ Successfully protected {count} files.")
