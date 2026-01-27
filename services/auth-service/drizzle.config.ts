import { defineConfig } from "drizzle-kit";

export default defineConfig({
    dialect: "postgresql",
    schema: "./src/utils/db/auth-schema.ts",
    out: "./drizzle",
    dbCredentials: {
        url:
            process.env.DATABASE_URL || "postgres://user:password@host:port/db",
    },
});
