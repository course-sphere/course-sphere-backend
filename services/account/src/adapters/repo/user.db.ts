import { eq } from "drizzle-orm";
import { User } from "../../domain/user";
import { UserRepository } from "../../ports/repo";
import { db } from "./db";
import { user } from "./db/schema";

export class DBUserRepository implements UserRepository {
    async get(id: string): Promise<User> {
        const [u] = await db.select().from(user).where(eq(user.id, id));

        return u;
    }
}
