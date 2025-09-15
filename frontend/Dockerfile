FROM oven/bun:1.1-alpine AS base

WORKDIR /app

COPY package*.json ./
COPY bun.lockb* ./
RUN bun install --frozen-lockfile

COPY . .

RUN bun run build

FROM oven/bun:1.1-alpine AS runtime
WORKDIR /app

COPY --from=base /app /.

EXPOSE 4173

CMD ["bun", "run", "preview", "--", "--host", "0.0.0.0", "--strictPort"]
