import { betterAuth } from "better-auth";
import { jwt, username } from "better-auth/plugins";
import { z } from "zod";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "./db";
import * as schema from "./db/schema";

const roleSchema = z.enum(["student", "instructor", "admin"]);

export const auth = betterAuth({
    baseURL: process.env.BETTER_AUTH_BASE_URL,
    emailAndPassword: {
        enabled: true,
    },
    plugins: [username(), jwt()],
    user: {
        additionalFields: {
            role: {
                type: "string",
                required: true,
                defaultValue: "student",
                validator: {
                    input: roleSchema,
                    output: roleSchema,
                },
            },
        },
    },
    databaseHooks: {
        user: {
            create: {
                before: async (user) => {
                    const role = roleSchema.parse(user.role ?? "student");
                    return { data: { ...user, role } };
                },
            },
        },
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
