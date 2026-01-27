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
console.log(process.env.CORS_ORIGIN);
console.log(process.env.BETTER_AUTH_BASE_URL);

app.on(["POST", "GET"], "/*", async (c) => {
    return await auth.handler(c.req.raw);
});

export default app;
