FROM oven/bun:latest AS base
WORKDIR /usr/src/app

FROM base AS install
RUN mkdir -p /temp
COPY package.json bun.lock /temp
RUN cd /temp && bun install --frozen-lockfile

FROM base AS release
COPY --from=install /temp/node_modules node_modules
COPY . .

USER bun
EXPOSE 3000/tcp
ENTRYPOINT [ "bun", "run", "src/index.ts" ]
