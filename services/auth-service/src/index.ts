import { Hono } from "hono";
import { auth } from "./utils/auth";
import { cors } from "hono/cors";

const app = new Hono();

app.use(
    "*",
    cors({
        origin: process.env.CORS_ORIGIN || "http://localhost:3001",
        allowHeaders: ["Content-Type", "Authorization"],
        allowMethods: ["POST", "GET", "OPTIONS"],
        exposeHeaders: ["Content-Length"],
        maxAge: 600,
        credentials: true,
    }),
);

app.on(["POST", "GET"], "/auth/*", (c) => auth.handler(c.req.raw));

export default app;
