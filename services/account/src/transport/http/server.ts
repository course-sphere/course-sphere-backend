import { Hono } from "hono";
import { auth } from "../../utils/auth";
import { cors } from "hono/cors";
import { UserRepository } from "../../ports/repo";
import * as user from "./handler/user";

export const serve = (userRepo: UserRepository) => {
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

    app.get("/user/:id", user.get(userRepo));

    app.on(["POST", "GET"], "/*", async (c) => {
        return await auth.handler(c.req.raw);
    });

    return app
};
