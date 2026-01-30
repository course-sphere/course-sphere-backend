import { defineConfig } from "drizzle-kit";

export default defineConfig({
    dialect: "postgresql",
    schema: "./src/utils/db/schema.ts",
    out: "./drizzle",
    dbCredentials: {
        url:
            process.env.DATABASE_URL || "postgres://user:password@host:port/db",
    },
});
