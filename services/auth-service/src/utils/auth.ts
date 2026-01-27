import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "./db";
import * as schema from "./db/auth-schema";

export const auth = betterAuth({
    baseURL: process.env.BETTER_AUTH_BASE_URL,
    emailAndPassword: {
        enabled: true,
    },
    database: drizzleAdapter(db, {
        provider: "pg",
        schema: schema,
    }),
    trustedOrigins: [process.env.CORS_ORIGIN || "http://localhost:3001"],
    logger: {
        level: "debug",
        log: (l, m, ...a) => console.log(`[better-auth][${l}]`, m, ...a),
    },
});
