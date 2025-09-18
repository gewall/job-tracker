# Direktori
WEB_DIR=apps/web
API_DIR=apps/api
API_CMD=$(API_DIR)/cmd

# Default commands
.PHONY: dev frontend backend build clean

# Jalankan frontend & backend bersamaan
dev:
	@echo "🚀 Running frontend (Next.js) & backend (Go)..."
	(cd $(WEB_DIR) && pnpm dev) & \
	(cd $(API_CMD) && go run main.go)

# Jalankan frontend saja
frontend:
	@echo "🌐 Running Next.js frontend..."
	cd $(WEB_DIR) && pnpm dev

# Jalankan backend saja
backend:
	@echo "⚙️  Running Go backend..."
	cd $(API_CMD) && go run main.go

# Build frontend & backend
build:
	@echo "🔨 Building frontend & backend..."
	(cd $(WEB_DIR) && pnpm build)
	(cd $(API_CMD) && go build -o ../../../bin/api ./)

# Bersihkan hasil build
clean:
	@echo "🧹 Cleaning build..."
	rm -rf $(WEB_DIR)/.next
	rm -rf bin/
