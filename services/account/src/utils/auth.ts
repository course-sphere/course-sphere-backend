import { betterAuth } from "better-auth";
import { jwt, twoFactor, username } from "better-auth/plugins";
import { z } from "zod";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "../adapters/repo/db";
import * as schema from "../adapters/repo/db/schema";
import { template, transporter } from "./mail";

const roleSchema = z.enum(["student", "instructor", "admin"]);

export const auth = betterAuth({
    appName: "Course Sphere",
    baseURL: process.env.BETTER_AUTH_BASE_URL,
    emailAndPassword: {
        enabled: true,
        requireEmailVerification: true,
        sendResetPassword: async ({ user, token }, request) => {
            const url = `${request?.headers.get("origin")}/auth/reset-password?token=${encodeURIComponent(token)}`;

            await transporter.sendMail({
                to: user.email,
                subject: "Reset your password",
                html: await template(
                    "Reset Password",
                    "Click the button below to reset your password.",
                    user,
                    url,
                    request,
                ),
            });
        },
    },
    emailVerification: {
        autoSignInAfterVerification: true,
        sendVerificationEmail: async ({ user, url }, request) => {
            await transporter.sendMail({
                to: user.email,
                subject: "Verify your email address",
                html: await template(
                    "Verify Email",
                    "Click the button below to verify your email address.",
                    user,
                    url,
                    request,
                ),
            });
        },
    },
    account: {
        accountLinking: {
            trustedProviders: [
                "google",
                "microsoft",
                "github",
                "email-password",
            ],
        },
    },
    socialProviders: {
        google: {
            prompt: "select_account",
            clientId: process.env.GOOGLE_CLIENT_ID as string,
            clientSecret: process.env.GOOGLE_CLIENT_SECRET as string,
        },
        microsoft: {
            prompt: "select_account",
            clientId: process.env.MICROSOFT_CLIENT_ID as string,
            clientSecret: process.env.MICROSOFT_CLIENT_SECRET as string,
        },
        github: {
            clientId: process.env.GITHUB_CLIENT_ID as string,
            clientSecret: process.env.GITHUB_CLIENT_SECRET as string,
        },
    },
    plugins: [username(), jwt(), twoFactor()],
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
    trustedOrigins: [process.env.CORS_ORIGIN || "http://localhost:3002"],
    logger: {
        level: "debug",
        log: (l, m, ...a) => console.log(`[better-auth][${l}]`, m, ...a),
    },
    advanced: {
        database: {
            generateId: "uuid",
        },
    },
});
